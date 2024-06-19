package errorx

import (
	"fmt"
	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

// 定义错误
var (
	ErrBadRequest      = NewErrorMsg(400, "请求发生错误")
	ErrNoPerm          = NewErrorMsg(401, "无访问权限")
	ErrInvalidToken    = NewErrorMsg(9999, "令牌失效")
	ErrNotFound        = NewErrorMsg(404, "资源不存在")
	ErrMethodNotAllow  = NewErrorMsg(405, "方法不被允许")
	ErrTooManyRequests = NewErrorMsg(429, "请求过于频繁")
	ErrInternalServer  = NewErrorMsg(500, "服务器发生错误")
)

// Error 定义业务错误
type Error struct {
	Code int   `json:"code"`          // 错误码
	Err  error `json:"err,omitempty"` // 响应错误
}

//NewError 新建
func NewError(code int, err error) *Error {
	return &Error{Code: code, Err: err}
}

//NewErrorMsg 新建
func NewErrorMsg(code int, msg string) *Error {
	return &Error{Code: code, Err: New(msg)}
}

//Error 错误输出
func (r *Error) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}
	return fmt.Sprintf("[%d]未定义错误码", r.Code)
}

//String 错误输出
func (r *Error) String() string {
	if r.Err != nil {
		return fmt.Sprintf("[%d]%s", r.Code, r.Err.Error())
	}
	return fmt.Sprintf("[%d]未定义错误码", r.Code)
}
