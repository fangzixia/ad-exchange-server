package rule

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/core/model"
	"log"
	"sync/atomic"
)

// RequestControlRule 请求控量规则（每秒最大100个请求）
type RequestControlRule struct {
	ruleName string
	maxQPS   int32
	curQPS   int32
}

// NewRequestControlRule 创建请求控量规则实例
func NewRequestControlRule() _interface.MediaRequestRule {
	return &RequestControlRule{
		ruleName: "request_control",
		maxQPS:   100, // 模拟每秒最大100个请求
		curQPS:   0,
	}
}

// Filter 执行请求控量
func (r *RequestControlRule) Filter(adReq *model.AdInternalRequest) (bool, error) {
	// 原子操作更新当前QPS
	current := atomic.AddInt32(&r.curQPS, 1)
	defer atomic.AddInt32(&r.curQPS, -1) // 模拟每秒重置，此处简化为请求结束后减1

	if current > r.maxQPS {
		log.Printf("当前QPS[%d]超过最大限制[%d]，过滤请求", current, r.maxQPS)
		return false, nil
	}

	return true, nil
}

// GetRuleName 获取规则名称
func (r *RequestControlRule) GetRuleName() string {
	return r.ruleName
}
