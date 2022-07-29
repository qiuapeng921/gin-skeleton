package config

import (
	"os"
	"time"

	"gitlab-ce.k8s.tools.vchangyi.com/common/go-toolbox/log"
)

//Config 配置组合
type Config struct {
	App         string `yaml:"app"`
	Mode        string `yaml:"mode"`
	Env         string `yaml:"env"`
	QYNoticeKey string `yaml:"qyNoticeKey"`

	Logger  LoggerCfg  `yaml:"logger"`
	Restful RestfulCfg `yaml:"restful"`
	Monitor MonitorCfg `yaml:"monitor"`
}

type JobCfg struct {
	SubCmd     string `yaml:"cmd" usage:"执行子命令"`
	Spec       string `yaml:"spec" usage:"可运行的特定标识。如果为空表示运行所有"`
	DelCmdSpec string `yaml:"delCmdSpec" usage:"删除灰度企业缓存"`
}

//Log 返回日志配置对象
func (c Config) Log(logToStderr bool) log.Config {
	if logToStderr {
		return log.Config{
			App:    c.App,
			Level:  c.Logger.Level,
			Format: c.Logger.Format,
			Output: os.Stderr,
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

//LoggerCfg 日志配置
type LoggerCfg struct {
	Level      string `yaml:"level" usage:"日志级别"`
	TargetType string `yaml:"targetType" usage:"输出目标类型【console/file/tcp/udp】"`
	Target     string `yaml:"target" usage:"日志输出格式【<file path>/<network address>】"`
	Format     string `yaml:"format" usage:"日志输出格式【yaml/raw】"`
}

//Cors 跨域配置
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

//Doc swagger 文档配置
type Doc struct {
	Title         string `yaml:"title"`
	Description   string `yaml:"description"`
	TermOfService string `yaml:"termOfService"`
	Contact       string `yaml:"contact"`
}

// RestfulCfg api 接口配置
type RestfulCfg struct {
	Host     string `yaml:"host"`
	BasePath string `yaml:"basePath"`
	Addr     string `yaml:"addr"`
	Cors     Cors   `yaml:"cors"`
	Doc      Doc    `yaml:"doc"`
}

//MonitorCfg 监控配置
type MonitorCfg struct {
	Enable bool   `yaml:"enable"`
	Addr   string `yaml:"addr"`
}
