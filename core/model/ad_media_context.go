package model

import (
	"ad-exchange-server/pkg/uuid"
	"time"
)

type MediaStatus int

type MediaMessage string

const (
	MediaStatusInit MediaStatus = iota
	MediaStatusChannelErr
	MediaStatusMediaErr
	MediaStatusUnmarshalErr
)

type AdMediaContent struct {
	RequestID          string
	ChannelId          int
	MediaType          string
	RequestTime        time.Time
	AdInternalRequest  *AdInternalRequest
	AdInternalResponse *AdInternalResponse
	MediaStatus        MediaStatus
	MediaMessage       string
}

func CreateMediaContent() *AdMediaContent {
	return &AdMediaContent{
		RequestID:   uuid.GenerateUUID(),
		RequestTime: time.Now(),
		MediaStatus: MediaStatusInit,
	}

}
