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
	Logger  LoggerCfg  `yaml:"logger"`
	DB      DbConfig   `yaml:"db"`
	RDB     RdbConfig  `yaml:"rdb"`
}

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

// RestfulCfg api 接口配置
type RestfulCfg struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"basePath"`
	Addr     string `yaml:"addr"`
	Cors     Cors   `yaml:"cors"`
}

type DbConfig struct {
	Driver  string            `yaml:"driver"`
	DnsList map[string]string `yaml:"dnsList"`
	Log     bool              `yaml:"log"`
	Pool    DbPoolCfg         `yaml:"pool"`
	Cluster string            `yaml:"cluster"`
}

type DbPoolCfg struct {
	MaxIdle     int           `yaml:"maxIdle"`
	MaxOpen     int           `yaml:"maxOpen"`
	MaxLifeTime time.Duration `yaml:"maxLifeTime" usage:"单位：秒"`
}
type RdbConfig struct {
	Addr        string        `yaml:"addr"`
	User        string        `yaml:"user"`
	Pwd         string        `yaml:"pwd"`
	DB          int           `yaml:"db"`
	DialTimeout time.Duration `yaml:"dialTimeout"`
	CmdTimeout  time.Duration `yaml:"cmdTimeout"`
	PoolSize    int           `yaml:"poolSize"`
	MinIdle     int           `yaml:"minIdle"`
}
