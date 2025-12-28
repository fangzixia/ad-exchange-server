package media

import (
	"ad-exchange-server/adapter/media/adlink"
	_interface "ad-exchange-server/core/interface"
)

// 确保后续实现满足接口约束
var (
	_ _interface.MediaAdapter = (*MediaAAdapter)(nil)
	_ _interface.MediaAdapter = (*MediaBAdapter)(nil)
	_ _interface.MediaAdapter = (*adlink.Adapter)(nil)
)
