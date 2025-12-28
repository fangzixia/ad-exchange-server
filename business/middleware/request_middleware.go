package middleware

import (
	"ad-exchange-server/core/model"
)

type RequestMiddleware func(req *model.AdMediaContent, next func()) bool

type RequestMiddlewareChain struct {
	middlewares []RequestMiddleware // 规则中间件列表（按执行顺序存储）
}

func NewRequestMiddlewareChain() *RequestMiddlewareChain {
	return &RequestMiddlewareChain{
		middlewares: make([]RequestMiddleware, 0),
	}
}

func (rc *RequestMiddlewareChain) Use(middlewares ...RequestMiddleware) {
	rc.middlewares = append(rc.middlewares, middlewares...)
}

func (rc *RequestMiddlewareChain) Execute(amc *model.AdMediaContent) {
	index := -1
	var next func()
	next = func() {
		index++
		// 递归终止条件：所有中间件执行完毕（规则全部通过）
		if index >= len(rc.middlewares) {
			return
		}
		// 获取当前要执行的规则中间件
		mw := rc.middlewares[index]
		// 执行当前规则中间件，传入 next 作为回调（当前规则通过后，调用 next 执行下一个）
		mw(amc, func() {
			// 此处空函数：仅作为“是否继续执行”的标识，实际递归逻辑在 next() 中
			// 若当前规则通过，会触发 next() 递归执行下一个规则
		})
	}
	next()
}
