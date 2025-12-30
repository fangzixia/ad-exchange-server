package middleware

import (
	"ad-exchange-server/business/dispatch"
	"ad-exchange-server/core/model"
	"ad-exchange-server/factory"
)

func DispatchTrafficMiddleware() PlatformHandlerFunc {

	return func(platformContent *model.AdPlatformContent) bool {
		// 6. 获取所有平台方适配器
		platformAdapters, err := factory.GetAllPlatformAdapters()
		if err != nil {
			platformContent.PlatformStatus = 100
		}
		// 7. 并发分发请求给平台方
		dispatcher := dispatch.NewPlatformDispatcher()
		dispatcher.Dispatch(platformContent, platformAdapters)
		return true
	}

}
