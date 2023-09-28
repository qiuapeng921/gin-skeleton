package errs

import (
	"micro-base/internal/pkg/core/errorx"
)

// ErrorStatus 错误状态码
var ErrorStatus = 401

// DefaultErrorCode 默认码
var DefaultErrorCode = 100000

// DefaultErrorMsg 默认错误信息
var DefaultErrorMsg = "参数错误"
var (
	// ParamsErr 参数错误
	ParamsErr = errorx.NewErrorMsg(DefaultErrorCode, DefaultErrorMsg)
)
