package selection

import "ad-exchange-server/core/model"

// AdSelectionStrategy 广告筛选策略统一接口
type AdSelectionStrategy interface {
	// Select 从平台方响应列表中筛选最优广告
	Select(platformResponses *model.AdPlatformContent)
	// GetStrategyName 获取策略名称
	GetStrategyName() string
}
