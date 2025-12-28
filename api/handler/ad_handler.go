package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"ad-exchange-server/business/dispatch"
	"ad-exchange-server/business/rule"
	"ad-exchange-server/business/selection"
	"ad-exchange-server/config"
	"ad-exchange-server/core/model"
	"ad-exchange-server/factory"
)

// AdRequestHandler 广告请求处理器
func AdRequestHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 获取媒体类型和渠道号
	vars := mux.Vars(r)
	mediaType := vars["mediaType"]
	channelId := vars["channelId"]
	channel, err := strconv.Atoi(channelId)
	if err != nil {
		http.Error(w, fmt.Sprintf("channel is illegal:%s", channelId), http.StatusBadRequest)
	}

	// 2. 创建媒体适配器
	mediaAdapter := factory.CreateMediaAdapter(mediaType)
	if mediaAdapter == nil {
		http.Error(w, fmt.Sprintf("un support media_type:%s", mediaType), http.StatusBadRequest)
		return
	}

	// 3. 媒体请求 -> 内部统一请求
	internalReq, err := mediaAdapter.UnmarshalRequest(r)
	if err != nil {
		log.Printf("媒体请求反序列化失败: %v", err)
		http.Error(w, "request unmarshal failed", http.StatusBadRequest)
		return
	}

	// 4. 创建 content
	mediaContent := model.CreateMediaContent(internalReq, channel)

	// 4. 初始化规则链并执行过滤
	rules, err := factory.CreateAllRules()
	if err != nil {
		log.Printf("创建规则失败: %v", err)
		http.Error(w, "create rules failed", http.StatusInternalServerError)
		return
	}

	ruleChain := rule.NewRuleChain(rules)
	pass, err := ruleChain.Execute(mediaContent)
	if !pass || err != nil {
		log.Printf("规则过滤失败: %v", err)
		// 返回空响应
		emptyResp := &model.AdInternalResponse{IsSuccess: false}
		respBytes, _ := mediaAdapter.MarshalResponse(emptyResp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respBytes)
		return
	}

	// 5. 构建平台方内部请求
	platformInternalReq := &model.AdInternalRequest{
		DeviceID:   internalReq.DeviceID,
		DeviceType: internalReq.DeviceType,
		UserID:     internalReq.UserID,
		AdSpaceID:  internalReq.AdSpaceID,
		RequestID:  strconv.FormatInt(internalReq.RequestTime, 10),
	}

	// 6. 获取所有平台方适配器
	platformAdapters, err := factory.GetAllPlatformAdapters()
	if err != nil {
		log.Printf("获取平台方适配器失败: %v", err)
		http.Error(w, "get platform adapters failed", http.StatusInternalServerError)
		return
	}

	// 7. 并发分发请求给平台方
	dispatcher := dispatch.NewPlatformDispatcher()
	platformResponses := dispatcher.Dispatch(platformInternalReq, platformAdapters)

	// 8. 广告筛选
	var selectStrategy selection.AdSelectionStrategy
	strategyType := config.GetSelectionStrategy()
	switch strategyType {
	case "price_priority":
		selectStrategy = selection.NewPricePriorityStrategy()
	case "quality_priority":
		selectStrategy = selection.NewQualityPriorityStrategy()
	default:
		selectStrategy = selection.NewPricePriorityStrategy()
	}
	adInternalResp := selectStrategy.Select(platformResponses)

	// 9. 内部响应 -> 媒体响应
	respBytes, err := mediaAdapter.MarshalResponse(adInternalResp)
	if err != nil {
		log.Printf("媒体响应序列化失败: %v", err)
		http.Error(w, "response marshal failed", http.StatusInternalServerError)
		return
	}

	// 10. 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)

	//11. 记录日志
}
