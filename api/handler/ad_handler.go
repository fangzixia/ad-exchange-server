package handler

import (
	"ad-exchange-server/business/middleware"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"ad-exchange-server/core/model"
	"ad-exchange-server/factory"
)

// AdRequestHandler 广告请求处理器
func AdRequestHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
	}()

	// 1. 获取媒体类型和渠道号
	vars := mux.Vars(r)
	mediaType := vars["mediaType"]
	channelId := vars["channelId"]
	channel, err := strconv.Atoi(channelId)
	if err != nil {
		http.Error(w, fmt.Sprintf("channel is illegal:%s", channelId), http.StatusBadRequest)
	}

	// 2. 创建媒体适配器
	mediaAdapter := factory.SelectMediaAdapter(mediaType)
	if mediaAdapter == nil {
		http.Error(w, fmt.Sprintf("un support media_type:%s", mediaType), http.StatusBadRequest)
		return
	}

	// 4. 创建 content
	mediaContent := model.CreateMediaContent()
	mediaContent.ChannelId = channel

	// 3. 媒体请求 -> 内部统一请求
	interReq := mediaAdapter.UnmarshalRequest(r)
	if interReq == nil {
		mediaContent.MediaStatus = model.MediaStatusUnmarshalErr
		http.Error(w, "request unmarshal failed", http.StatusBadRequest)
		return
	}
	mediaContent.AdInternalRequest = interReq

	mediaHandMiddleware := middleware.NewMediaHandMiddleware()
	mediaHandMiddleware.Use(middleware.RequestLogMiddleware())
	mediaHandMiddleware.Use(middleware.RequestLDispatchMiddleware())
	mediaHandMiddleware.Do(mediaContent)

	// 9. 内部响应 -> 媒体响应
	respBytes, err := mediaAdapter.MarshalResponse(mediaContent.AdInternalResponse)
	if err != nil {
		log.Printf("媒体响应序列化失败: %v", err)
		http.Error(w, "response marshal failed", http.StatusInternalServerError)
		return
	}

	// 10. 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
