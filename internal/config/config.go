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
func Init() {
	yamlFile, err := ioutil.ReadFile(DefaultConfigFile)
	err = yaml.Unmarshal(yamlFile, &CfgData)
	if err != nil {
		fmt.Println("解析 YAML 失败：", err)
		return
	}
}
