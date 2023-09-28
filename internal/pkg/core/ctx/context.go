package ctx

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/language"
)

// RequestIDKey id 存储 key
const (
	RequestIDKey string = "x-request-id"
	LocaleKey           = "locale"
	LangKey             = "language"
	TokenIDKey          = "TokenIDKey"
	AuthInfoKey         = "AuthInfoKey"
)

type Context interface {
	context.Context
	ID() string
	AuthInfo() any
	SetAuthInfo(tokenId string, info any)
}

type defaultContext struct {
	context.Context
	id       string
	tokenID  string
	authInfo any
}

func (d *defaultContext) AuthInfo() any {
	return d.authInfo
}

func (d *defaultContext) SetAuthInfo(tokenID string, info any) {
	d.tokenID = tokenID
	d.authInfo = info
}

func (d *defaultContext) ID() string {
	return d.id
}

// New 新建 context
func New() Context {
	p, id := xReqIDCtx(context.TODO())
	return &defaultContext{Context: p, id: id}
}

// Wrap 包装成一个 ctx.Context
func Wrap(c context.Context) Context {
	if cc, ok := c.(Context); ok {
		return cc
	}

	p, id := xReqIDCtx(c)
	locale := p.Value(LocaleKey)
	if locale == nil {
		locale = "zh"
	}
	lang := p.Value(LangKey)
	if lang == nil {
		lang = language.SimplifiedChinese
	}
	tokenID := ""
	if t, ok := p.Value(TokenIDKey).(string); ok {
		tokenID = t
	}
	authInfo := p.Value(AuthInfoKey)
	return &defaultContext{Context: p, id: id, tokenID: tokenID, authInfo: authInfo}
}

// Request 根据 id 生成 ctx.Context
func Request(id string) Context {
	return &defaultContext{Context: context.TODO(), id: id}
}

func xReqIDCtx(c context.Context) (context.Context, string) {
	if c == nil {
		panic("xReqIDCtx：context.Context 不能传入 nil")
	}
	val := c.Value(RequestIDKey)
	if id, ok := val.(string); ok && id != "" {
		return c, id
	}
	id := uuid.NewV4().String()
	return context.WithValue(c, RequestIDKey, id), id
}
