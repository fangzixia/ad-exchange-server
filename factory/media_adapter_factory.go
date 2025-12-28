package factory

import (
	"ad-exchange-server/adapter/media"
	"ad-exchange-server/adapter/media/adlink"
	_interface "ad-exchange-server/core/interface"
)

// CreateMediaAdapter 根据媒体类型创建对应的媒体适配器
func CreateMediaAdapter(mediaType string) _interface.MediaAdapter {
	switch mediaType {
	case "media_a":
		return media.NewMediaAAdapter()
	case "media_b":
		return media.NewMediaBAdapter()
	case "adlink":
		return adlink.NewAdapter()
	default:
		return nil
	}
}
