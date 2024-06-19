package config

import "time"

// Cors 跨域配置
type Cors struct {
	Enable           bool          `yaml:"enable"`
	AllowAllOrigins  bool          `yaml:"allowAllOrigins"`
	AllowOrigins     []string      `yaml:"allowOrigins"`
	AllowMethods     []string      `yaml:"allowMethods"`
	AllowHeaders     []string      `yaml:"allowHeaders"`
	AllowCredentials bool          `yaml:"allowCredentials"`
	ExposeHeaders    []string      `yaml:"exposeHeaders"`
	MaxAge           time.Duration `yaml:"maxAge"`
	AllowWebSockets  bool          `yaml:"allowWebSockets"`
	AllowFiles       bool          `yaml:"allowFiles"`
}
