package config

import (
	"micro-base/internal/pkg/core/log"
	"time"
)

// Config 配置组合
type Config struct {
	App     string     `yaml:"app"`
	Mode    string     `yaml:"mode"`
	Env     string     `yaml:"env"`
	Restful RestfulCfg `yaml:"restful"`

	Logger LoggerCfg
}

// Log 返回日志配置对象
func (c Config) Log(logToStderr bool) log.Config {
	if logToStderr {
		return log.Config{
			App:    c.App,
			Level:  c.Logger.Level,
			Format: c.Logger.Format,
		}
	}
	return log.Config{
		App:        c.App,
		Level:      c.Logger.Level,
		TargetType: c.Logger.TargetType,
		Target:     c.Logger.Target,
		Format:     c.Logger.Format,
	}
}

// LoggerCfg 日志配置
type LoggerCfg struct {
	Level      string `yaml:"level" usage:"日志级别"`
	TargetType string `yaml:"targetType" usage:"输出目标类型 console/file"`
	Target     string `yaml:"target" usage:"日志输出格式 file path"`
	Format     string `yaml:"format" usage:"日志输出格式 yaml/raw"`
}

// Cors 跨域配置
type Cors struct {
	Enable           bool          `yaml:"enable"`
	AllowAllOrigins  bool          `yaml:"AllowAllOrigins"`
	AllowOrigins     []string      `yaml:"AllowOrigins"`
	AllowMethods     []string      `yaml:"AllowMethods"`
	AllowHeaders     []string      `yaml:"AllowHeaders"`
	AllowCredentials bool          `yaml:"AllowCredentials"`
	ExposeHeaders    []string      `yaml:"ExposeHeaders"`
	MaxAge           time.Duration `yaml:"MaxAge"`
	AllowWebSockets  bool          `yaml:"AllowWebSockets"`
	AllowFiles       bool          `yaml:"AllowFiles"`
}

// RestfulCfg api 接口配置
type RestfulCfg struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"basePath"`
	Addr     string `yaml:"addr"`
	Cors     Cors   `yaml:"cors"`
}
