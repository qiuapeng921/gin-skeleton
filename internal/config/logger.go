package config

import "micro-base/internal/pkg/core/log"

// Log 返回日志配置对象
func (c Config) Log() log.Config {
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
	Format     string `yaml:"format" usage:"日志输出格式 json/raw"`
}
