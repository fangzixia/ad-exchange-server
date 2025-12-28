package model

// AdInternalRequest 内部统一广告请求模型（所有媒体协议最终转换为此格式）
type AdInternalRequest struct {
	RequestID   string            // 请求 ID
	MediaID     string            // 媒体ID
	DeviceID    string            // 设备ID
	DeviceType  string            // 设备类型（android/ios）
	UserID      string            // 用户ID
	RequestTime int64             // 请求时间戳
	AdSpaceID   string            // 广告位ID
	ExtInfo     map[string]string // 扩展信息
}
