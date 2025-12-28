package factory

import (
	"ad-exchange-server/business/rule"
	_interface "ad-exchange-server/core/interface"
)

// CreateAllRules 创建所有规则实例（设备过滤、请求控量、频控）
func CreateAllRules() ([]_interface.MediaRequestRule, error) {
	var rules []_interface.MediaRequestRule
	// 按执行顺序创建规则
	rules = append(rules, rule.NewDeviceFilterRule())
	rules = append(rules, rule.NewRequestControlRule())
	rules = append(rules, rule.NewFrequencyControlRule())
	return rules, nil
}
