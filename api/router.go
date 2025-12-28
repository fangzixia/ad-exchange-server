package api

import (
	"ad-exchange-server/api/handler"
	"github.com/gorilla/mux"
)

// InitRouter 初始化路由
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	// 广告请求接口
	router.HandleFunc("/v2/adx/{mediaType}/ch/{channelId}", handler.AdRequestHandler)
	return router
}
