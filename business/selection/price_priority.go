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
func (p *PricePriorityStrategy) Select(platformResponses []*model.AdInternalResponse) *model.AdInternalResponse {
	if len(platformResponses) == 0 {
		return &model.AdInternalResponse{IsSuccess: false}
	}

	// 找出价格最高的响应
	maxPriceResp := platformResponses[0]
	for _, resp := range platformResponses {
		if resp.Price > maxPriceResp.Price {
			maxPriceResp = resp
		}
	}

	// 转换为内部统一响应
	return &model.AdInternalResponse{
		AdID:        maxPriceResp.AdID,
		AdTitle:     maxPriceResp.AdTitle,
		AdContent:   maxPriceResp.AdContent,
		RedirectURL: "http://localhost:8080/redirect/" + maxPriceResp.AdID,
		Price:       maxPriceResp.Price,
		IsSuccess:   true,
	}
}

// GetStrategyName 获取策略名称
func (p *PricePriorityStrategy) GetStrategyName() string {
	return p.strategyName
}
