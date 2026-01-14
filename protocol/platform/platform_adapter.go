package platform

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/protocol/platform/titanvol"
)

// 确保后续实现满足接口约束
var (
	_ _interface.PlatformAdapter = (*PlatformXAdapter)(nil)
	_ _interface.PlatformAdapter = (*PlatformYAdapter)(nil)
	_ _interface.PlatformAdapter = (*titanvol.Adapter)(nil)
)
