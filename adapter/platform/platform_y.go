package platform

import (
	"ad-exchange-server/core/model"
	"encoding/json"
)

// PlatformY 原始请求格式（字段名与PlatformX不一致）
type PlatformYRequest struct {
	DeviceId     string `json:"device_id"`
	DeviceCat    string `json:"device_cat"`
	UserUuid     string `json:"user_uuid"`
	AdPositionId string `json:"ad_position_id"`
	ReqId        string `json:"req_id"`
}

// PlatformY 原始响应格式（字段名与PlatformX不一致，出价字段为BidPrice）
type PlatformYResponse struct {
	PlatformId   string  `json:"platform_id"`
	AdvertId     string  `json:"advert_id"`
	Title        string  `json:"title"`
	Desc         string  `json:"description"`
	BidPrice     float64 `json:"bid_price"`
	QualityLevel int     `json:"quality_level"`
	Available    bool    `json:"available"`
}

// PlatformYAdapter 平台方Y适配器
type PlatformYAdapter struct {
	platformName string
	platformURL  string
}

// NewPlatformYAdapter 创建平台方Y适配器实例
func NewPlatformYAdapter() *PlatformYAdapter {
	return &PlatformYAdapter{
		platformName: "platform_y",
		platformURL:  "http://localhost:8081/platform/y/ad", // 模拟地址
	}
}

// MarshalRequest 内部平台方请求 -> PlatformY请求
func (b *PlatformYAdapter) MarshalRequest(internalReq *model.AdPlatformContent) ([]byte, error) {
	platformYReq := PlatformYRequest{}

	return json.Marshal(platformYReq)
}

// UnmarshalResponse PlatformY响应 -> 内部统一响应
func (b *PlatformYAdapter) UnmarshalResponse(respBytes []byte) (*model.AdInternalResponse, error) {
	var platformYResp PlatformYResponse
	if err := json.Unmarshal(respBytes, &platformYResp); err != nil {
		return nil, err
	}

	return &model.AdInternalResponse{}, nil
}

// GetPlatformName 获取平台方名称
func (b *PlatformYAdapter) GetPlatformName() string {
	return b.platformName
}

// GetPlatformURL 获取平台方接口地址
func (b *PlatformYAdapter) GetPlatformURL() string {
	return b.platformURL
}
