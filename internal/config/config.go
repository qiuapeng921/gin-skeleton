package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const DefaultConfigFile = "app.yaml"

var CfgData *Config

func Init(configFile string) error {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &CfgData)
	return err
}

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
