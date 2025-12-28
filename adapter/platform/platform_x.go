package platform

import (
	"ad-exchange-server/core/model"
	"encoding/json"
)

// PlatformX 原始请求格式
type PlatformXRequest struct {
	Device_UUID string `json:"device_uuid"`
	Device_Type string `json:"device_type"`
	User_ID     string `json:"user_id"`
	Ad_Space_ID string `json:"ad_space_id"`
	Request_ID  string `json:"request_id"`
}

// PlatformX 原始响应格式
type PlatformXResponse struct {
	Platform_Id string  `json:"platform_id"`
	Ad_Id       string  `json:"ad_id"`
	Ad_Title    string  `json:"ad_title"`
	Ad_Content  string  `json:"ad_content"`
	Price       float64 `json:"price"`
	Quality     int     `json:"quality_score"`
	Valid       bool    `json:"is_valid"`
}

// PlatformXAdapter 平台方X适配器
type PlatformXAdapter struct {
	platformName string
	platformURL  string
}

// NewPlatformXAdapter 创建平台方X适配器实例
func NewPlatformXAdapter() *PlatformXAdapter {
	return &PlatformXAdapter{
		platformName: "platform_x",
		platformURL:  "http://localhost:8081/platform/x/ad", // 模拟地址
	}
}

// MarshalRequest 内部平台方请求 -> PlatformX请求
func (b *PlatformXAdapter) MarshalRequest(internalReq *model.AdInternalRequest) ([]byte, error) {
	platformXReq := PlatformXRequest{
		Device_UUID: internalReq.DeviceID,
		Device_Type: internalReq.DeviceType,
		User_ID:     internalReq.UserID,
		Ad_Space_ID: internalReq.AdSpaceID,
	}

	return json.Marshal(platformXReq)
}

// UnmarshalResponse PlatformX响应 -> 内部统一响应
func (b *PlatformXAdapter) UnmarshalResponse(respBytes []byte) (*model.AdInternalResponse, error) {
	var platformXResp PlatformXResponse
	if err := json.Unmarshal(respBytes, &platformXResp); err != nil {
		return nil, err
	}

	return &model.AdInternalResponse{
		AdID:      platformXResp.Ad_Id,
		AdTitle:   platformXResp.Ad_Title,
		AdContent: platformXResp.Ad_Content,
		Price:     platformXResp.Price,
	}, nil
}

// GetPlatformName 获取平台方名称
func (b *PlatformXAdapter) GetPlatformName() string {
	return b.platformName
}

// GetPlatformURL 获取平台方接口地址
func (b *PlatformXAdapter) GetPlatformURL() string {
	return b.platformURL
}
