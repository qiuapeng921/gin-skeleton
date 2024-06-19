package log

import (
	"github.com/rs/zerolog"
	"micro-base/internal/pkg/core/ctx"
)

type Logger interface {
	Init(cfg Config)
	Trace(c ctx.Context)
	Debug(c ctx.Context)
	Info(c ctx.Context)
	Warn(c ctx.Context)
	Error(c ctx.Context)
	Err(c ctx.Context, err error)
	Fatal(c ctx.Context)
	Panic(c ctx.Context)
	builder(event *zerolog.Event, c ctx.Context)
}
