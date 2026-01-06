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
	ar.SeatBid = bids
	return ar

}
