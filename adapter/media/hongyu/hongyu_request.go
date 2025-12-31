package hongyu

// DEFAULT_VERSION 默认协议版本号常量
const DEFAULT_VERSION = "1.6.0"

// AdRequest AD-Link星联移动ADX广告平台请求DTO
type AdRequest struct {
	RequestID  string      `json:"requestID"`  // 标识一次请求的唯一ID
	AdSlot     *AdSlot     `json:"adslot"`     // 广告位信息
	Site       *Site       `json:"site"`       // 网站信息
	App        *App        `json:"app"`        // APP信息
	Device     *Device     `json:"device"`     // 设备信息
	User       *User       `json:"user"`       // 用户信息
	At         int         `json:"at"`         // Auction类型，1:一价结算；2:二价结算；默认为2
	TMax       int         `json:"tmax"`       // 最大返回响应时长，单位毫秒，包括网络传输时间
	BCat       string      `json:"bcat"`       // 屏蔽行业，以逗号分隔，见IAB标准分类定义
	BAdv       string      `json:"badv"`       // 屏蔽域名，以逗号分隔，见IAB标准分类定义
	Cur        []string    `json:"cur"`        // 货币类型：USD–美元，CNY–元，默认CNY
	Https      int         `json:"https"`      // 是否支持https，0:不限，1:只支持https
	Support302 int         `json:"support302"` // 是否支持302跳转类型链接，1:支持，0:不支持
	Deeplink   int         `json:"deeplink"`   // 是否支持deeplink，0:未知，1:支持，2:不支持
	MediaTime  int64       `json:"mediaTime"`  // 媒体端发出请求的时间，单位：ms，建议填写
	SspTime    int64       `json:"sspTime"`    // 渠道端发出请求的时间，单位：ms，建议填写
	Ext        interface{} `json:"ext"`        // 扩展信息
}

// AdSlot 广告位信息
type AdSlot struct {
	TemplateID  string   `json:"templateID"`  // 媒体请求素材模板标识
	SlotID      string   `json:"slotID"`      // 媒体在DSP平台注册的广告位唯一标识
	AdType      int      `json:"adType"`      // 广告类型：banner=1，插屏=2，开屏=3，信息流=4，贴片=5，激励视频=6；其它原生广告=7
	Pos         int      `json:"pos"`         // 广告展现位置：顶部=1；底部=2；信息流内=3；中部=4；全屏=5
	BidType     int      `json:"bidType"`     // 允许的竞价类型，目前仅支持CPM(0-CPM,1-CPC,2-CPA)
	BidFloor    int64    `json:"bidFloor"`    // 最低千次展现单价，单位人民币分
	W           int      `json:"w"`           // 广告位实际尺寸宽，单位像素
	H           int      `json:"h"`           // 广告位实际尺寸高，单位像素
	WMax        int      `json:"wmax"`        // 允许的最大宽度值
	HMax        int      `json:"hmax"`        // 允许的最大高度值
	WMin        int      `json:"wmin"`        // 允许的最小宽度值
	HMin        int      `json:"hmin"`        // 允许的最小高度值
	MinDuration int      `json:"minDuration"` // 允许的视频最短时长，单位：秒
	MaxDuration int      `json:"maxDuration"` // 允许的视频最大时长，单位：秒
	MinBitRate  int      `json:"minBitRate"`  // 允许的视频最小比特率，单位:Kbps
	MaxBitRate  int      `json:"maxBitRate"`  // 允许的视频最大比特率，单位:Kbps
	Skip        int      `json:"skip"`        // 是否允许跳过
	SkipAfter   int      `json:"skipAfter"`   // 几秒后允许跳过
	VideoType   int      `json:"videoType"`   // 视频类型，0-普通视频，1-激励视频
	BannerType  int      `json:"bannerType"`  // Banner广告类型，0-JSON，1-HTML。Banner必传
	Mimes       []string `json:"mimes"`       // 支持的素材内容类型
	Size        int      `json:"size"`        // 允许的最大文件大小，单位:KB
	Ext         *Ext     `json:"ext"`         // 扩展信息
}

// Site 网站信息
type Site struct {
	SiteID string `json:"siteID"` // 网站在DSP平台唯一标识
	Name   string `json:"name"`   // 网站名称
	Domain string `json:"domain"` // 网站域名
	Page   string `json:"page"`   // 当前页面
	Ref    string `json:"ref"`    // referer
	Cat    string `json:"cat"`    // 行业分类代码，见附录
}

// App APP信息
type App struct {
	AppID    string `json:"appID"`    // 应用在DSP平台唯一标识
	Name     string `json:"name"`     // 应用名称
	Bundle   string `json:"bundle"`   // 应用包名
	Ver      string `json:"ver"`      // 应用版本号
	StoreUrl string `json:"storeUrl"` // App应用商店地址
	Cat      string `json:"cat"`      // 行业分类代码，见附录
	Keywords string `json:"keywords"` // App关键字描述
	Paid     int    `json:"paid"`     // 0-免费，1-付费
}

// Device 设备信息
type Device struct {
	Type                 int      `json:"type"`                 // 设备类型，0-未知，1-phone，2-tablet，3-其他
	Dnt                  int      `json:"dnt"`                  // DoNotTrack标记 0-cantrack 1-don'ttrack
	Ua                   string   `json:"ua"`                   // User agent
	Ip                   string   `json:"ip"`                   // IP地址
	Ipv6                 string   `json:"ipv6"`                 // IPv6地址
	Make                 string   `json:"make"`                 // 制造商
	Model                string   `json:"model"`                // 设备型号
	Os                   int      `json:"os"`                   // 操作系统，iOS=1,Android=2
	Osv                  string   `json:"osv"`                  // 操作系统版本号
	W                    int      `json:"w"`                    // 屏幕宽，单位像素
	H                    int      `json:"h"`                    // 屏幕高，单位像素
	Carrier              string   `json:"carrier"`              // 运营商
	Connection           int      `json:"connection"`           // 设备网络连接类型，0-未知，1-ethernet，2-2G，3-3G，4-4G，5-5G，6-wifi
	Oaid                 string   `json:"oaid"`                 // OAID，Android10以上必传
	Imei                 string   `json:"imei"`                 // IMEI原文
	ImeiMd5              string   `json:"imeiMd5"`              // IMEI的MD5值
	ImeiSha1             string   `json:"imeiSha1"`             // IMEI的SHA1值
	AndroidID            string   `json:"androidID"`            // AndroidID原文
	AndroidIDMd5         string   `json:"androidIDMd5"`         // AndroidID的MD5值
	AndroidIDSha1        string   `json:"androidIDSha1"`        // AndroidID的SHA1值
	Imsi                 string   `json:"imsi"`                 // IMSI
	Idfa                 string   `json:"idfa"`                 // IDFA原文
	IdfaMd5              string   `json:"idfaMd5"`              // IDFA的MD5值
	IdfaSha1             string   `json:"idfaSha1"`             // IDFA的SHA1值
	Caids                []string `json:"caids"`                // caid版本号和实际值
	Mac                  string   `json:"mac"`                  // MAC地址
	MacSha1              string   `json:"macSha1"`              // MAC的SHA1值
	MacMd5               string   `json:"macMd5"`               // MAC的MD5值
	Ssid                 string   `json:"ssid"`                 // 无线网SSID名称
	Bssid                string   `json:"bssid"`                // Wifi MAC地址
	Paid                 string   `json:"paid"`                 // Pinduoduo Advertising Identifier
	Aaid                 string   `json:"aaid"`                 // 匿名广告标识符
	Geo                  *Geo     `json:"geo"`                  // 地理位置信息
	Dpi                  int      `json:"dpi"`                  // 每英寸像素个数
	Density              float32  `json:"density"`              // dpi/160
	Language             string   `json:"language"`             // 语言，如zh-CN
	SysVer               string   `json:"sysVer"`               // 手机ROM版本号
	AppStorePackage      string   `json:"appStorePackage"`      // 安卓设备应用商店包名
	HwagVer              string   `json:"hwagVer"`              // 安卓设备应用商店版本号
	HmsVer               string   `json:"hmsVer"`               // 华为机型HMSCore版本号
	HwModel              string   `json:"hwModel"`              // 硬件设备型号
	HwName               string   `json:"hwName"`               // 硬件设备名称
	HwNameMd5            string   `json:"hwNameMd5"`            // 硬件设备名称MD5
	HwMachine            string   `json:"hwMachine"`            // 硬件系统型号
	SysMemory            string   `json:"sysMemory"`            // 系统内存
	SysDisksize          string   `json:"sysDisksize"`          // 用户设置的手机设备名称
	DeviceName           string   `json:"deviceName"`           // 用户设置的手机设备名称
	DeviceNameMd5        string   `json:"deviceNameMd5"`        // 用户设置的手机设备名称MD5
	BootMark             string   `json:"bootMark"`             // 系统启动标识
	UpdateMark           string   `json:"updateMark"`           // 系统更新标识
	DeviceInitializeTime string   `json:"deviceInitializeTime"` // 设备初始化时间(秒)
	BootTimeSec          string   `json:"bootTimeSec"`          // 系统启动时间标识(秒)
	OsUpdateTimeSec      string   `json:"osUpdateTimeSec"`      // 系统更新时间标识(秒)
	Brand                string   `json:"brand"`                // 设备品牌,iOS系统设备请统一填写为Apple
	Idfv                 string   `json:"idfv"`                 // IOS移动设备idfv值
	OpenUdid             string   `json:"openUdid"`             // IOS移动设备OpenUDID
	SysCompilingTime     string   `json:"sysCompilingTime"`     // 系统编译时间戳，单位：毫秒
	SerialNo             string   `json:"serialNo"`             // 设备序列号
	CpuNum               int      `json:"cpuNum"`               // 设备CPU个数
	BatteryStatus        int      `json:"batteryStatus"`        // 电池充电状态
	BatteryPower         int      `json:"batteryPower"`         // 电池电量百分比
	Ppi                  int      `json:"ppi"`                  // 屏幕每英寸像素数
}

// User 用户信息
type User struct {
	UserID   string   `json:"userID"`   // 用户ID
	Tags     string   `json:"tags"`     // 用户标签
	Gender   string   `json:"gender"`   // 性别
	Age      int      `json:"age"`      // 年龄
	Keywords string   `json:"keywords"` // 关键字
	Applist  []string `json:"applist"`  // 应用列表
}

// Geo 地理位置信息
type Geo struct {
	Lat      float32 `json:"lat"`      // 纬度，浮点数
	Lon      float32 `json:"lon"`      // 经度，浮点数
	City     string  `json:"city"`     // 城市
	Province string  `json:"province"` // 省份
	District string  `json:"district"` // 区县
}

// Ext 扩展信息基类
type Ext struct{}
