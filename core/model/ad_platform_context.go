package model

import (
	"ad-exchange-server/pkg/uuid"
	"time"
)

type PlatformStatus int

type PlatformMessage string

var PlatformStatusInit PlatformStatus = 0

type AdPlatformContent struct {
	RequestID                string
	ChannelId                int
	RequestTime              time.Time
	AdInternalRequest        *AdInternalRequest
	AdInternalResponses      map[string]*AdInternalResponse
	FinalAdInternalResponses *AdInternalResponse
	PlatformStatus           PlatformStatus
	PlatformMessage          string
}

func CreatePlatformContent(request *AdInternalRequest, channelId int) *AdPlatformContent {
	return &AdPlatformContent{
		RequestID:           uuid.GenerateUUID(),
		RequestTime:         time.Now(),
		ChannelId:           channelId,
		AdInternalRequest:   request,
		PlatformStatus:      PlatformStatusInit,
		AdInternalResponses: make(map[string]*AdInternalResponse),
	}

}
