package middleware

import (
	"ad-exchange-server/business/selection"
	"ad-exchange-server/config"
	"ad-exchange-server/core/model"
)

func DispatchAdSelectMiddleware() PlatformHandlerFunc {

	return func(platformContent *model.AdPlatformContent) bool {
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
		selectStrategy.Select(platformContent)
		return true
	}

}
