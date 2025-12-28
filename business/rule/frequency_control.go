package rule

import (
	_interface "ad-exchange-server/core/interface"
	"ad-exchange-server/core/model"
	"ad-exchange-server/infra/cache"
	"log"
	"time"
)

// FrequencyControlRule 设备频控规则（同一设备10秒内最多5次请求）
type FrequencyControlRule struct {
	ruleName string
	maxCount int
	expire   time.Duration
	cache    cache.MemoryCache
}

// NewFrequencyControlRule 创建设备频控规则实例
func NewFrequencyControlRule() _interface.MediaRequestRule {
	return &FrequencyControlRule{
		ruleName: "frequency_control",
		maxCount: 5, // 10秒内最多5次
		expire:   10 * time.Second,
		cache:    cache.NewMemoryCache(),
	}
}

// Filter 执行设备频控
func (f *FrequencyControlRule) Filter(adReq *model.AdInternalRequest) (bool, error) {
	cacheKey := "freq_control_" + adReq.DeviceID
	// 获取当前设备请求次数
	count, err := f.cache.GetInt(cacheKey)
	if err != nil && err != cache.ErrKeyNotFound {
		return false, err
	}

	// 检查是否超过最大次数
	if count >= f.maxCount {
		log.Printf("设备[%s]10秒内请求次数[%d]超过最大限制[%d]，过滤请求", adReq.DeviceID, count, f.maxCount)
		return false, nil
	}

	// 更新缓存次数
	if err := f.cache.Set(cacheKey, count+1, f.expire); err != nil {
		return false, err
	}

	return true, nil
}

// GetRuleName 获取规则名称
func (f *FrequencyControlRule) GetRuleName() string {
	return f.ruleName
}
