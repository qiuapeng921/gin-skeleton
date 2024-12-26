package main

import (
	"fmt"
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

// var addr = "imap.qiye.aliyun.com:993"
// var username = "us-cs001@khdtek.com"
// var password = "khd=20221208"
var addr = "mail.vecelo.com:993"
var username = "support@vecelo.com"
var password = "*uvcG]7?=Y}a"

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
	fmt.Print("Enter Password: ")
	return
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

	// 第二步：实时监听新邮件
	//listenForNewEmails(c, mbox.Messages)

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
	for msg := range messages {
		form := msg.Envelope.From[0]
		address := form.MailboxName + "@" + form.HostName
		log.Printf("邮件编号: %d | 时间: %s | 主题: %s | 发件人: %s \n", msg.SeqNum, msg.Envelope.Date, msg.Envelope.Subject, address)
	}
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

			go func() {
				if err = c.Fetch(seqSet, []imap.FetchItem{imap.FetchEnvelope}, messages); err != nil {
					log.Printf("获取新邮件失败: %v", err)
				}
			}()

			for msg := range messages {
				form := msg.Envelope.From[0]
				address := form.MailboxName + "@" + form.HostName
				log.Printf("新邮件: %s | 发件人: %s\n", msg.Envelope.Subject, address)
			}

			// 更新最后一封邮件的编号
			lastSeen = mbox.Messages
		}
	}
}
