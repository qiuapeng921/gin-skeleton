package config

import "time"

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
