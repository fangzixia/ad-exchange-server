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
func (b *Adapter) UnmarshalResponse(respBytes []byte) (*model.AdInternalResponse, error) {
	var titanvolResp AdResponse
	if err := json.Unmarshal(respBytes, &titanvolResp); err != nil {
		return nil, err
	}

	return &model.AdInternalResponse{
		AdInfos: make([]model.AdInfo, 1),
	}, nil
}

// GetPlatformName 获取平台方名称
func (b *Adapter) GetPlatformName() string {
	return b.platformName
}

// GetPlatformURL 获取平台方接口地址
func (b *Adapter) GetPlatformURL() string {
	return b.platformURL
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
