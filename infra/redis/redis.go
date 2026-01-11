package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client

func InitRedis() error {
	// 创建客户端实例，配置连接信息和连接池
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis 地址
		Password: "",               // Redis 密码，无则为空字符串
		DB:       0,                // 使用的数据库编号
		// 连接池配置（关键：优化性能，避免频繁创建连接）
		PoolSize:        10,               // 连接池最大活跃连接数，默认 10
		MinIdleConns:    5,                // 连接池最小空闲连接数，避免频繁创建/销毁连接
		ConnMaxIdleTime: 30 * time.Second, // 空闲连接超时时间
	})

	// 测试连接（Ping）
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 释放上下文资源
	_, err := rdb.Ping(ctx).Result()
	return err
}

func SetValue(key string, value any, duration time.Duration) error {
	return rdb.Set(context.Background(), key, value, duration).Err()
}

func GetValue(key string) (any, error) {
	return rdb.Get(context.Background(), key).Result()
}

func HSetValue(key string, hKey string, value any, duration time.Duration) error {
	return rdb.HSet(context.Background(), key, hKey, value, duration).Err()
}

func HGetValue(key string, hKey string) (any, error) {
	return rdb.HGet(context.Background(), key, hKey).Result()
}
