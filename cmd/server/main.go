package main

import (
	"ad-exchange-server/bootstrap"
	"log"
	"net/http"
	"strconv"

	"ad-exchange-server/api"
	"ad-exchange-server/config"
)

func main() {
	// 1. 初始化基础设施
	bootstrap.Init()

	// 2. 初始化路由
	router := api.InitRouter()

	// 3. 获取服务端口
	port := config.GetServerPort()
	addr := ":" + strconv.Itoa(port)

	// 4. 启动HTTP服务
	log.Printf("广告网关服务启动，监听地址: %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
