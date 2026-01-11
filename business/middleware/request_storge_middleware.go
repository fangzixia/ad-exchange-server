package middleware

import (
	"ad-exchange-server/core/model"
	"ad-exchange-server/infra/redis"
	"time"
)

type Guid struct {
	Key string
}

func RequestStorgeAddMiddleware() MediaHandlerFunc {

	return func(mediaContent *model.AdMediaContent) bool {
		if mediaContent == nil || mediaContent.AdInternalResponse == nil || len(mediaContent.AdInternalResponse.AdInfos) == 0 {
			return true
		}
		guid := &Guid{
			Key: "",
		}
		_ = redis.SetValue(guid.Key, guid, 10*time.Minute)
		return true
	}

}
