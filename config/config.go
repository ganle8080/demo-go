package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host string
	Port int
}

var config *Config

// InitSysConfig 读取配置文件，初始化config变量
func InitSysConfig() *Config {
	// 获取环境变量，如果没有指明是prod还是dev，默认是dev
	env := os.Getenv("ENV")
	if env == "" || env == "dev" {
		env = "dev"
		viper.SetConfigName("config" + "-" + env) // 环境特定配置文件名 (不带扩展名)
	} else {

	}

	// 加载环境特定的配置文件

	viper.SetConfigType("yaml") // 环境特定配置文件类型
	viper.AddConfigPath(".")    // 查找环境特定配置文件的路径

	return nil
}

// Get 读取config变量
func Get() *Config {
	return config
}
