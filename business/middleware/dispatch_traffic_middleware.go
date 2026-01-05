package middleware

import (
	"ad-exchange-server/business/dispatch"
	"ad-exchange-server/core/model"
	"ad-exchange-server/factory"
)

func DispatchTrafficMiddleware() PlatformHandlerFunc {

	return func(platformContent *model.AdPlatformContent) bool {
		// 6. 获取所有平台方适配器
		platformAdapters := factory.GetAllPlatformAdapters()
		if platformAdapters == nil || len(platformAdapters) == 0 {
			platformContent.PlatformStatus = 100
			return false
		}
		// 7. 并发分发请求给平台方
		dispatcher := dispatch.NewPlatformDispatcher()
		internalResponses := dispatcher.Dispatch(platformContent, platformAdapters)
		platformContent.AdInternalResponses = internalResponses
		return true
	}

}
