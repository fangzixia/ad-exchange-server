package model

// AdInternalResponse 内部统一广告响应模型（所有媒体协议最终转换为此格式）
type AdInternalResponse struct {
	RequestId string   // 对应的广告请求 id
	AdInfos   []AdInfo // 广告填充实例
}

// AdInfo 广告请求填充物
type AdInfo struct {
	Aid           int64           // 广告位 id
	AdFormat      int             // 广告展现形式
	Price         int             // 上游平台竞标出价
	WinNotice     string          // 竞价成功通知地址
	WinFailNotice string          // 竞价失败通知地址
	ClickAction   int             // 点击动作类型：0-打开网页 1-下载 2-appstore 3-deeplink 4-微信小程序 5-广点通二次下载
	LandingPage   string          // 落地页地址
	Source        string          // 广告来源标识文字
	SourceLogoUrl string          // 广告来源标识的logo图片url
	Deeplink      string          // deeplink地址
	DeepULink     string          // IOS通用唤起链接(UniversalLink)
	IsGDT         bool            // 是否广点通广告
	Advertiser    string          // 广告主名称
	AdvertiserId  string          // 广告主ID
	App           AppInfo         // APP信息（若广告是APP下载类型）
	MiniProgram   MiniProgram     // 小程序信息（若广告是小程序类型）
	ThirdTracings []ThirdTracking // 监测地址
	Creative      Creative        // 创意对象列表
}

// AppInfo 下载的应用信息
type AppInfo struct {
	Name              string   // 应用名称
	Bundle            string   // 应用包名
	ItunesId          string   // iOS应用商店Id
	MarketUrl         string   // 应用市场地址（仅安卓端存在）
	Relation          int      // 广告的应用安装定向类型：0-不限 1-定向未安装 2-定向已安装
	IconUrl           string   // 应用图标
	IconWidth         int      // 图标宽度
	IconHeight        int      // 图标高度
	Version           string   // 应用版本号
	VersionCode       int      // 应用版本代码（数字）
	UpdateTime        string   // 应用更新时间
	QuickUrl          string   // 快应用地址
	DownloadUrl       string   // 应用下载地址
	ApkMD5            string   // 应用APK文件的MD5值（仅安卓）
	Developer         string   // 应用开发者
	Introduction      string   // 应用介绍
	IntroductionLink  string   // 应用介绍链接
	PrivacyPolicy     string   // 应用隐私政策
	PrivacyPolicyLink string   // 应用隐私政策链接
	PermissionsLink   string   // 应用权限链接
	Permissions       string   // 应用权限描述
	Size              int      // 应用大小（单位KB）
	Rating            float32  // 应用评分
	Downloads         int      // 应用下载次数
	Comments          int      // 应用评论数
	Snapshots         []string // 应用截图列表
	Description       string   // 应用描述
	Tags              []string // 应用标签
	Icp               string   // ICP备案信息
}

// MiniProgram 小程序信息（对应 MiniProgram.java）
type MiniProgram struct {
	Id    string // 小程序原始id
	Path  string // 小程序跳转路径
	Name  string // 小程序名称
	AppId string // 开放平台id
}

// ThirdTracking 第三方广告监测实体（对应 ThirdTracking.java）
type ThirdTracking struct {
	Type int      // 监测类型（映射 TrackingType 枚举，Go 中用 int 替代）
	Urls []string // 监测URL列表
}

// Creative 广告创意类（对应 Creative.java）
type Creative struct {
	Title              string          // 创意标题
	Description        string          // 创意描述
	Cta                string          // 广告按钮文字
	Rating             float32         // 评分等级
	Icon               Image           // 广告创意Icon
	Images             []Image         // 广告创意图片
	ImageMode          int             // 素材模式：小图=2 大图=3 组图=4 横屏视频=5 竖屏视频=6
	CreativeVideo      CreativeVideo   // 视频类型创意
	CreativeId         string          // 创意ID
	HtmlSnippet        string          // Html类型的创意代码
	Index              int             // 创意在广告请求中的唯一序号
	LocalTrackingEvent []TrackingEvent // 自身平台的监测（已废弃）
	ImpUrls            []string        // 上游平台返回的曝光监测地址列表
	ClickUrls          []string        // 上游平台返回的点击监测地址列表
	ThirdTrackingUrls  []ThirdTracking // 其他事件监测
}

// Image 图片信息（对应 Image.java）
type Image struct {
	Url    string   // 图片地址
	Width  int      // 图片宽度
	Height int      // 图片高度
	Size   int      // 图片大小（单位字节）
	Ratio  float32  // 图片宽高比
	Mimes  []string // 支持的图片内容类型：image/jpg、image/jpeg、image/gif、image/png
}

// CreativeVideo 视频类型创意（对应 CreativeVideo.java）
type CreativeVideo struct {
	Oriented           int              // 视频方向：1-竖屏 2-横屏
	Title              string           // 视频播放完成后显示的广告标题
	Description        string           // 视频播放完成后显示的广告描述
	Width              int              // 视频宽度
	Height             int              // 视频高度
	Url                string           // 素材地址
	CoverUrl           string           // 视频封面图片地址
	CoverWidth         int              // 视频封面图片宽度
	CoverHeight        int              // 视频封面图片高度
	EndCoverUrl        string           // 视频结束封面图片地址
	Duration           int              // 播放时长（单位秒）
	MinDuration        int              // 最短播放时长（单位秒）
	Size               int              // 视频大小（单位KB）
	Resolution         string           // 视频分辨率（如 1920*1080）
	AfterHtml          string           // 视频播放完成后显示的HTML代码
	AfterButtonText    string           // 视频播放完成后显示的按钮文字
	AfterButtonUrl     string           // 视频播放完成后按钮点击跳转链接
	ValidTime          int              // 激励视频预加载有效时间（单位秒）
	PlayType           int              // 播放类型：1-自动播放 2-点击播放
	PrefetchEnable     bool             // 是否开启预加载
	Mimes              []string         // 支持的视频内容类型
	PointTrackers      []PointTrack     // 视频按秒上报事件
	BackgroundImage    string           // 视频背景图片（不播放视频时显示）
	VideoTracking      VideoTracking    // 视频监测信息
	ImpressionTracking map[int][]string // 播放监测URL
	KeepSecond         int              // 强制观看秒数
	Icon               string           // 激励视频特有字段
}

// PointTrack 视频按秒上报事件（CreativeVideo 内部类）
type PointTrack struct {
	Ts   int      // 视频播放到ts秒时上报
	Urls []string // 上报地址列表
}

type TrackingEvent struct{} // 自身平台监测事件（对应 TrackingEvent.java）
type VideoTracking struct{} // 视频监测信息（对应 VideoTracking.java）
type Event struct {         // 创意交互事件（对应 Event.java）
	Type  int    // 交互类型（映射 Interaction 枚举，Go 中用 int 替代）
	Value string // 广告点击互动内容
}
