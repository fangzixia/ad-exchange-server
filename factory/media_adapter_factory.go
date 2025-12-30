package factory

import (
	"ad-exchange-server/adapter/media"
	"ad-exchange-server/adapter/media/adlink"
	_interface "ad-exchange-server/core/interface"
)

var mediaAdapterMap = map[string]_interface.MediaAdapter{
	"adlink":  adlink.NewAdapter(),
	"media_a": media.NewMediaAAdapter(),
	"media_b": media.NewMediaBAdapter(),
}

// SelectMediaAdapter 根据媒体类型创建对应的媒体适配器
func SelectMediaAdapter(mediaType string) _interface.MediaAdapter {
	return mediaAdapterMap[mediaType]
}
