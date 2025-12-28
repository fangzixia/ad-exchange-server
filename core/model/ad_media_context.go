package model

import (
	"ad-exchange-server/pkg/uuid"
	"time"
)

type MediaStatus int

type MediaMessage string

var MediaStatusInit MediaStatus = 0

type AdMediaContent struct {
	RequestID          string
	ChannelId          int
	RequestTime        time.Time
	AdInternalRequest  *AdInternalRequest
	AdInternalResponse *AdInternalResponse
	MediaStatus        MediaStatus
	MediaMessage       string
}

func CreateMediaContent(request *AdInternalRequest, channelId int) *AdMediaContent {
	return &AdMediaContent{
		RequestID:         uuid.GenerateUUID(),
		RequestTime:       time.Now(),
		ChannelId:         channelId,
		AdInternalRequest: request,
		MediaStatus:       MediaStatusInit,
	}

}
