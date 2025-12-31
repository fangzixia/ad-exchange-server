package hongyu

import (
	"ad-exchange-server/core/model"
	"encoding/json"
	"net/http"
)

type Adapter struct {
	mediaType string
}

// NewAdapter 创建媒体A适配器实例
func NewAdapter() *Adapter {
	return &Adapter{
		mediaType: "ad-link",
	}
}

// UnmarshalRequest 媒体A请求 -> 内部统一请求
func (m *Adapter) UnmarshalRequest(r *http.Request) (*model.AdInternalRequest, error) {
	var mediaAReq AdRequest
	if err := json.NewDecoder(r.Body).Decode(&mediaAReq); err != nil {
		return nil, err
	}

	return &model.AdInternalRequest{}, nil
}

// MarshalResponse 内部统一响应 -> 媒体A响应
func (m *Adapter) MarshalResponse(internalResp *model.AdInternalResponse) ([]byte, error) {
	mediaAResp := AdResponse{}

	return json.Marshal(mediaAResp)
}

// GetMediaType 获取媒体类型
func (m *Adapter) GetMediaType() string {
	return m.mediaType
}
