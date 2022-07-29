package http

import (
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/clients"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/ctx"
	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/errorx"
)

// ResponseBase 基础返回参数
type ResponseBase struct {
	App       string      `json:"app"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data"`
}

// ClientObject http请求代理变量
var ClientObject = NewHTTPClient()

// SdkClient 调用内部服务接口Client
type SdkClient struct {
	Client *clients.HTTPClient
}

// NewHTTPClient 获取http请求代理
func NewHTTPClient(options ...clients.HTTPClientOption) *SdkClient {
	var client SdkClient
	client.Client = clients.NewHTTPClient(options...)
	return &client
}

// Post 发送 Post 请求
func (s *SdkClient) Post(c ctx.Context, url string, options ...clients.Option) (*clients.ResponseProxy, error) {
	return s.Client.Post(c, url, options...)
}

// Get 发送 Get 请求
func (s *SdkClient) Get(c ctx.Context, url string, options ...clients.Option) (*clients.ResponseProxy, error) {
	return s.Client.Get(c, url, options...)
}

// CheckResponse 校验返回数据结构
func CheckResponse(response ResponseBase) error {
	if response.Code != 0 {
		return nil
	}
	return errorx.NewErrorMsg(response.Code, response.Msg)
}
