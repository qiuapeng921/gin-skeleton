package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"os"
	"strings"
	"time"
)

// 全局变量
var (
	parentMessageId  any
	chatSessionId    = ""
	authorization    = ""
	client           = resty.New()
	defaultAuthToken = "9v8X5BNIQM0hj8hI4p4qAgI9CwdPsYWDaiqHSHxnFZ5vF6vjMMZEcdgAuJxEFI0S"
)

// 主函数
func main() {
	inputToken()
	for {
		action := chooseAction()
		switch action {
		case "1":
			if err := chatList(); err != nil {
				fmt.Println("获取会话列表失败:", err)
			}
		case "2":
			if err := chatCreate(); err != nil {
				fmt.Println("创建新会话失败:", err)
			}
		case "exit":
			fmt.Println("程序已退出。")
			return
		default:
			fmt.Println("无效输入，请重新选择。")
		}
		chatLoop()
	}
}

// 输入 Token
func inputToken() {
	for {
		fmt.Print("请输入授权 Token [留空则使用默认 Token]: ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		authorization = strings.TrimSpace(input)

		if authorization == "" {
			authorization = defaultAuthToken
			fmt.Println("已使用默认 Token。")
		}

		if len(authorization) < 10 { // 假设 Token 长度至少为 10
			fmt.Println("Token 格式错误，请重新输入。")
			continue
		}

		fmt.Println("Token 输入成功。")
		break
	}
}

// 用户选择操作
func chooseAction() string {
	fmt.Println("\n请选择操作：")
	fmt.Println("1. 选择会话")
	fmt.Println("2. 创建新会话")
	fmt.Print("请输入选项 [ 1或2 输入 exit 退出]: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(input)
}

// 聊天循环
func chatLoop() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n进入聊天模式，输入 'exit' 退出会话。")
	for {
		fmt.Print("提问: ")
		inputText, _ := reader.ReadString('\n')
		inputText = strings.TrimSpace(inputText)

		if inputText == "exit" {
			fmt.Println("退出当前会话。")
			break
		}

		if err := sendChatRequest(inputText); err != nil {
			fmt.Println("请求处理错误:", err)
		}
		fmt.Println("\n------------------------")
		time.Sleep(500 * time.Millisecond)
	}
}

// 获取会话列表
func chatList() error {
	var resp SessionListResp

	_, err := client.R().
		SetAuthToken(authorization).
		SetQueryParam("count", "500").
		SetResult(&resp).
		Get("https://chat.deepseek.com/api/v0/chat_session/fetch_page")

	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("获取会话失败: %s", resp.Msg)
	}

	chatSessions := make([]string, 0)
	for _, session := range resp.Data.BizData.ChatSessions {
		fmt.Printf("ID: %s, Title: %s\n", session.Id, session.Title)
		chatSessions = append(chatSessions, session.Id)
	}

	return selectChatSession(chatSessions)
}

// 选择会话
func selectChatSession(chatSessions []string) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入会话ID: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("会话ID 不能为空，请重新输入。")
			continue
		}

		if !contains(chatSessions, input) {
			fmt.Printf("无效的会话ID: %s，请重新输入。\n", input)
			continue
		}

		currentMessageId, err := fetchHistoryMessages(input)
		if err != nil {
			fmt.Printf("获取会话信息失败: %s\n", err.Error())
			continue
		}

		chatSessionId = input
		parentMessageId = currentMessageId
		fmt.Printf("已选择会话ID: %s\n", chatSessionId)
		break
	}

	return nil
}

// 创建新会话
func chatCreate() error {
	var resp struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			BizData struct {
				Id string `json:"id"`
			} `json:"biz_data"`
		} `json:"data"`
	}

	_, err := client.R().
		SetAuthToken(authorization).
		SetBody(map[string]string{"agent": "chat"}).
		SetResult(&resp).
		Post("https://chat.deepseek.com/api/v0/chat_session/create")

	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("创建会话失败: %s", resp.Msg)
	}

	chatSessionId = resp.Data.BizData.Id
	parentMessageId = nil
	fmt.Printf("已创建新会话，ID: %s\n", chatSessionId)
	return nil
}

// 获取会话历史消息
func fetchHistoryMessages(sessionId string) (int, error) {
	var resp historyMessagesResp

	_, err := client.R().
		SetAuthToken(authorization).
		SetQueryParam("chat_session_id", sessionId).
		SetResult(&resp).
		Get("https://chat.deepseek.com/api/v0/chat/history_messages")

	if err != nil {
		return 0, fmt.Errorf("请求错误: %w", err)
	}

	if resp.Code != 0 {
		return 0, fmt.Errorf("获取会话失败: %s", resp.Msg)
	}

	return resp.Data.BizData.ChatSession.CurrentMessageId, nil
}

// 发送聊天请求
func sendChatRequest(text string) error {
	requestBody := completionReq{
		ChatSessionId:     chatSessionId,
		ParentMessageId:   parentMessageId,
		Prompt:            text,
		RefFileIds:        []any{},
		ThinkingEnabled:   false,
		SearchEnabled:     false,
		ChallengeResponse: nil,
	}

	resp, err := client.R().
		SetDoNotParseResponse(true).
		SetAuthToken(authorization).
		SetBody(requestBody).
		Post("https://chat.deepseek.com/api/v0/chat/completion")

	if err != nil {
		return fmt.Errorf("请求错误: %w", err)
	}
	defer resp.RawBody().Close()

	return processStreamResponse(resp.RawBody())
}

// 处理流式响应
func processStreamResponse(body io.Reader) error {
	fmt.Print("答复: ")
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.HasPrefix(line, "data: ") {
			continue
		}
		line = strings.TrimPrefix(line, "data: ")
		var data struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
			MessageId int `json:"message_id"`
		}
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			continue
		}
		for _, choice := range data.Choices {
			fmt.Print(choice.Delta.Content)
			time.Sleep(50 * time.Millisecond) // 模拟流效果
		}
		parentMessageId = data.MessageId
	}

	return scanner.Err()
}

// 判断切片中是否包含目标值
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

// 请求结构体
type completionReq struct {
	ChatSessionId     string        `json:"chat_session_id"`
	ParentMessageId   interface{}   `json:"parent_message_id"`
	Prompt            string        `json:"prompt"`
	RefFileIds        []interface{} `json:"ref_file_ids"`
	ThinkingEnabled   bool          `json:"thinking_enabled"`
	SearchEnabled     bool          `json:"search_enabled"`
	ChallengeResponse interface{}   `json:"challenge_response"`
}

type SessionListResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		BizData struct {
			ChatSessions []struct {
				Id    string `json:"id"`
				Title string `json:"title"`
			} `json:"chat_sessions"`
		} `json:"biz_data"`
	} `json:"data"`
}

// 历史消息响应结构
type historyMessagesResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		BizData struct {
			ChatSession struct {
				CurrentMessageId int `json:"current_message_id"`
			} `json:"chat_session"`
		} `json:"biz_data"`
	} `json:"data"`
}
