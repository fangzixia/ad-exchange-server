package middleware

import "ad-exchange-server/core/model"

func RequestLogMiddleware() MediaHandlerFunc {

	return func(mediaContent *model.AdMediaContent) bool {
		defer func() {
			println("RequestLogMiddleware")
		}()
		return true
	}

}
