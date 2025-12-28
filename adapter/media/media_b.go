package media

import (
	"ad-exchange-server/core/model"
	"encoding/json"
	"net/http"
	"time"
)

// MediaB 原始请求格式（外部协议：字段名与MediaA不一致）
type MediaBRequest struct {
	MediaIdentifier string            `json:"media_identifier"`
	DeviceUUID      string            `json:"device_uuid"`
	DeviceCategory  string            `json:"device_category"`
	UserIdentifier  string            `json:"user_identifier"`
	AdSlotId        string            `json:"ad_slot_id"`
	Extra           map[string]string `json:"extra"`
}

// MediaB 原始响应格式（外部协议：字段名与MediaA不一致）
type MediaBResponse struct {
	AdvertisementId string `json:"advertisement_id"`
	AdName          string `json:"ad_name"`
	AdDesc          string `json:"ad_description"`
	ClickUrl        string `json:"click_url"`
	IsOk            bool   `json:"is_ok"`
}

// MediaBAdapter 媒体B适配器
type MediaBAdapter struct {
	mediaType string
}

// NewMediaBAdapter 创建媒体B适配器实例
func NewMediaBAdapter() *MediaBAdapter {
	return &MediaBAdapter{
		mediaType: "media_b",
	}
}

// UnmarshalRequest 媒体B请求 -> 内部统一请求
func (m *MediaBAdapter) UnmarshalRequest(r *http.Request) (*model.AdInternalRequest, error) {
	var mediaBReq MediaBRequest
	if err := json.NewDecoder(r.Body).Decode(&mediaBReq); err != nil {
		return nil, err
	}

	return &model.AdInternalRequest{
		MediaID:     mediaBReq.MediaIdentifier,
		DeviceID:    mediaBReq.DeviceUUID,
		DeviceType:  mediaBReq.DeviceCategory,
		UserID:      mediaBReq.UserIdentifier,
		RequestTime: time.Now().Unix(),
		AdSpaceID:   mediaBReq.AdSlotId,
		ExtInfo:     mediaBReq.Extra,
	}, nil
}

// MarshalResponse 内部统一响应 -> 媒体B响应
func (m *MediaBAdapter) MarshalResponse(internalResp *model.AdInternalResponse) ([]byte, error) {
	mediaBResp := MediaBResponse{
		AdvertisementId: internalResp.AdID,
		AdName:          internalResp.AdTitle,
		AdDesc:          internalResp.AdContent,
		ClickUrl:        internalResp.RedirectURL,
		IsOk:            internalResp.IsSuccess,
	}

	return json.Marshal(mediaBResp)
}

// GetMediaType 获取媒体类型
func (m *MediaBAdapter) GetMediaType() string {
	return m.mediaType
}
