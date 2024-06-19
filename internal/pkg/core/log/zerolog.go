package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"micro-base/internal/pkg/core/ctx"
	"os"
	"strings"
)

var logger zerolog.Logger

// Config 日志组件配置参数
type Config struct {
	App        string    `yaml:"app" json:"app" usage:"日志模块名称"`
	Level      string    `json:"level" usage:"日志级别"`
	Format     string    `json:"format" usage:"日志输出格式:json/raw"`
	TargetType string    `json:"target_type"  usage:"输出目标类型:console/file"`
	Target     string    `json:"target" usage:"日志输出格式: file path"`
	Output     io.Writer `json:"-" yaml:"-" toml:"-"`
}

// Init 初始化日志组件
func Init(cfg Config) error {
	level, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		log.Warn().AnErr("err", err).Msg("日志级别配置错误，使用默认 debug 级别")
		level = zerolog.DebugLevel
	}
	cfg = mergeDefaultCfg(cfg)

	zerolog.SetGlobalLevel(level)
	logger = log.
		// 日志基本配置
		Output(cfg.Output).
		Level(level).
		// 日志默认输出字段配置
		With().
		Caller().
		Str("app", cfg.App).
		Logger()
	log.Logger = logger
	return err
}

func mergeDefaultCfg(cfg Config) Config {
	if cfg.Output == nil {
		var output io.Writer
		switch cfg.TargetType {
		case "file":
			output = NewFileOutput(cfg.Target)
		default:
			output = os.Stdout
		}

		cfg.Output = output
		if cfg.Format == "json" {
			cfg.Output = output
		} else {
			cfg.Output = zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) { w.Out = output })
		}
	}

	return cfg
}

// NewConsoleFileOutput console 格式输出到文件
func NewConsoleFileOutput(logfile string) io.Writer {
	return zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = NewFileOutput(logfile)
		w.NoColor = true
	})
}

// NewFileOutput json 格式输出到文件
func NewFileOutput(logfile string) io.Writer {
	lf, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	return lf
}

// Option zerolog.Event 配置
type Option func(*zerolog.Event)

// Trace 以 trace 级别开始一个新的日志
//
// 必须调用 Msg 方法才能输出日志
func Trace(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Trace(), c, options...)
}

// Debug 以 debug 级别开始一个新的日志
//
// 必须调用 Msg 方法才能输出日志
func Debug(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Debug(), c, options...)
}

// Info 以 info 级别开始一个新的日志
//
// 必须调用 Msg 方法才能输出日志
func Info(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Info(), c, options...)
}

// Warn 以 warn 级别开始一个新的日志
//
// 必须调用 Msg 方法才能输出日志
func Warn(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Warn(), c, options...)
}

// Error 以 error 级别开始一个新的日志
//
// 必须调用 Msg 方法才能输出日志
func Error(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Error(), c, options...)
}

// Err 以 error 或 info 级别开始一个新的日志。
// > 如果 err == nil，开启 info  级别
// > 如果 err != nil，开启 error 级别
//
// 必须调用 Msg 方法才能输出日志
func Err(c ctx.Context, err error, options ...Option) *zerolog.Event {
	return builder(logger.Err(err), c, options...)
}

// Fatal 以 fatal 级别开始一个新的日志。
// 日志输出过后会调用 os.Exit(1) 退出程序
//
// 必须调用 Msg 方法才能输出日志
func Fatal(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Fatal(), c, options...)
}

// Panic 以 panic 级别开始一个新的日志。
// 日志输出过后会调用 panic() 方法，抛出输出的消息
//
// 必须调用 Msg 方法才能输出日志
func Panic(c ctx.Context, options ...Option) *zerolog.Event {
	return builder(logger.Panic(), c, options...)
}

// builder 业务字段日志输出字段配置
func builder(event *zerolog.Event, c ctx.Context, options ...Option) *zerolog.Event {
	for _, option := range options {
		option(event)
	}
	return event.Str("x-request-id", c.ID())
}
