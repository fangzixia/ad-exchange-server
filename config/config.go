package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
}

func InitDB(cfg DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Params)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db, err := open.DB()
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(cfg.MaxOpenConnection)
	db.SetMaxIdleConns(cfg.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionMaxLifetime) * time.Second)
}

type DatabaseConfig struct {
	Driver                string `yaml:"driver"`
	Username              string `yaml:"username"`
	Password              string `yaml:"password"`
	Host                  string `yaml:"host"`
	Port                  int    `yaml:"port"`
	DbName                string `yaml:"dbName"`
	Params                string `yaml:"params"`
	MaxOpenConnection     int    `yaml:"max_open_connection"`
	MaxIdleConnection     int    `yaml:"max_idle_connection"`
	ConnectionMaxLifetime int    `yaml:"connection_max_lifetime"`
}

// InitConfig 初始化配置
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
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
