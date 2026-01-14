package factory

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/protocol/media/adlink"
	"ad-exchange-server/protocol/media/hongyu"
)

var mediaAdapterMap = map[string]_interface.MediaAdapter{
	"adlink": adlink.NewAdapter(),
	"hongyu": hongyu.NewAdapter(),
}

// SelectMediaAdapter 根据媒体类型选择对应的媒体适配器
func SelectMediaAdapter(mediaType string) _interface.MediaAdapter {
	return mediaAdapterMap[mediaType]
}
