package titanvol

import (
	"ad-exchange-server/core/model"
	"encoding/json"
)

// Adapter 钛量平台适配器
type Adapter struct {
	platformName string
	platformURL  string
}

// NewAdapter 创建平台适配器实例
func NewAdapter() *Adapter {
	return &Adapter{
		platformName: "titanvol",
		platformURL:  "http://ca.adx.pw/media_a/hongyu?mid=21033", // 模拟地址
	}
}

// MarshalRequest 内部平台方请求 -> PlatformY请求
func (b *Adapter) MarshalRequest(internalReq *model.AdInternalRequest) ([]byte, error) {
	titanvolReq := AdRequest{}

	return json.Marshal(titanvolReq)
}

// UnmarshalResponse PlatformY响应 -> 内部统一响应
func (b *Adapter) UnmarshalResponse(respBytes []byte) (*model.AdInternalResponse, error) {
	var titanvolResp AdResponse
	if err := json.Unmarshal(respBytes, &titanvolResp); err != nil {
		return nil, err
	}

	return &model.AdInternalResponse{
		AdInfos: make([]model.AdInfo, 1),
	}, nil
}

// GetPlatformName 获取平台方名称
func (b *Adapter) GetPlatformName() string {
	return b.platformName
}

// GetPlatformURL 获取平台方接口地址
func (b *Adapter) GetPlatformURL() string {
	return b.platformURL
}
