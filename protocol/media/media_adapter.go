package media

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/protocol/media/adlink"
	"ad-exchange-server/protocol/media/hongyu"
)

// 确保后续实现满足接口约束
var (
	_ _interface.MediaAdapter = (*hongyu.Adapter)(nil)
	_ _interface.MediaAdapter = (*adlink.Adapter)(nil)
)
