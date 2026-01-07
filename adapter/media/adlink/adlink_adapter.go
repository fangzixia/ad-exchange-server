package adlink

import (
	"ad-exchange-server/core/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Adapter 媒体adLink适配器
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
func (m *Adapter) UnmarshalRequest(r *http.Request) *model.AdInternalRequest {
	var mediaAReq AdRequest
	if err := json.NewDecoder(r.Body).Decode(&mediaAReq); err != nil {
		return nil
	}
	return adaptRequest(&mediaAReq)
}

// MarshalResponse 内部统一响应 -> 媒体A响应
func (m *Adapter) MarshalResponse(internalResp *model.AdInternalResponse) ([]byte, error) {
	mediaAResp := adaptResponse(internalResp)
	return json.Marshal(mediaAResp)
}

// GetMediaType 获取媒体类型
func (m *Adapter) GetMediaType() string {
	return m.mediaType
}

func adaptRequest(r *AdRequest) *model.AdInternalRequest {
	ir := &model.AdInternalRequest{
		Id:        r.RequestID,
		Version:   model.AD_INTERNAL_VERSION,
		Timestamp: time.Now(),
	}
	ir.AdSlots = createAdSlot(r)
	ir.App = createApp(r)
	ir.Geo = createGeo(r)
	ir.User = createUser(r)
	ir.Device = createDevice(r)
	return ir

}

func createAdSlot(r *AdRequest) []*model.AdSlot {
	if r.AdSlot != nil {
		ia := &model.AdSlot{
			BidType:       "CPM",
			Wmax:          r.AdSlot.WMax,
			Wmin:          r.AdSlot.WMin,
			Hmax:          r.AdSlot.HMax,
			Hmin:          r.AdSlot.HMin,
			SkipEnable:    r.AdSlot.Skip == 1,
			SkipAfterTime: r.AdSlot.SkipAfter,
			VideoType:     r.AdSlot.VideoType,
			Mimes:         r.AdSlot.Mimes,
			Id:            r.RequestID,
			Aid:           r.AdSlot.SlotID,
			MediaAid:      r.AdSlot.SlotID,
			Width:         r.AdSlot.W,
			Height:        r.AdSlot.H,
			BidFloor:      r.AdSlot.BidFloor,
			SupportHttps:  r.Https,
		}
		return []*model.AdSlot{ia}
	}
	return nil
}

func createApp(r *AdRequest) *model.App {
	if r.App != nil {
		a := &model.App{
			Id:          r.App.AppID,
			Name:        r.App.Name,
			Bundle:      r.App.Bundle,
			Ver:         r.App.Ver,
			StoreUrl:    r.App.StoreUrl,
			AppCategory: r.App.Cat,

			IsPaid: r.App.Paid == 1,
		}
		if r.App.Keywords != "" {
			a.Keywords = strings.Split(r.App.Keywords, ",")
		}

		return a
	}
	return nil
}

func createGeo(r *AdRequest) *model.Geo {
	if r.Device != nil && r.Device.Geo != nil {
		g := &model.Geo{
			Latitude:  r.Device.Geo.Lat,
			Longitude: r.Device.Geo.Lon,
			District:  r.Device.Geo.District,
			City:      r.Device.Geo.City,
			State:     r.Device.Geo.Province,
		}
		return g
	}
	return nil
}

func createUser(r *AdRequest) *model.User {
	if r.User != nil {
		u := &model.User{
			Id:   r.User.UserID,
			Tags: r.User.Tags,
			Age:  r.User.Age,
		}
		if r.User.Gender == "male" {
			u.Gender = 1
		} else if r.User.Gender == "female" {
			u.Gender = 2
		} else {
			u.Gender = 0
		}
		if r.User.Keywords != "" {
			u.Keywords = strings.Split(r.User.Keywords, ",")
		}
		u.AppList = r.User.Applist
		return u
	}
	return nil
}

func createDevice(r *AdRequest) *model.Device {
	if r.Device == nil {
		return nil
	}
	d := &model.Device{}

	d.Ip = r.Device.Ip
	d.Ipv6 = r.Device.Ipv6
	d.Ua = r.Device.Ua
	d.Brand = r.Device.Brand
	d.Model = r.Device.Model
	d.Make = r.Device.Make
	d.OsVersion = r.Device.Osv
	d.ScreenWidth = r.Device.W
	d.ScreenHeight = r.Device.H
	d.ScreenSize = fmt.Sprintf("%d*%d", r.Device.W, r.Device.H)
	d.Oaid = r.Device.Oaid
	d.Imei = r.Device.Imei
	d.ImeiMd5 = r.Device.ImeiMd5
	d.ImeiSha1 = r.Device.ImeiSha1
	d.Adid = r.Device.AndroidID
	d.AdidMd5 = r.Device.AndroidIDMd5
	d.AndroidSha1 = r.Device.AndroidIDSha1
	d.Imsi = r.Device.Imsi
	d.Idfa = r.Device.Idfa
	d.IdfaMd5 = r.Device.IdfaMd5
	d.IdfaSha1 = r.Device.IdfaSha1

	d.Mac = r.Device.Mac
	d.MacMd5 = r.Device.MacMd5
	d.MacSha1 = r.Device.MacSha1
	d.Ssid = r.Device.Ssid
	d.Bssid = r.Device.Bssid
	d.Aaid = r.Device.Aaid
	d.Dpi = strconv.Itoa(r.Device.Dpi)
	d.Density = strconv.FormatFloat(float64(r.Device.Density), 'f', 2, 32)
	d.Language = r.Device.Language
	d.DeviceRomVer = r.Device.SysVer

	d.HwagVer = r.Device.HwagVer
	d.HmsVer = r.Device.HmsVer
	d.HwModel = r.Device.HwModel
	d.HwName = r.Device.HwName
	d.HwNameMd5 = r.Device.HwNameMd5
	d.HwMachine = r.Device.HwMachine
	sysMemory, _ := strconv.Atoi(r.Device.SysMemory)
	d.MemorySize = sysMemory
	sysDiskSize, _ := strconv.Atoi(r.Device.SysDisksize)
	d.DiskSize = sysDiskSize
	d.DeviceName = r.Device.DeviceName
	d.DeviceNameMd5 = r.Device.DeviceNameMd5
	d.BootMark = r.Device.BootMark
	d.UpdateMark = r.Device.UpdateMark
	d.BirthTime = r.Device.DeviceInitializeTime
	d.BootTime = r.Device.BootTimeSec
	d.UpdateTime = r.Device.OsUpdateTimeSec
	d.Idfv = r.Device.Idfv
	d.SysCompileTime = r.Device.SysCompilingTime
	d.Sn = r.Device.SerialNo
	d.CpuNum = r.Device.CpuNum
	//d.CpuFreq = r.Device.CP
	d.BatteryStatus = r.Device.BatteryStatus
	d.BatteryPower = r.Device.BatteryPower

	d.Udid = r.Device.OpenUdid
	d.Aaid = r.Device.Aaid
	d.Sn = r.Device.SerialNo
	if r.Device.Type == 1 {
		d.DeviceType = 1
	} else if r.Device.Type == 2 {
		d.DeviceType = 2
	} else {
		d.DeviceType = 0
	}

	if r.Device.Os == 1 {
		d.Os = 2
	} else if r.Device.Os == 2 {
		d.Os = 1
	} else {
		d.Os = 0
	}

	if r.Device.Connection == 6 {
		d.Network = 1
	} else if r.Device.Connection == 2 {
		d.Network = 2
	} else if r.Device.Connection == 3 {
		d.Network = 3
	} else if r.Device.Connection == 4 {
		d.Network = 4
	} else if r.Device.Connection == 5 {
		d.Network = 5
	} else {
		d.Network = 0
	}

	if "46000" == r.Device.Carrier {
		d.Carrier = 2
	} else if "46001" == r.Device.Carrier {
		d.Carrier = 3
	} else if "46002" == r.Device.Carrier {
		d.Carrier = 1
	} else if "46003" == r.Device.Carrier {
		d.Carrier = 4
	} else if "46020" == r.Device.Carrier {
		d.Carrier = 1
	} else {
		d.Carrier = 0
	}

	if r.Device.Caids != nil && len(r.Device.Caids) > 0 {
		caid := make([]*model.Caid, len(r.Device.Caids))
		for i, c := range r.Device.Caids {
			split := strings.Split(c, ",")
			mc := &model.Caid{
				Caid:    split[0],
				Version: split[1],
			}
			caid[i] = mc
		}
		d.Caids = caid
	}
	d.Ppi = strconv.Itoa(r.Device.Ppi)
	return d
}

func adaptResponse(ir *model.AdInternalResponse) *AdResponse {
	ar := &AdResponse{}
	if ir.AdInfos == nil || len(ir.AdInfos) == 0 {
		ar.Code = 0
		ar.Message = "广告内容为空"
		return ar
	}
	bids := make([]*SeatBid, len(ir.AdInfos))
	for i, adInfo := range ir.AdInfos {
		bid := &SeatBid{}
		b := &Bid{}

		b.BidType = 0
		b.Price = float64(adInfo.Price)
		b.Crid = adInfo.Creative.CreativeId
		b.Nurl = adInfo.WinNotice
		b.Lurl = adInfo.WinFailNotice

		mm := &MaterialMeta{}

		mm.Title = adInfo.Creative.Title
		mm.Description = adInfo.Creative.Description
		mm.Cta = adInfo.Creative.Cta
		mm.Rating = strconv.FormatFloat(float64(adInfo.Creative.Rating), 'f', 2, 64)
		if adInfo.Creative.Icon != nil {
			icon := Image{}
			icon.Url = adInfo.Creative.Icon.Url
			icon.W = adInfo.Creative.Icon.Width
			icon.H = adInfo.Creative.Icon.Height
			icon.AspectRatio = float64(adInfo.Creative.Icon.Ratio)
			mm.Icon = &icon
		}

		if adInfo.Creative.Images != nil && len(adInfo.Creative.Images) > 0 {
			images := make([]*Image, len(adInfo.Creative.Images))
			for ii, img := range adInfo.Creative.Images {
				im := Image{}
				im.Url = img.Url
				im.W = img.Width
				im.H = img.Height
				im.AspectRatio = float64(img.Ratio)
				if ii == 0 {
					mm.Image = &im
				}
				images[ii] = &im
			}
			mm.Images = images
		}
		mm.ImageMode = adInfo.Creative.ImageMode
		if adInfo.Creative.CreativeVideo != nil {

			v := Video{}

			v.Url = adInfo.Creative.CreativeVideo.Url
			v.Duration = float32(adInfo.Creative.CreativeVideo.Duration)
			v.Size = float32(adInfo.Creative.CreativeVideo.Size)
			v.Resolution = adInfo.Creative.CreativeVideo.Resolution
			v.Cover = adInfo.Creative.CreativeVideo.CoverUrl
			v.Coverh = adInfo.Creative.CreativeVideo.CoverHeight
			v.Coverw = adInfo.Creative.CreativeVideo.CoverWidth
			v.Title = adInfo.Creative.CreativeVideo.Title
			v.Desc = adInfo.Creative.CreativeVideo.Description
			v.AfterHtml = adInfo.Creative.CreativeVideo.AfterHtml
			v.ForceDuration = adInfo.Creative.CreativeVideo.MinDuration
			v.CompleteCoverUrl = adInfo.Creative.CreativeVideo.EndCoverUrl
			v.Skip = adInfo.Creative.CreativeVideo.Oriented
			if adInfo.Creative.CreativeVideo.PrefetchEnable {
				v.Preload = 1
			} else {
				v.Preload = 0
			}
			mm.Video = &v
		}
		mm.HtmlSnippet = adInfo.Creative.HtmlSnippet
		if adInfo.MiniProgram != nil {
			mp := MiniProgram{}
			mp.ID = adInfo.MiniProgram.Id
			mp.Path = adInfo.MiniProgram.Path
			mp.Name = adInfo.MiniProgram.Name
			mp.AppId = adInfo.MiniProgram.AppId
			mm.MiniProgram = &mp
		}

		if adInfo.App != nil {
			appData := AppData{}
			appData.Name = adInfo.App.Name
			appData.Bundle = adInfo.App.Bundle
			appData.Ver = adInfo.App.Version
			appData.Icon = adInfo.App.IconUrl
			appData.Intro = adInfo.App.Introduction
			appData.Size = float64(adInfo.App.Size)
			appData.Privacy = adInfo.App.PrivacyPolicy
			appData.PrivacyAgreement = adInfo.App.PrivacyPolicyLink
			appData.Developer = adInfo.App.Developer
			appData.PaymentType = adInfo.App.IconUrl
			b.App = &appData
		}
		b.ClickAction = adInfo.ClickAction
		b.Landing = adInfo.LandingPage
		b.Source = adInfo.Source
		b.SourceLogo = adInfo.SourceLogoUrl
		b.Deeplink = adInfo.Deeplink

		if adInfo.Creative.ImpUrls != nil && len(adInfo.Creative.ImpUrls) > 0 {
			track := b.EventTrack
			if track == nil {
				track = &EventTrack{}
			}
			track.ImpTracks = adInfo.Creative.ImpUrls

		}

		if adInfo.Creative.ClickUrls != nil && len(adInfo.Creative.ClickUrls) > 0 {
			track := b.EventTrack
			if track == nil {
				track = &EventTrack{}
			}
			track.ClkTracks = adInfo.Creative.ClickUrls

		}

		if adInfo.Creative.LocalTrackingEvent != nil {
			track := b.EventTrack
			track.ImpTracks = adInfo.Creative.ImpUrls
			track.ImpTracks = append(track.ImpTracks)
		}

		b.Creative = mm
		bid.Bid = b

		bids[i] = bid
	}
	ar.SeatBid = bids
	return ar

}
