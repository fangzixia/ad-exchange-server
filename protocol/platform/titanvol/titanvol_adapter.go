package titanvol

import (
	"ad-exchange-server/core/model"
	"encoding/json"
	"strconv"
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
		platformURL:  "http://127.0.0.1:8188/rtbhelper/receiver/adx/", // 模拟地址
	}
}

// MarshalRequest 内部平台方请求 -> PlatformY请求
func (b *Adapter) MarshalRequest(internalReq *model.AdPlatformContent) ([]byte, error) {
	bidReq := adaptRequest(internalReq)
	return json.Marshal(bidReq)
}

// UnmarshalResponse PlatformY响应 -> 内部统一响应
func (b *Adapter) UnmarshalResponse(c *model.AdPlatformContent, respBytes []byte) (*model.AdInternalResponse, error) {
	var bidResp AdResponse
	if err := json.Unmarshal(respBytes, &bidResp); err != nil {
		return nil, err
	}
	return adaptResponse(&bidResp), nil
}

// GetPlatformName 获取平台方名称
func (b *Adapter) GetPlatformName() string {
	return b.platformName
}

// GetPlatformURL 获取平台方接口地址
func (b *Adapter) GetPlatformURL() string {
	return b.platformURL
}

func adaptResponse(bidResp *AdResponse) *model.AdInternalResponse {
	resp := &model.AdInternalResponse{}
	if bidResp == nil || bidResp.SeatBid == nil || len(bidResp.SeatBid) == 0 {
		return nil
	}
	adInfos := make([]*model.AdInfo, len(bidResp.SeatBid))
	for i, bid := range bidResp.SeatBid {
		b := bid.Bid[0]
		adInfo := &model.AdInfo{}

		adInfo.Aid = 0
		adInfo.AdFormat = 0
		adInfo.Advertiser = b.AdverName
		adInfo.AdvertiserId = b.AdverId
		adInfo.Price = b.Price
		adInfo.WinNotice = b.Nurl
		adInfo.WinFailNotice = b.Lurl

		if (b.ClickAction == 1) || (b.ClickAction == 3) {
			adInfo.ClickAction = 0
		} else if b.ClickAction == 2 {
			adInfo.ClickAction = 2
		} else if b.ClickAction == 2 {
			adInfo.ClickAction = 1
		} else if b.ClickAction == 5 {
			adInfo.ClickAction = 3
		} else if b.ClickAction == 6 {
			adInfo.ClickAction = 5
		} else if b.ClickAction == 11 {
			adInfo.ClickAction = 4
		} else {
			adInfo.ClickAction = 0
		}

		adInfo.LandingPage = b.Landing
		adInfo.Source = b.Source
		adInfo.SourceLogoUrl = b.SourceLogo
		adInfo.Deeplink = b.DeepLink
		adInfo.DeepULink = b.DeepUlink
		if b.Isgdt == 1 {
			adInfo.IsGDT = true
		}

		if b.App != nil {
			app := model.AppInfo{}
			app.Name = b.App.Name
			app.Bundle = b.App.Bundle
			app.ItunesId = b.App.ItunesId
			app.MarketUrl = b.App.MarketUrl
			app.Relation = b.App.Relation
			app.IconUrl = b.App.IconUrl
			app.IconWidth = b.App.IconW
			app.IconHeight = b.App.IconH
			app.Version = b.App.Version
			app.VersionCode = b.App.VersionCode
			app.UpdateTime = b.App.UpdateTime
			app.QuickUrl = b.App.QuickUrl
			app.DownloadUrl = b.App.DownUrl
			app.ApkMD5 = b.App.Md5
			app.Developer = b.App.Developer
			app.Introduction = b.App.Intro
			app.IntroductionLink = b.App.IntroLink
			app.PrivacyPolicy = b.App.PrivatePolicy
			app.PrivacyPolicyLink = b.App.PrivatePolicyLink
			app.Permissions = b.App.Permissions
			app.PermissionsLink = b.App.PermissionsLink
			app.Size = b.App.Size
			app.Rating = b.App.Rating
			app.Downloads = b.App.Downloads
			app.Comments = b.App.Comments
			app.Snapshots = b.App.Snapshots
			app.Description = b.App.Desc
			app.Tags = b.App.Tag
			app.Icp = b.App.Icp
			adInfo.App = &app
		}
		if b.MiniProgram != nil {
			mp := model.MiniProgram{}
			mp.Id = b.MiniProgram.Id
			mp.AppId = b.MiniProgram.AppId
			mp.Path = b.MiniProgram.Path
			mp.Name = b.MiniProgram.Name
			adInfo.MiniProgram = &mp
		}
		if b.Trackings != nil && len(b.Trackings) > 0 {
			ts := make([]*model.EventTracking, len(b.Trackings))
			for i, t := range b.Trackings {
				ts[i] = &model.EventTracking{
					Type: t.EventType,
					Urls: t.Urls,
				}
			}
			adInfo.EventTracking = ts
		}

		creative := &model.Material{}
		creative.Title = b.Creative.Title
		creative.Description = b.Creative.Description
		creative.Cta = b.Creative.Cta
		creative.Rating = b.Creative.Rating
		if b.Creative.Icon == nil {
			icon := &model.Image{}
			icon.Url = b.Creative.Icon.Url
			icon.Width = b.Creative.Icon.Width
			icon.Height = b.Creative.Icon.Height
			icon.Size = b.Creative.Icon.Size
			icon.Ratio = b.Creative.Icon.Ratio
			icon.Mimes = b.Creative.Icon.Mimes
			creative.Icon = icon
		}
		if b.Creative.Images == nil && len(b.Creative.Images) > 0 {
			images := make([]*model.Image, len(b.Creative.Images))
			for mi, ci := range b.Creative.Images {
				image := &model.Image{}
				image.Url = ci.Url
				image.Width = ci.Width
				image.Height = ci.Height
				image.Size = ci.Size
				image.Ratio = ci.Ratio
				image.Mimes = ci.Mimes
				images[mi] = image
			}
			creative.Images = images
		}
		if b.Creative.Video != nil {
			creativeVideo := model.Video{}
			creativeVideo.Oriented = b.Creative.Video.Type
			creativeVideo.Title = b.Creative.Video.Title
			creativeVideo.Description = b.Creative.Video.Desc
			creativeVideo.Width = b.Creative.Video.Width
			creativeVideo.Height = b.Creative.Video.Height
			creativeVideo.Url = b.Creative.Video.Url
			creativeVideo.CoverUrl = b.Creative.Video.CoverUrl
			creativeVideo.CoverWidth = b.Creative.Video.CoverWidth
			creativeVideo.CoverHeight = b.Creative.Video.CoverHeight
			creativeVideo.Duration = b.Creative.Video.Duration
			creativeVideo.Size = b.Creative.Video.Size
			creativeVideo.Resolution = b.Creative.Video.Resolution
			creativeVideo.MinDuration = b.Creative.Video.MinDuration
			creativeVideo.AfterHtml = b.Creative.Video.Afterhtml
			creativeVideo.AfterButtonText = b.Creative.Video.AfterBtnText
			creativeVideo.AfterButtonUrl = b.Creative.Video.AfterBtnUrl
			creativeVideo.ValidTime = b.Creative.Video.ValidTime
			creativeVideo.PlayType = b.Creative.Video.PlayType
			if b.Creative.Video.Prefetch == 1 {
				creativeVideo.PrefetchEnable = true
			} else {
				creativeVideo.PrefetchEnable = false
			}
			creativeVideo.Mimes = b.Creative.Video.Mimes
			if b.Creative.Video.Trackings != nil && len(b.Creative.Video.Trackings) > 0 {
				pts := make([]*model.PointTrack, len(b.Creative.Video.Trackings))
				for pi, t := range b.Creative.Video.Trackings {
					pointTrack := model.PointTrack{
						Ts:   t.Ts,
						Urls: t.Urls,
					}
					pts[pi] = &pointTrack
				}
				creativeVideo.PointTrackers = pts
			}

			creative.CreativeVideo = &creativeVideo
		}

		creative.ImageMode = b.Creative.ImageMode
		adInfo.Creative = creative

		adInfos[i] = adInfo
	}
	resp.AdInfos = adInfos
	return resp
}

func adaptRequest(c *model.AdPlatformContent) *AdRequest {
	internalRequest := c.AdInternalRequest
	bidReq := AdRequest{}
	slots := internalRequest.AdSlots
	impressions := make([]*Impression, len(slots))
	for i, slot := range slots {
		impression := Impression{}
		impression.Aid = ""
		impression.W = 1024
		impression.H = 768
		impression.AdFormat = 0
		impression.BidFloor = slot.BidFloor
		impression.Bidcur = slot.BidCurrency
		if slot.BidType == "CPM" {
			impression.BidType = 1
		} else if slot.BidType == "CPM" {
			impression.BidType = 2
		}

		impression.Wmax = slot.Wmax
		impression.Hmax = slot.Hmax
		impression.Wmin = slot.Wmin
		impression.Hmin = slot.Hmin
		impression.Mimes = slot.Mimes
		if slot.SkipEnable {
			impression.Skip = 1
		} else {
			impression.Skip = 0
		}

		impression.SkipAfter = slot.SkipAfterTime
		impression.VideoType = slot.VideoType
		impression.Https = slot.SupportHttps
		impression.Deeplink = slot.SupportDeepLink
		impression.BlockCat = slot.BlockCat
		impression.Size = slot.MaterialSize
		impressions[i] = &impression
	}
	bidReq.Impression = impressions
	rd := internalRequest.Device
	device := Device{}
	device.Ua = rd.Ua
	if internalRequest.Geo != nil {
		device.Geo = &Geo{
			Lon:           internalRequest.Geo.Longitude,
			Lat:           internalRequest.Geo.Latitude,
			Type:          internalRequest.Geo.Type,
			Country:       internalRequest.Geo.Country,
			Province:      internalRequest.Geo.State,
			City:          internalRequest.Geo.City,
			District:      internalRequest.Geo.District,
			LocationAccur: internalRequest.Geo.LocationAccur,
		}
	}

	device.Ip = rd.Ip
	device.Ipv6 = rd.Ipv6
	device.DeviceType = rd.DeviceType
	device.Make = rd.Make
	device.Model = rd.Model
	device.Os = rd.Os
	device.Osv = rd.OsVersion
	device.Carrier = rd.Carrier
	device.SerialNo = rd.Sn
	device.Meid = rd.Meid
	device.ConnectionType = rd.Network
	device.Imsi = rd.Imsi
	device.Imei = rd.Imei
	device.ImeiMd5 = rd.ImeiMd5
	device.ImeiSha1 = rd.ImeiSha1
	device.Oaid = rd.Oaid
	device.OaidMd5 = rd.OaidMd5
	device.Adid = rd.Adid
	device.AdidMd5 = rd.AdidMd5
	device.Idfa = rd.Idfa
	device.IdfaMd5 = rd.IdfaMd5
	device.IdfaSha1 = rd.IdfaSha1
	device.Idfv = rd.Idfv
	device.IdfvMd5 = rd.IdfvMd5
	device.Mac = rd.Mac
	device.MacMd5 = rd.MacMd5
	device.MacSha1 = rd.MacSha1
	device.W = rd.ScreenWidth
	device.H = rd.ScreenHeight
	device.Dpi = rd.Dpi
	density, _ := strconv.ParseFloat(rd.Density, 10)
	device.Density = float32(density)
	if rd.Caids != nil && len(rd.Caids) > 0 {
		caids := make([]*Caid, len(rd.Caids))
		for i, caid := range rd.Caids {
			caids[i] = &Caid{
				Caid:    caid.Caid,
				Version: caid.Version,
			}
		}
		device.CaidList = caids
	}
	device.Aaid = rd.Aaid
	device.BootMark = rd.BootMark
	device.UpdateMark = rd.UpdateMark
	device.DeviceRomVer = rd.DeviceRomVer
	device.Hmscore = rd.Hmscore
	device.Paid = rd.PddAid
	device.Paid14 = rd.PddAid14
	device.ScreenOrientation = rd.ScreenOrientation
	device.Applist = rd.Applist
	device.HmsVer = rd.HmsVer
	device.HwagVer = rd.HwagVer
	device.DeviceName = rd.DeviceName
	device.DeviceNameMd5 = rd.DeviceNameMd5
	device.TimeZone = rd.TimeZone
	device.LocalTzName = rd.LocalTzName
	device.AppStoreVer = rd.AppStoreVer
	device.AppStoreUrl = rd.AppStoreUrl
	device.ApiLevel = rd.ApiLevel
	device.Ssid = rd.Ssid
	device.Bssid = rd.Bssid
	device.MiuiVer = rd.MiuiVer
	device.AuthStatus = rd.AuthStatus
	device.Openudid = rd.Udid
	device.BootTime = rd.BootTime
	device.BirthTime = rd.BirthTime
	device.UpdateTime = rd.UpdateTime
	device.ElapseTime = rd.ElapseTime
	device.SysCompileTime = rd.SysCompileTime
	device.MemorySize = strconv.Itoa(rd.MemorySize)
	device.DiskSize = strconv.Itoa(rd.DiskSize)
	device.DiskFreeSpace = strconv.Itoa(rd.DiskFreeSpace)
	device.BatteryStatus = rd.BatteryStatus
	device.BatteryPower = rd.BatteryPower
	device.CpuNum = rd.CpuNum
	device.CpuFreq = rd.CpuFreq
	device.HwModel = rd.HwModel
	device.HwName = rd.HwName
	device.HwMachine = rd.HwMachine
	device.HwVersion = rd.HwVersion
	device.AndroidSha1 = rd.AndroidSha1
	bidReq.Device = &device
	return &bidReq
}
