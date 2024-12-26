package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestBody struct {
	Text      string `json:"text"`
	SessionId int    `json:"sessionId"`
	Files     []any  `json:"files"`
}

type ResponseData struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ChatListResp struct {
	Resp
	Data []ChatDataResp `json:"data"`
}

type ChatDataResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var sessionId = 0 // 默认会话ID
var authorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjg4Niwic2lnbiI6Ijc1MWUwYTRkYjQ3MGM3Y2I1ZTQ0ZDIzMTEwNjBjZGE3Iiwicm9sZSI6InVzZXIiLCJleHAiOjE3Mzc1OTcxOTgsIm5iZiI6MTczNDkxODc5OCwiaWF0IjoxNzM0OTE4Nzk4fQ.qgJIYJ9GIKmwijUBpGFCv2cvMupODMZMzm2Q_d3Fc80"

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 让用户选择操作
	for {
		fmt.Println("请选择操作：")
		fmt.Println("1. 从会话列表中选择")
		fmt.Println("2. 创建新会话")
		fmt.Print("请输入选项(1或2, 输入exit退出): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("退出程序。")
			break
		}

		switch input {
		case "1":
			err := chatList() // 调用获取会话列表函数
			if err != nil {
				fmt.Println("获取会话列表失败:", err)
			}
		case "2":
			err := chatCreate() // 调用创建会话函数
			if err != nil {
				fmt.Println("创建新会话失败:", err)
			}
		default:
			fmt.Println("无效输入，请重新选择。")
		}

		// 进入聊天主循环
		chatLoop()
	}
}

// 聊天主循环
func chatLoop() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入你的问题(输入exit退出): ")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.TrimSpace(inputText)

		if inputText == "exit" {
			fmt.Println("退出当前会话。")
			break
		}

		err := sendChatRequest(inputText)
		if err != nil {
			fmt.Println("请求处理错误:", err)
		}

		fmt.Println("\n ------------------------")
		time.Sleep(500 * time.Millisecond) // 延迟，避免请求过于频繁
	}
}

// 获取会话列表
func chatList() error {
	client := &http.Client{}
	url := "https://gpt-all.chat/api/chat/session"

	headers := map[string]string{
		"Authorization": authorization,
		"Content-Type":  "application/json",
	}

	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求错误: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var chatListResp ChatListResp
	if err = json.NewDecoder(resp.Body).Decode(&chatListResp); err != nil {
		return fmt.Errorf("解析响应错误: %w", err)
	}

	if chatListResp.Code != 0 {
		return fmt.Errorf("获取会话列表失败: %s", chatListResp.Msg)
	}

	// 显示会话列表
	fmt.Println("可用会话列表:")
	for _, chat := range chatListResp.Data {
		fmt.Printf("ID: %d, Name: %s\n", chat.Id, chat.Name)
	}

	// 选择会话
	fmt.Print("请输入会话ID: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	selectedId, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("无效的会话ID: %w", err)
	}

	sessionId = selectedId // 设置全局会话ID
	fmt.Printf("已选择会话ID: %d\n", sessionId)
	return nil
}

// 创建新会话
func chatCreate() error {
	client := &http.Client{}
	url := "https://gpt-all.chat/api/chat/session"

	headers := map[string]string{
		"Authorization": authorization,
		"Content-Type":  "application/json",
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求错误: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var chatResp struct {
		Code int          `json:"code"`
		Msg  string       `json:"msg"`
		Data ChatDataResp `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return fmt.Errorf("解析响应错误: %w", err)
	}

	if chatResp.Code != 0 {
		return fmt.Errorf("创建会话失败: %s", chatResp.Msg)
	}

	// 设置全局会话ID
	sessionId = chatResp.Data.Id
	fmt.Printf("已创建新会话，ID: %d, 名称: %s\n", chatResp.Data.Id, chatResp.Data.Name)
	return nil
}

// 发送聊天请求
func sendChatRequest(text string) error {
	client := &http.Client{}
	url := "https://gpt-all.chat/api/chat/completions"

	headers := map[string]string{
		"Authorization": authorization,
		"Content-Type":  "application/json",
		"Accept":        "text/event-stream",
		"Connection":    "keep-alive",
	}

	// 构造请求体
	requestBody := RequestBody{
		Text:      text,
		SessionId: sessionId,
		Files:     []any{},
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("JSON 编码错误: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("创建请求错误: %w", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}
	defer resp.Body.Close()

	// 处理响应流
	return processResponse(resp.Body)
}

// 处理响应流
func processResponse(body io.ReadCloser) error {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return fmt.Errorf("读取响应错误: %w", err)
	}

	responseLines := bytes.Split(buf.Bytes(), []byte("\n"))

	for _, line := range responseLines {
		trimmedLine := bytes.TrimSpace(line)
		if len(trimmedLine) == 0 {
			continue // 跳过空行
		}

		jsonString := bytes.Replace(trimmedLine, []byte("data: "), []byte(""), 1)

		var jsonData ResponseData
		if err := json.Unmarshal(jsonString, &jsonData); err == nil && jsonData.Type == "string" {
			fmt.Print(jsonData.Data)
			time.Sleep(50 * time.Millisecond) // 模拟流式输出
		}
	}

	return nil
}
