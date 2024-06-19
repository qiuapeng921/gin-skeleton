package ctx

import (
	"context"
	"github.com/google/uuid"
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

// Context 上下文包装类
//
// 1. 不要把Context放在结构体中，要以参数的方式传递
// 2. 以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
// 3. 给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用 `context.TODO_
// 4. Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
// 5. Context是线程安全的，可以放心的在多个goroutine中传递
type Context interface {
	context.Context
	ID() string
	Locale() string
	Language() language.Tag
	AuthInfo() interface{}
	SetAuthInfo(tokenId string, info interface{})
}

type defaultContext struct {
	context.Context
	id     string
	locale string
	lang   language.Tag
	// auth
	tokenID  string
	authInfo interface{}
}

func (d *defaultContext) AuthInfo() interface{} {
	return d.authInfo
}

func (d *defaultContext) SetAuthInfo(tokenID string, info interface{}) {
	d.tokenID = tokenID
	d.authInfo = info
}

func (d *defaultContext) ID() string {
	return d.id
}

func (d *defaultContext) Locale() string {
	return d.locale
}

func (d *defaultContext) Language() language.Tag {
	return d.lang
}

// New 新建 context
func New() Context {
	p, id := xReqIDCtx(context.TODO())
	return &defaultContext{Context: p, id: id, locale: "zh", lang: language.SimplifiedChinese}
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
	return &defaultContext{
		Context: p, id: id, locale: locale.(string), lang: lang.(language.Tag),
		tokenID: tokenID, authInfo: authInfo,
	}
}

// Request 根据 id 生成 ctx.Context
func Request(id string) Context {
	return &defaultContext{Context: context.TODO(), id: id, locale: "zh", lang: language.SimplifiedChinese}
}

func xReqIDCtx(c context.Context) (context.Context, string) {
	if c == nil {
		panic("xReqIDCtx：context.Context 不能传入 nil")
	}
	val := c.Value(RequestIDKey)
	if id, ok := val.(string); ok && id != "" {
		return c, id
	}
	id := uuid.New().String()
	return context.WithValue(c, RequestIDKey, id), id
}
