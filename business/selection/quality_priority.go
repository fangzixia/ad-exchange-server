package selection

import (
	"ad-exchange-server/core/model"
)

// QualityPriorityStrategy 质量优先筛选策略
type QualityPriorityStrategy struct {
	strategyName string
}

// NewQualityPriorityStrategy 创建质量优先策略实例
func NewQualityPriorityStrategy() AdSelectionStrategy {
	return &QualityPriorityStrategy{
		strategyName: "quality_priority",
	}
}

// Select 筛选质量分最高的广告
func (q *QualityPriorityStrategy) Select(platformResponses []*model.AdInternalResponse) *model.AdInternalResponse {
	if len(platformResponses) == 0 {
		return &model.AdInternalResponse{IsSuccess: false}
	}

	// 找出质量分最高的响应
	maxQualityResp := platformResponses[0]
	//for _, resp := range platformResponses {
	//	if resp.QualityScore > maxQualityResp.QualityScore {
	//		maxQualityResp = resp
	//	}
	//}

	// 转换为内部统一响应
	return &model.AdInternalResponse{
		AdID:        maxQualityResp.AdID,
		AdTitle:     maxQualityResp.AdTitle,
		AdContent:   maxQualityResp.AdContent,
		RedirectURL: "http://localhost:8080/redirect/" + maxQualityResp.AdID,
		Price:       maxQualityResp.Price,
		IsSuccess:   true,
	}
}

// GetStrategyName 获取策略名称
func (q *QualityPriorityStrategy) GetStrategyName() string {
	return q.strategyName
}
