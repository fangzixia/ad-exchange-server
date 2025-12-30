package middleware

import (
	"ad-exchange-server/core/model"
)

func RequestLDispatchMiddleware() MediaHandlerFunc {

	return func(mediaContent *model.AdMediaContent) bool {
		internalReq := mediaContent.AdInternalRequest
		platformContent := model.CreatePlatformContent(internalReq, mediaContent.ChannelId)

		platformHandleMiddleware := NewPlatformHandMiddleware()
		platformHandleMiddleware.Use(DispatchTrafficMiddleware())
		platformHandleMiddleware.Use(DispatchAdSelectMiddleware())
		platformHandleMiddleware.Do(platformContent)

		return true
	}

}
