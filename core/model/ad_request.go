package model

import "time"

// AdInternalRequest AdInternalRequest 内部统一广告请求模型（所有媒体协议最终转换为此格式）
type AdInternalRequest struct {
	// 请求ID
	Id string
	// 协议版本号
	Version string
	// 请求时间戳（毫秒）
	Timestamp time.Time
	// 广告(曝光)信息，同一请求可包含多个曝光请求
	AdSlots []*AdSlot
	// 设备信息
	Device *Device
	// 媒体应用信息
	App *App
	// 地域信息
	Geo *Geo
	// 用户信息
	User *User
}

// Geo 广告请求地理信息类，处理位置相关业务逻辑
type Geo struct {
	// 国家ID
	CountryId int
	// 国家名称
	Country string
	// 州/省 ID
	StateId int
	// 州/省名称
	State string
	// 城市 ID
	CityId int
	// 城市名称
	City string
	// 区 ID
	DistrictId int
	// 区名称
	District string
	// 设备所在地理位置的经度
	Longitude float32
	// 设备所在地理位置的纬度
	Latitude float32
	// 定位类型：
	// 1-gps 全球定位系统
	// 2-wgs-84 世界大地测量系统1984
	// 3-GCJ-02 火星坐标系/国家测绘局02坐标系
	// 4-BD-09 百度坐标系
	// 5-CGCS2000 国家大地坐标系
	Type int

	// 经纬度精度半径，单位为米
	LocationAccur float32
}

// AppInfo 应用对象
type App struct {
	// 应用ID
	Id string
	// 应用名称
	Name string
	// 应用包名
	Bundle string
	// APP版本号
	Ver string
	// 应用商店地址
	StoreUrl string
	// app分类
	AppCategory string
	// app描述
	Desc string
	// 应用关键字
	Keywords []string
	// 是否付费应用
	IsPaid bool
}

// AdSlot 曝光信息抽象类
type AdSlot struct {
	// 曝光ID
	Id string
	// 广告位ID
	Aid string
	// 媒体广告位ID
	MediaAid string
	// 底价，单位是分
	BidFloor int
	// 竞价币种：USD-美元、CNY-人民币
	BidCurrency string
	// 竞价类型（原Java中为枚举，暂定义为字符串，可根据实际枚举值补充）
	BidType string
	// 当前广告位所在媒体环境的关键字（多个关键字用半角/全角逗号分割）
	Keywords []string
	// 广告位的当前宽度
	Width int
	// 广告位的当前高度
	Height int
	// 广告位允许的最大宽度
	Wmax int
	// 广告位允许的最大高度
	Hmax int
	// 广告位允许的最小宽度
	Wmin int
	// 广告位允许的最小高度
	Hmin int
	// 广告位支持的 MIME 类型列表，如["image/jpg","image/gif"]
	Mimes []string
	// 广告位是否支持视频：true-支持、false-不支持
	SupportVideo bool
	// 广告位视频是否允许跳过：1-可跳过，0-不可跳过
	SkipEnable bool
	// 广告位视频允许跳过的时间，单位：秒
	SkipAfterTime int
	// 视频广告类型（原Java中为枚举，暂定义为字符串）
	VideoType string
	// 是否支持微信小程序：true-支持、false-不支持
	SupportWx bool
	// 是否支持302跳转：true-支持、false-不支持
	Support302 bool
	// 是否支持深度链接：0-未知，1-支持，2-不支持
	SupportDeepLink int
	// 是否支持https：0-不限、1-支持、2-不支持
	SupportHttps int
	// 屏蔽行业
	BlockCat []string
	// 素材大小，单位：KB
	MaterialSize int
}

// Device 设备对象
type Device struct {
	// 设备客户端User Agent信息
	Ua string
	// 设备所在ip地址（公网IP）
	Ip string
	// 设备所在ipv6地址（公网IPv6）
	Ipv6 string
	// 设备类型（枚举：手机/平板/PC等，暂定义为字符串）
	DeviceType string
	// 制造商
	Make string
	// 设备型号
	Model string
	// 设备品牌
	Brand string
	// 用户终端操作系统（枚举：Android/iOS等，暂定义为字符串）
	Os string
	// 操作系统版本号
	OsVersion string
	// 运营商（枚举：移动/联通/电信等，暂定义为字符串）
	Carrier string
	// 设备序列号
	Sn string
	// 手机设备的meid号（电信必备）
	Meid string
	// 网络类型（枚举：4G/WiFi等，暂定义为字符串）
	Network string
	// 国际移动客户识别码
	Imsi string
	// 用户终端的IMEI（md5加密）
	ImeiMd5 string
	// 原始imei
	Imei string
	// imei sha1加密
	ImeiSha1 string
	// 匿名设备标识符
	Oaid string
	// oaid原始值md5加密后的值
	OaidMd5 string
	// 安卓id
	Adid string
	// 安卓id md5加密后的值
	AdidMd5 string
	// IOS IDFA（适用于IOS6及以上，md5加密）
	IdfaMd5 string
	// 原始idfa
	Idfa string
	// idfa sha1加密
	IdfaSha1 string
	// ios idfv（>=iOS6，安卓无需填写）
	Idfv string
	// idfv md5加密
	IdfvMd5 string
	// 用户终端eth0接口的MAC地址（大写去除冒号分隔符，md5加密）
	MacMd5 string
	// 原始mac
	Mac string
	// 用户终端eth0接口的MAC地址（大写保留冒号分隔符，md5加密）
	MacSha1 string
	// 屏幕宽
	ScreenWidth int
	// 屏幕高
	ScreenHeight int
	// 设备屏幕图像密度（物理像素和设备独立像素的比例）
	Dpi string
	// 设备屏幕像素密度（每英寸屏幕像素数，如286ppi）
	Ppi string
	// 屏幕密度
	Density string
	// 屏幕尺寸
	ScreenSize string
	// caid列表
	Caids []Caid
	// 匿名广告设备标识符（阿里AAID）
	Aaid string
	// 系统启动标识（原始传输）
	BootMark string
	// iOS：设备开机/启动时间（保留9位小数，纳秒），示例：1713834795.235361342
	BootTime string
	// iOS：系统初始化时间（保留9位小数，纳秒），示例：1595214620.383940897
	BirthTime string
	// OS：开机时长（单位毫秒）
	ElapseTime int
	// iOS/Android：系统编译时间（保留9位小数，纳秒），示例:1722989578.886315266
	SysCompileTime string
	// iOS：系统更新时间（保留9位小数，纳秒），示例：1595214620.383940897
	UpdateTime string
	// 系统更新标识（原始传输）
	UpdateMark string
	// 手机rom的版本号
	DeviceRomVer string
	// 拼多多广告ID
	PddAid string
	// 拼多多广告ID 1.4版本
	PddAid14 string
	// 屏幕朝向：0-未知、1-竖屏、2-横屏
	ScreenOrientation int
	// 当前手机上的安装app列表 [应用名：包名，应用名：包名]
	Applist []string
	// 华为HMS Core版本号
	Hmscore string
	// 华为应用市场的版本号（华为必填）
	HmsVer string
	// 华为AG的版本（华为必填，即应用市场版本号）
	HwagVer string
	// IOS：用户设置的手机设备名称
	DeviceName string
	// IOS：用户设置的手机设备名称MD5
	DeviceNameMd5 string
	// iOS：当前系统所在时区，示例：GMT+8/GMT+08:00
	TimeZone string
	// iOS：local时区与格林威治的时间差（单位s），如"Asia/Shanghai"=>28800（iOS必填）
	LocalTzTime int
	// iOS：local时区，如"Asia/Shanghai"
	LocalTzName string
	// iOS：设备语言（ISO-639-1/alpha-2），示例：zh-Hans-CN（iOS）,zh（安卓）
	Language string
	// 厂商应用商店版本号（vivo、小米、oppo等）
	AppStoreVer string
	// 应用商店地址
	AppStoreUrl string
	// Android API level（iOS不填）
	ApiLevel string
	// 当前连接WiFi的SSID名称
	Ssid string
	// WIFI路由器MAC地址
	Bssid string
	// MIUI版本号（小米设备必填）
	MiuiVer string
	// IOS：广告标识授权情况：0-未允许、1-受限制、2-拒绝、3-授权通过
	AuthStatus int
	// IOS UDID
	Udid string
	// iOS：系统内存（单位kb），示例：524288
	MemorySize int
	// iOS：系统硬盘大小（单位kb），示例：262144
	DiskSize int
	// 系统硬盘剩余空间（单位kb）
	DiskFreeSpace int
	// iOS：电池充电状态：1-未知、2-不在充电、3-正在充电
	BatteryStatus int
	// iOS：电量百分比（1-100）
	BatteryPower int
	// iOS：cpu数量
	CpuNum int
	// iOS：手机cpu频率（单位GHz）
	CpuFreq int
	// 硬件设备型号：安卓："Redmi K30Pro"，iOS："iPhone"
	HwModel string
	// 硬件设备名称（原值与md5填其一，iOS获取不到IDFA时必填）
	HwName string
	// 硬件设备名称MD5（原值与md5填其一，iOS获取不到IDFA时必填）
	HwNameMd5 string
	// 硬件系统型号：安卓："QKQ1.191117.002"，iOS："iPhone8,1"
	HwMachine string
	// 硬件型号版本（iOS必填），例如iPhone5S中的5S
	HwVersion string
	// 安卓SHA1值（签名证书指纹识别），示例：9A:B4:44:EF:43:D1:G5:85:44:65:16:3F:33:BA:DC:1F:28:6F:44:BC
	AndroidSha1 string
	// cookie
	Cookie string
	// 来源页
	Referer string
	// 电话号码
	PhoneNumber string
}

// Caid 热云数据设备标识类
// iOS中idfa/idfv/caid必须传一项（iOS14+可能获取不到idfa）
type Caid struct {
	// CAID设备标识符
	Caid string
	// CAID版本号
	Version string
}

// User 用户对象
type User struct {
	// 媒体侧用户id
	Id string
	// 用户标签（多标签逗号隔开）
	Tags string
	// 性别：0-未知、1-男、2-女
	Gender int
	// 年龄
	Age int
	// 用户画像关键字列表（多关键字逗号隔开）
	Keywords []string
	// 出生日期，格式为YYYYMMDD
	DateOfBirth int
}
