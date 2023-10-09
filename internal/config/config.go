package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// DefaultConfigFile ...
const DefaultConfigFile = "app.yaml"

// CfgData ...
var CfgData *Config

// Init 业务数据初始化，在解析命令行参数过后执行
func Init(configFile string) {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Errorf("读取配置文件失败: %s", err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &CfgData)
	if err != nil {
		fmt.Errorf("解析配置文件失败: %s", err.Error())
	}
}
