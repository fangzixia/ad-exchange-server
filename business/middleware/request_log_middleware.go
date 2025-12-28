package middleware

import "ad-exchange-server/core/model"

func RequestLogMiddleware() RequestMiddleware {
	return func(req *model.AdMediaContent, next func(*model.AdMediaContent)) bool {
		return next(req)
	}
}
