package app

import (
	"micro-base/internal/pkg/core/ctx"
)

type Ctx ctx.Context

type Int interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Types interface {
	~string | Int | ~bool | Float
}
