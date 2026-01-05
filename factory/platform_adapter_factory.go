package factory

import (
	"ad-exchange-server/adapter/platform"
	"ad-exchange-server/adapter/platform/titanvol"
	_interface "ad-exchange-server/core/interface"
)

var platformAdapterMap = map[string]_interface.PlatformAdapter{
	"titanvol": titanvol.NewAdapter(),
	"x":        platform.NewPlatformXAdapter(),
}

// SelectPlatformAdapter 根据平台方名称创建对应的平台方适配器
func SelectPlatformAdapter(PlatformName string) _interface.PlatformAdapter {
	return platformAdapterMap[PlatformName]

}

// GetAllPlatformAdapters 获取所有支持的平台方适配器
func GetAllPlatformAdapters() []_interface.PlatformAdapter {
	var adapters []_interface.PlatformAdapter
	// 模拟配置中配置的平台方列表
	platformNames := []string{"titanvol", "y"}
	for _, name := range platformNames {
		adapter := SelectPlatformAdapter(name)
		if adapter != nil {
			adapters = append(adapters, adapter)
		}
	}
	return adapters
}
