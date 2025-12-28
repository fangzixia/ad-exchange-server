package media

import (
	"ad-exchange-server/core/model"
	"encoding/json"
	"net/http"
	"time"
)

// MediaA 原始请求格式（外部协议）
type MediaARequest struct {
	MediaId    string `json:"media_id"`
	DeviceId   string `json:"device_id"`
	DeviceType string `json:"device_type"`
	UserId     string `json:"user_id"`
	AdSpaceId  string `json:"ad_space_id"`
}

// MediaA 原始响应格式（外部协议）
type MediaAResponse struct {
	AdId        string `json:"ad_id"`
	Title       string `json:"ad_title"`
	Content     string `json:"ad_content"`
	RedirectUrl string `json:"redirect_url"`
	Success     bool   `json:"success"`
}

// MediaAAdapter 媒体A适配器
type MediaAAdapter struct {
	mediaType string
}

// NewMediaAAdapter 创建媒体A适配器实例
func NewMediaAAdapter() *MediaAAdapter {
	return &MediaAAdapter{
		mediaType: "ad-link",
	}
}

// UnmarshalRequest 媒体A请求 -> 内部统一请求
func (m *MediaAAdapter) UnmarshalRequest(r *http.Request) (*model.AdInternalRequest, error) {
	var mediaAReq MediaARequest
	if err := json.NewDecoder(r.Body).Decode(&mediaAReq); err != nil {
		return nil, err
	}

	return &model.AdInternalRequest{
		MediaID:     mediaAReq.MediaId,
		DeviceID:    mediaAReq.DeviceId,
		DeviceType:  mediaAReq.DeviceType,
		UserID:      mediaAReq.UserId,
		RequestTime: time.Now().Unix(),
		AdSpaceID:   mediaAReq.AdSpaceId,
		ExtInfo:     make(map[string]string),
	}, nil
}

// MarshalResponse 内部统一响应 -> 媒体A响应
func (m *MediaAAdapter) MarshalResponse(internalResp *model.AdInternalResponse) ([]byte, error) {
	mediaAResp := MediaAResponse{
		AdId:        internalResp.AdID,
		Title:       internalResp.AdTitle,
		Content:     internalResp.AdContent,
		RedirectUrl: internalResp.RedirectURL,
		Success:     internalResp.IsSuccess,
	}

	return json.Marshal(mediaAResp)
}

// GetMediaType 获取媒体类型
func (m *MediaAAdapter) GetMediaType() string {
	return m.mediaType
}
