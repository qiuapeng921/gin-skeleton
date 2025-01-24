package main

import (
	"bytes"
	"fmt"
	"log"
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

	log.Println(result)

	return result
}

// 连接到 SSH 服务器并执行命令
func connectSSH(hostname string, port int, username string, password string) bool {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port), config)
	if err != nil {
		log.Printf("%s 连接失败: %v\n", hostname, err.Error())
		return false
	}
	defer func() {
		_ = sshClient.Close()
	}()

	session, err := sshClient.NewSession()
	if err != nil {
		log.Printf("%s 创建会话失败: %v\n", hostname, err)
		return false
	}
	defer func() {
		_ = session.Close()
	}()

	// 执行命令
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("apt install curl unzip -y && ./agent.sh uninstall"); err != nil {
		log.Printf("%s 执行命令失败: %v\n", hostname, err)
		return false
	}

	log.Printf("%s 连接成功并执行命令\n", hostname)
	return true
}

func main() {
	// 连接到 IMAP 服务器
	log.Println("连接到 IMAP 服务器...")
	imapClient, err := client.DialTLS(IMAPServer, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = imapClient.Logout()
	}()

	// 登录
	if err := imapClient.Login(Username, Password); err != nil {
		log.Fatal(err)
	}
	log.Println("登录成功")

	// 选择 INBOX 文件夹
	mbox, err := imapClient.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("INBOX 中共有 %d 封邮件\n", mbox.Messages)

	// 搜索主题为 "PiNetWorkNode" 的邮件
	criteria := imap.NewSearchCriteria()
	criteria.Header.Set("Subject", "PiNetWorkNode")
	ids, err := imapClient.Search(criteria)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal("未找到邮件")
		}

		r := msg.GetBody(&imap.BodySectionName{})
		if r == nil {
			log.Fatal("未找到邮件正文")
		}

		// 解析邮件
		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Fatal(err)
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

	// 将结果写入 Excel 文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("PiNetWorkNode")
	if err != nil {
		log.Fatal(err)
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
	if err := file.Save("PiNetWorkNode_emails.xlsx"); err != nil {
		log.Fatal(err)
	}
	log.Println("Excel 文件已保存")

	// 读取 Excel 文件并连接 SSH
	excelFile, err := xlsx.OpenFile("PiNetWorkNode_emails.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	success := 0
	for _, sheet := range excelFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 4 {
				continue
			}

			hostname := row.Cells[0].String()
			port := row.Cells[1].String()
			username := row.Cells[2].String()
			password := row.Cells[3].String()

			if hostname == "" || port == "" || username == "" || password == "" {
				continue
			}

			// 连接 SSH
			if connectSSH(hostname, 22, username, password) {
				success++
			}
		}
	}

	log.Printf("共计: %d 台连接成功\n", success)
}
