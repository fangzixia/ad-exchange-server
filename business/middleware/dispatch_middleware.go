package middleware

import (
	"ad-exchange-server/core/model"
)

type PlatformHandlerFunc func(*model.AdPlatformContent) bool

type PlatformHandMiddleware struct {
	index       int
	middlewares []PlatformHandlerFunc
}

func NewPlatformHandMiddleware() *PlatformHandMiddleware {
	return &PlatformHandMiddleware{
		middlewares: make([]PlatformHandlerFunc, 0),
		index:       -1,
	}
}

func (rmw *PlatformHandMiddleware) Use(middleware ...PlatformHandlerFunc) {
	rmw.middlewares = append(rmw.middlewares, middleware...)
}

func (rmw *PlatformHandMiddleware) Do(amc *model.AdPlatformContent) {
	rmw.index++
	for ; rmw.index < len(rmw.middlewares); rmw.index++ {
		b := rmw.middlewares[rmw.index](amc)
		if !b {
			break
		}
	}
}

func (rmw *PlatformHandMiddleware) Reset() {
	rmw.index = -1
}
