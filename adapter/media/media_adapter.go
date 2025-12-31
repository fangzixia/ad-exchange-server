package media

import (
	"ad-exchange-server/adapter/media/adlink"
	"ad-exchange-server/adapter/media/hongyu"
	_interface "ad-exchange-server/core/interface"
)

// 确保后续实现满足接口约束
var (
	_ _interface.MediaAdapter = (*hongyu.Adapter)(nil)
	_ _interface.MediaAdapter = (*adlink.Adapter)(nil)
)
