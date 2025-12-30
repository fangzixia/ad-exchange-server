package middleware

import "ad-exchange-server/core/model"

func RequestLogMiddleware() MediaHandlerFunc {

	return func(mediaContent *model.AdMediaContent) bool {
		return true
	}

}
