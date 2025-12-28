package media

import "ad-exchange-server/core/model"

type ChainFilter interface {
	DoFilter(c *model.AdMediaContent) bool
}

type FilterChain struct {
	filters []ChainFilter
}

func NewFilterChain() *FilterChain {
	return &FilterChain{
		filters: make([]ChainFilter, 0),
	}
}

func (f *FilterChain) AddFilter(filter ChainFilter) {
	f.filters = append(f.filters, filter)
}

func (f *FilterChain) Execute(c *model.AdMediaContent) {
	for _, filter := range f.filters {
		doNext := filter.DoFilter(c)
		if !doNext {
			break
		}
	}
}
