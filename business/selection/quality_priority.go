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
func (q *QualityPriorityStrategy) Select(platformContent *model.AdPlatformContent) {
	platformResponses := platformContent.AdInternalResponses
	if len(platformResponses) == 0 {
		return
	}
	// 找出质量分最高的响应
	//maxQualityResp := platformResponses[0]
	//platformContent.FinalAdInternalResponses = maxQualityResp

}

// GetStrategyName 获取策略名称
func (q *QualityPriorityStrategy) GetStrategyName() string {
	return q.strategyName
}
