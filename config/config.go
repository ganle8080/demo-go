package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host string
	Port int
}

var config *Config

func InitSysConfig() *Config {
	// 获取环境变量，如果没有指明是prod还是dev，默认是dev
	env := os.Getenv("ENV")
	fmt.Printf("env: %v\n", env)
	return nil
}

func Get() *Config {
	return config
}
