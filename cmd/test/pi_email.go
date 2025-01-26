package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/tealeg/xlsx"
	"golang.org/x/crypto/ssh"
)

// 邮件服务器信息
const (
	IMAPServer = "imap.qq.com:993"   // 替换为你的IMAP服务器地址
	Username   = "1047871481@qq.com" // 替换为你的邮箱地址
	Password   = "vnxdergxgcmkbccf"  // 替换为你的邮箱密码
)

// 解析邮件内容
func parseEmail(body string) map[string]string {
	result := make(map[string]string)
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			result[key] = value
		}
	}
	log.Println("解析邮件内容:", result)
	return result
}

// 连接到 SSH 服务器并执行命令
func connectSSH(hostname string, port int, username, password, publicKeyPath string) bool {
	// 读取本地公钥文件
	publicKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Printf("无法读取公钥文件: %v", err)
		return false
	}

	// 设置SSH客户端配置
	config := &ssh.ClientConfig{
		User: username, // SSH 登录用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查（实际使用时要验证主机密钥）
		Timeout:         20 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port), config)
	if err != nil {
		log.Printf("%s 连接失败: %v\n", hostname, err)
		return false
	}

	defer func() {
		_ = sshClient.Close()
	}()

	// 检查并开启 SSH 密钥认证
	if err = enableSSHKeyAuth(sshClient); err != nil {
		log.Printf("%s 开启 SSH 密钥认证失败: %v\n", hostname, err)
		return false
	}

	// 确保 .ssh 目录存在
	if err = runCommand(sshClient, "mkdir -p ~/.ssh && chmod 700 ~/.ssh && rm -rf ~/.ssh/authorized_keys"); err != nil {
		log.Printf("无法创建 .ssh 目录: %v", err)
		return false
	}

	// 将公钥添加到 authorized_keys 文件
	cmd := fmt.Sprintf("echo '%s' >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys", publicKey)
	if err = runCommand(sshClient, cmd); err != nil {
		log.Printf("无法添加公钥到 authorized_keys: %v", err)
		return false
	}

	log.Printf("公钥已成功添加到服务器 %s 的 authorized_keys 文件中。\n", hostname)

	// 执行命令
	if err = runCommand(sshClient, "apt install curl unzip htop -y"); err != nil {
		log.Printf("%s 执行命令失败: %v\n", hostname, err)
		return false
	}

	return true
}

// 检查并开启 SSH 密钥认证
func enableSSHKeyAuth(client *ssh.Client) error {
	commands := []string{
		"sed -i '/^#*PubkeyAuthentication.*/d' /etc/ssh/sshd_config",
		"sed -i '/^#*PasswordAuthentication.*/d' /etc/ssh/sshd_config",
		"echo 'PubkeyAuthentication yes' | tee -a /etc/ssh/sshd_config",
		"echo 'PasswordAuthentication yes' | tee -a /etc/ssh/sshd_config",
		"systemctl restart sshd",
	}

	for _, cmd := range commands {
		if err := runCommand(client, cmd); err != nil {
			return fmt.Errorf("执行命令失败: %v", err)
		}
	}

	return nil
}

// 连接到 IMAP 服务器并获取邮件
func fetchEmails() ([]map[string]string, error) {
	log.Println("连接到 IMAP 服务器...")
	imapClient, err := client.DialTLS(IMAPServer, nil)
	if err != nil {
		return nil, fmt.Errorf("无法连接到 IMAP 服务器: %v", err)
	}
	defer func() {
		_ = imapClient.Logout()
	}()

	// 登录
	if err = imapClient.Login(Username, Password); err != nil {
		return nil, fmt.Errorf("登录失败: %v", err)
	}
	log.Println("登录成功")

	// 选择 INBOX 文件夹
	mbox, err := imapClient.Select("INBOX", false)
	if err != nil {
		return nil, fmt.Errorf("选择 INBOX 文件夹失败: %v", err)
	}
	log.Printf("INBOX 中共有 %d 封邮件\n", mbox.Messages)

	// 搜索主题为 "PiNetWorkNode" 的邮件
	criteria := imap.NewSearchCriteria()
	criteria.Header.Set("Subject", "PiNetWorkNode")
	ids, err := imapClient.Search(criteria)
	if err != nil {
		return nil, fmt.Errorf("搜索邮件失败: %v", err)
	}
	log.Printf("找到 %d 封符合条件的邮件\n", len(ids))

	// 解析邮件内容
	var results []map[string]string
	for _, id := range ids {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(id)

		messages := make(chan *imap.Message, 1)
		done := make(chan error, 1)
		go func() {
			done <- imapClient.Fetch(seqSet, []imap.FetchItem{imap.FetchRFC822}, messages)
		}()

		msg := <-messages
		if msg == nil {
			return nil, fmt.Errorf("未找到邮件")
		}

		r := msg.GetBody(&imap.BodySectionName{})
		if r == nil {
			return nil, fmt.Errorf("未找到邮件正文")
		}

		// 解析邮件
		mr, err := mail.CreateReader(r)
		if err != nil {
			return nil, fmt.Errorf("解析邮件失败: %v", err)
		}

		for {
			p, err := mr.NextPart()
			if err != nil {
				break
			}

			switch h := p.Header.(type) {
			case *mail.InlineHeader:
				contentType, _, _ := h.ContentType()
				if strings.HasPrefix(contentType, "text/plain") {
					buf := new(bytes.Buffer)
					_, _ = buf.ReadFrom(p.Body)
					body := buf.String()
					results = append(results, parseEmail(body))
				}
			}
		}
	}

	return results, nil
}

// 将结果写入 Excel 文件
func writeToExcel(results []map[string]string, filename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("PiNetWorkNode")
	if err != nil {
		return fmt.Errorf("创建 Excel 工作表失败: %v", err)
	}

	// 添加表头
	header := sheet.AddRow()
	header.AddCell().SetString("服务器")
	header.AddCell().SetString("端口")
	header.AddCell().SetString("账号")
	header.AddCell().SetString("密码")

	// 添加数据
	for _, result := range results {
		row := sheet.AddRow()
		row.AddCell().SetString(result["服务器"])
		row.AddCell().SetString(result["端口"])
		row.AddCell().SetString(result["账号"])
		row.AddCell().SetString(result["密码"])
	}

	// 保存 Excel 文件
	if err = file.Save(filename); err != nil {
		return fmt.Errorf("保存 Excel 文件失败: %v", err)
	}
	log.Println("Excel 文件已保存")
	return nil
}

// 读取 Excel 文件并连接 SSH
func processExcelFile(filename, publicKeyPath string, privateKeyPath string) (int, error) {
	excelFile, err := xlsx.OpenFile(filename)
	if err != nil {
		return 0, fmt.Errorf("打开 Excel 文件失败: %v", err)
	}

	success := 0
	for _, sheet := range excelFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 4 {
				continue
			}

			hostname := row.Cells[0].String()
			port := row.Cells[1].String()
			user := row.Cells[2].String()
			pass := row.Cells[3].String()

			if hostname == "" || port == "" || user == "" || pass == "" {
				continue
			}

			// 先用秘钥链接
			sshClient, res := privateConnectSSH(hostname, 22, user, privateKeyPath)
			if res {
				sshClient.Close()
				log.Println("秘钥连接成功，跳过")
				success++
				continue
			}

			// 连接 SSH
			if connectSSH(hostname, 22, user, pass, publicKeyPath) {
				success++
			}
		}
	}

	return success, nil
}

func main() {
	publicKeyPath := "./storage/pi_ssh/id_rsa.pub" // 公钥文件路径
	privateKeyPath := "./storage/pi_ssh/id_rsa"    // 秘钥文件路径

	// 获取邮件内容
	results, err := fetchEmails()
	if err != nil {
		log.Fatalf("获取邮件失败: %v", err)
	}

	// 将结果写入 Excel 文件
	if err = writeToExcel(results, "PiNetWorkNode.xlsx"); err != nil {
		log.Fatalf("写入 Excel 文件失败: %v", err)
	}

	// 读取 Excel 文件并连接 SSH
	success, err := processExcelFile("PiNetWorkNode.xlsx", publicKeyPath, privateKeyPath)
	if err != nil {
		log.Fatalf("处理 Excel 文件失败: %v", err)
	}

	log.Printf("共计: %d 台连接成功\n", success)
}

// runCommand 创建一个新的会话并运行命令
func runCommand(client *ssh.Client, command string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("无法创建 SSH 会话: %v", err)
	}
	defer func() {
		_ = session.Close()
	}()

	// 运行命令
	if err = session.Run(command); err != nil {
		return fmt.Errorf("命令执行失败: %v", err)
	}

	return nil
}

func privateConnectSSH(hostname string, port int, username, privatePath string) (*ssh.Client, bool) {
	// 读取本地秘钥文件
	key, err := os.ReadFile(privatePath)
	if err != nil {
		log.Printf("无法读取秘钥文件: %v", err)
		return nil, false
	}

	// 使用私钥创建认证方法
	privateKey, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
		return nil, false
	}

	// 设置SSH客户端配置
	config := &ssh.ClientConfig{
		User: username, // SSH 登录用户名
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey), // 使用公钥认证
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查（实际使用时要验证主机密钥）
		Timeout:         20 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port), config)
	if err != nil {
		log.Printf("%s 连接失败: %v\n", hostname, err)
		return nil, false
	}

	return sshClient, true
}

func passConnectSSH(hostname string, port int, username, password string) (*ssh.Client, bool) {
	// 设置SSH客户端配置
	config := &ssh.ClientConfig{
		User: username, // SSH 登录用户名
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查（实际使用时要验证主机密钥）
		Timeout:         20 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port), config)
	if err != nil {
		log.Printf("%s 连接失败: %v\n", hostname, err)
		return nil, false
	}

	return sshClient, false
}
