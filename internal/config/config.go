package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Appname    string
	MySQL      MySQLConfig
	StaticPath PathConfig
	MsgChannelType MsgChannelType
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

// 文件地址
type PathConfig struct {
	FilePath string
}

// 消息队列类型及相关信息
// gochannel为单机使用go默认的channel进行消息传递
// kafka 是使用 kafka 作为消息队列，可以分布式扩展消息聊天程序
type MsgChannelType struct {
	ChannelType string
	KafkaHosts  string
	KafkaTopic  string
}

var c Config

func init() {
	viper.SetConfigName("configs")    // 设置要读取配置文件的名称
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
