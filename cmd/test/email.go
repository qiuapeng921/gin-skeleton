package main

import (
	"fmt"
	"github.com/emersion/go-message/mail"
	"golang.org/x/net/html"
	"io"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

var addr = "imap.qiye.aliyun.com:993"
var username = "us-cs001@khdtek.com"
var password = "khd=20221208"

//var addr = "mail.vecelo.com:993"
//var username = "support@vecelo.com"
//var password = "*uvcG]7?=Y}a"

func login() *client.Client {
	c, err := client.DialTLS(addr, nil) // 替换为你的 IMAP 服务器
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Login(username, password); err != nil { // 替换为你的邮箱和密码
		log.Fatal(err)
	}
	return c
}

func main() {
	// 连接到 IMAP 服务器
	log.Println("连接到 IMAP 服务器...")

	c := login()

	// 选择 INBOX 文件夹
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("邮箱中共有 %d 封邮件\n", mbox.Messages)

	// 第一步：获取所有邮件并存储
	getAllEmails(c, mbox.Messages)

	log.Println("开始监听新邮件...")
	// 第二步：实时监听新邮件
	listenForNewEmails(c, mbox.Messages)

	// Create a channel to receive mailbox updates
	updates := make(chan client.Update)
	c.Updates = updates

	// Start idling
	stopped := false
	stop := make(chan struct{})
	done := make(chan error, 1)
	go func() {
		done <- c.Idle(stop, &client.IdleOptions{PollInterval: time.Minute})
	}()

	// Listen for updates
	for {
		select {
		case update := <-updates:
			log.Println("New update:", update)
			if !stopped {
				close(stop)
				stopped = true
			}
		case err = <-done:
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Not idling anymore")
			return
		}
	}
}

// 获取所有邮件并存储
func getAllEmails(c *client.Client, total uint32) {
	if total == 0 {
		log.Println("没有邮件需要获取")
		return
	}

	// 创建范围（1 到 total）
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(1, total)

	messages := make(chan *imap.Message, 10)
	section := &imap.BodySectionName{}

	// 异步获取邮件
	go func() {
		if err := c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages); err != nil {
			c = login()
			log.Fatal("获取邮件失败: ", err)
		}
	}()

	// 遍历并存储邮件
	log.Println("开始获取所有邮件...")
	messageBody(messages)

	log.Println("所有邮件获取完成")
}

// 实时监听新邮件
func listenForNewEmails(c *client.Client, lastSeen uint32) {
	for {
		time.Sleep(time.Second * 3)
		// 刷新邮箱，检查是否有新邮件
		mbox, err := c.Select("INBOX", false)
		if err != nil {
			log.Println("刷新邮箱失败,重新登录")
			c = login()
			continue
		}

		if mbox.Messages > lastSeen {
			newMailCount := mbox.Messages - lastSeen
			log.Printf("检测到 %d 封新邮件", newMailCount)

			// 获取新邮件详情
			seqSet := new(imap.SeqSet)
			seqSet.AddRange(lastSeen+1, mbox.Messages)
			messages := make(chan *imap.Message, newMailCount)

			section := &imap.BodySectionName{}

			go func() {
				if err := c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}, messages); err != nil {
					c = login()
					log.Fatal("获取邮件失败: ", err)
				}
			}()

			messageBody(messages)

			// 更新最后一封邮件的编号
			lastSeen = mbox.Messages

		}
	}
}

func messageBody(messages chan *imap.Message) {
	for msg := range messages {
		//form := msg.Envelope.From[0]
		//address := form.MailboxName + "@" + form.HostName
		//log.Printf("新邮件: %s | 发件人: %s\n", msg.Envelope.Subject, address)
		if strings.Contains(msg.Envelope.Subject, "Account data access attempt") {
			r := msg.GetBody(&imap.BodySectionName{})
			if r == nil {
				log.Fatal("服务器未返回邮件正文")
			}
			mr, err := mail.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			// 处理邮件正文
			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal("NextPart:err ", err)
				}

				switch p.Header.(type) {
				case *mail.InlineHeader:
					// 读取正文内容
					b, _ := io.ReadAll(p.Body)
					bodyText := string(b)

					// 解析HTML内容
					doc, err := html.Parse(strings.NewReader(bodyText))
					if err != nil {
						log.Fatal("Failed to parse HTML:", err)
					}

					extractText(msg, doc)
				}
			}
		}
	}

}

// extractText 从 HTML 文档中提取 <p> 标签中的 6 位数字
func extractText(msg *imap.Message, doc *html.Node) {
	if doc.Type == html.ElementNode && (doc.Data == "p" || doc.Data == "span") {
		for c := doc.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				text := strings.TrimSpace(c.Data)
				// 匹配6位数字
				re := regexp.MustCompile(`^\d{6}$`)
				if re.MatchString(text) {
					fmt.Printf("%s 验证码在 <%s>标签, 值为: %s \n", msg.Envelope.Date, doc.Data, text)
				}
			}
		}
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		extractText(msg, c)
	}
}
