package media

import "ad-exchange-server/core/model"

type MediaRequestLogFilter struct {
	next ChainFilter
}

func (f *MediaRequestLogFilter) Handle(*model.AdMediaContent) {

}

func (f *MediaRequestLogFilter) Next(chainFilter ChainFilter) {
	f.next = chainFilter
}
