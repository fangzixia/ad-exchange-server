package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("selection.strategy", "price_priority")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("读取配置文件失败，使用默认配置: %v", err)
	}
}

// GetServerPort 获取服务端口
func GetServerPort() int {
	return viper.GetInt("server.port")
}

// GetSelectionStrategy 获取广告筛选策略
func GetSelectionStrategy() string {
	return viper.GetString("selection.strategy")
}
