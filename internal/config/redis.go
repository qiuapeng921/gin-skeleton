package config

import "time"

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
