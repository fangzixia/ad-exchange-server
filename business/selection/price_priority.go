package selection

import (
	"ad-exchange-server/core/model"
)

// PricePriorityStrategy 价格优先筛选策略
type PricePriorityStrategy struct {
	strategyName string
}

// NewPricePriorityStrategy 创建价格优先策略实例
func NewPricePriorityStrategy() AdSelectionStrategy {
	return &PricePriorityStrategy{
		strategyName: "price_priority",
	}
}

// Select 筛选出价最高的广告
func (p *PricePriorityStrategy) Select(platformContent *model.AdPlatformContent) {
	platformResponses := platformContent.AdInternalResponses
	if len(platformResponses) == 0 {
		return
	}

	// 找出价格最高的响应
	//maxPriceResp := platformResponses[0]
	//for _, resp := range platformResponses {
	//	if resp.Price > maxPriceResp.Price {
	//		maxPriceResp = resp
	//	}
	//}
	//platformContent.FinalAdInternalResponses = maxPriceResp
}

// GetStrategyName 获取策略名称
func (p *PricePriorityStrategy) GetStrategyName() string {
	return p.strategyName
}
