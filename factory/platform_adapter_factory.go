package factory

import (
	"ad-exchange-server/adapter/platform"
	_interface "ad-exchange-server/core/interface"
	"errors"
)

// CreatePlatformAdapter 根据平台方名称创建对应的平台方适配器
func CreatePlatformAdapter(PlatformName string) (_interface.PlatformAdapter, error) {
	switch PlatformName {
	case "platform_x":
		return platform.NewPlatformXAdapter(), nil
	case "platform_y":
		return platform.NewPlatformYAdapter(), nil
	default:
		return nil, errors.New("unsupported Platform name: " + PlatformName)
	}
}

// GetAllPlatformAdapters 获取所有支持的平台方适配器
func GetAllPlatformAdapters() ([]_interface.PlatformAdapter, error) {
	var adapters []_interface.PlatformAdapter
	// 模拟配置中配置的平台方列表
	platformNames := []string{"platform_x", "platform_y"}
	for _, name := range platformNames {
		adapter, err := CreatePlatformAdapter(name)
		if err != nil {
			return nil, err
		}
		adapters = append(adapters, adapter)
	}
	return adapters, nil
}
