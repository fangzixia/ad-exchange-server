package middleware

import (
	"ad-exchange-server/core/model"
)

type MediaHandlerFunc func(*model.AdMediaContent) bool

type MediaHandMiddleware struct {
	index       int
	middlewares []MediaHandlerFunc
}

func NewMediaHandMiddleware() *MediaHandMiddleware {
	return &MediaHandMiddleware{
		middlewares: make([]MediaHandlerFunc, 0),
		index:       -1,
	}
}

func (rmw *MediaHandMiddleware) Use(middleware ...MediaHandlerFunc) {
	rmw.middlewares = append(rmw.middlewares, middleware...)
}

func (rmw *MediaHandMiddleware) Do(amc *model.AdMediaContent) {
	rmw.index++
	for ; rmw.index < len(rmw.middlewares); rmw.index++ {
		b := rmw.middlewares[rmw.index](amc)
		if !b {
			break
		}
	}
}

func (rmw *MediaHandMiddleware) Reset() {
	rmw.index = -1
}
