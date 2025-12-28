package _interface

import "ad-exchange-server/core/model"

// MediaRequestRule 媒体请求过滤规则统一接口（策略模式）
type MediaRequestRule interface {
	// Filter 执行过滤逻辑
	Filter(adReq *model.AdMediaContent) (bool, error)
	// GetRuleName 获取规则名称
	GetRuleName() string
}
