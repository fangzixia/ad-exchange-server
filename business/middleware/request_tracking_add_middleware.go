package middleware

import "ad-exchange-server/core/model"

func RequestTrackingAddMiddleware() MediaHandlerFunc {

	return func(mediaContent *model.AdMediaContent) bool {
		if mediaContent == nil || mediaContent.AdInternalResponse == nil || len(mediaContent.AdInternalResponse.AdInfos) == 0 {
			return true
		}
		for _, adInfo := range mediaContent.AdInternalResponse.AdInfos {
			tracking := adInfo.EventTracking
			adInfo.EventTracking = append(tracking, &model.EventTracking{
				Type: 1,
				Urls: []string{"localImp"},
			})
			adInfo.EventTracking = append(tracking, &model.EventTracking{
				Type: 2,
				Urls: []string{"localClick"},
			})
		}

		return true
	}

}
