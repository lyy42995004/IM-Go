package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Appname string
	MySQL   MySQLConfig
}

// MySQL 配置
type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

var c Config

func init() {
	viper.SetConfigName("config")    // 设置要读取配置文件的名称
	viper.SetConfigType("toml")      // 指定配置文件的类型为 TOML
	viper.AddConfigPath("./configs") // 添加配置文件的搜索路径
	viper.AutomaticEnv()             // 自动读取环境变量

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c) // 解析配置信息
}

func GetConfig() Config {
	return c
}
