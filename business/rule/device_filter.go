package rule

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/core/model"
	"log"
)

// DeviceFilterRule 设备过滤规则（仅允许android/ios设备）
type DeviceFilterRule struct {
	ruleName           string
	allowedDeviceTypes []string
}

// NewDeviceFilterRule 创建设备过滤规则实例
func NewDeviceFilterRule() _interface.MediaRequestRule {
	return &DeviceFilterRule{
		ruleName:           "device_filter",
		allowedDeviceTypes: []string{"android", "ios"},
	}
}

// Filter 执行设备过滤
func (d *DeviceFilterRule) Filter(adReq *model.AdMediaContent) (bool, error) {
	if adReq.DeviceID == "" {
		log.Printf("设备ID为空，过滤请求")
		return false, nil
	}

	// 检查设备类型是否在允许列表中
	for _, allowedType := range d.allowedDeviceTypes {
		if adReq.DeviceType == allowedType {
			return true, nil
		}
	}

	log.Printf("设备类型[%s]不被允许，过滤请求", adReq.DeviceType)
	return false, nil
}

// GetRuleName 获取规则名称
func (d *DeviceFilterRule) GetRuleName() string {
	return d.ruleName
}
