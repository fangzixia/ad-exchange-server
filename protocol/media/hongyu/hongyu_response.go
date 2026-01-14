package hongyu

type AdResponse struct {
	// responseid，必须和requestid保持一致
	ID string `json:"id"`
	// 返回代码，0为参与竞价，非0放弃竞价
	Code int `json:"code"`
	// 开发使用的返回信息，协议中不需要
	Message string `json:"message,omitempty"`
	// seatbid列表，如果参与竞价，必须包含一个seatbid对象
	SeatBid []SeatBid `json:"seatBid"`
	// 由DSP生成的responseid，用于logging/tracking
	BidID string `json:"bidID,omitempty"`
	// 扩展信息
	Ext interface{} `json:"ext,omitempty"`
}

// SeatBid 代表DSP端的一个广告主
type SeatBid struct {
	// 竞价信息
	Bid Bid `json:"bid"`
	// 该次竞价是代表谁参与的，一般设置成广告主id，用于logging/tracking
	Seat string `json:"seat,omitempty"`
}

// Bid 广告竞价信息
type Bid struct {
	// 出价类型，目前只支持CPM类型：0–CPM、1–CPC、2–CPA
	BidType int `json:"bidType"`
	// 出价，单位人民币分
	Price float64 `json:"price"`
	// 广告ID
	Adid string `json:"adid,omitempty"`
	// 创意ID，用于loggging/tracking
	Crid string `json:"crid,omitempty"`
	// 拼多多广告ID
	PddAdid string `json:"pddAdid,omitempty"`
	// winnotice地址
	Nurl string `json:"nurl,omitempty"`
	// 竞价失败通知地址，建议支持回传
	Lurl string `json:"lurl,omitempty"`
	// 广告素材
	Creative MaterialMeta `json:"creative"`
	// 如果响应的素材是app下载，则包含该app的信息
	App AppData `json:"app,omitempty"`
	// 点击行为：0-打开网页，1-下载，2-appstore，3-deeplink，4-微信小程序，5-广点通二次下载
	ClickAction int `json:"clickAction"`
	// 点击落地页，如果是app下载则为app下载地址。媒体点击地址会优先使用deeplink，landing次之
	Landing string `json:"landing,omitempty"`
	// 广告来源标识字符
	Source string `json:"source,omitempty"`
	// 广告来源标识的logo图片url
	SourceLogo string `json:"sourceLogo,omitempty"`
	// deeplink地址，如果是调起广告。媒体点击地址会优先使用deeplink，landing次之
	Deeplink string `json:"deeplink,omitempty"`
	// event上报地址，曝光、点击、deeplink唤醒上报等
	EventTrack EventTrack `json:"eventTrack,omitempty"`
	// 广告过期时间，单位秒，针对预加载广告
	ExpirationTime int `json:"expirationTime,omitempty"`
	// 扩展信息
	Ext interface{} `json:"ext,omitempty"`
}

// MaterialMeta 广告素材类
type MaterialMeta struct {
	// 广告标题
	Title string `json:"title,omitempty"`
	// 广告描述
	Description string `json:"description,omitempty"`
	// 广告按钮文字
	Cta string `json:"cta,omitempty"`
	// 评分等级
	Rating string `json:"rating,omitempty"`
	// 广告icon
	Icon Image `json:"icon,omitempty"`
	// 广告图片主体，一张图片
	Image Image `json:"image,omitempty"`
	// 广告图片主体，多张图片
	Images []Image `json:"images,omitempty"`
	// 素材模式：小图=2，大图=3，组图=4，横屏视频=5，竖屏视频=6
	ImageMode int `json:"imageMode,omitempty"`
	// 视频广告资源
	Video Video `json:"video,omitempty"`
	// Banner广告HTML资源
	HtmlSnippet string `json:"htmlSnippet,omitempty"`
	// 小程序广告资源
	MiniProgram MiniProgram `json:"miniProgram,omitempty"`
}

// AppData App下载信息（仅素材为app下载时包含）
type AppData struct {
	// App名称
	Name string `json:"name,omitempty"`
	// Android为packagename，iOS为bundleID
	Bundle string `json:"bundle,omitempty"`
	// App版本号
	Ver string `json:"ver,omitempty"`
	// App图标URL，下载类广告填写
	Icon string `json:"icon,omitempty"`
	// App的介绍，下载类广告填写
	Intro string `json:"intro,omitempty"`
	// 当前版本特征描述，下载类广告填写
	Feature string `json:"feature,omitempty"`
	// 包体大小，单位KB，下载类广告填写
	Size float64 `json:"size,omitempty"`
	// 隐私权限（格式：大标题:介绍;大标题:介绍）
	Privacy string `json:"privacy,omitempty"`
	// 隐私协议，外网可访问链接，下载类广告填写
	PrivacyAgreement string `json:"privacy_agreement,omitempty"`
	// 开发者，下载类广告填写
	Developer string `json:"developer,omitempty"`
	// 付费类型：1-免费版 2-试用版 3-包含付费内容 4-试用版+付费内容
	PaymentType string `json:"payment_type,omitempty"`
}

// Image 图片素材信息
type Image struct {
	// 图片URL地址
	Url string `json:"url,omitempty"`
	// 支持的图片格式类型（image/jpg、image/png等）
	Mimes string `json:"mimes,omitempty"`
	// 图片宽度
	W int `json:"w,omitempty"`
	// 图片高度
	H int `json:"h,omitempty"`
	// 图片宽高比
	AspectRatio float64 `json:"aspectRatio,omitempty"`
}

// Video 视频素材信息
type Video struct {
	// 视频URL地址
	Url string `json:"url,omitempty"`
	// 支持的视频格式类型（video/mp4、video/ogg等）
	Mimes string `json:"mimes,omitempty"`
	// 视频时长，单位：秒
	Duration float32 `json:"duration,omitempty"`
	// 视频大小，单位：KB
	Size float32 `json:"size,omitempty"`
	// 视频分辨率，如1280*720（width*height）
	Resolution string `json:"resolution,omitempty"`
	// 视频封面图片地址，激励视频用于渲染endcard
	Cover string `json:"cover,omitempty"`
	// 封面图宽度，单位像素
	Coverw int `json:"coverw,omitempty"`
	// 封面图高度，单位像素
	Coverh int `json:"coverh,omitempty"`
	// 视频信息流标题
	Title string `json:"title,omitempty"`
	// 视频信息流描述
	Desc string `json:"desc,omitempty"`
	// 视频播放完成后展示的html
	AfterHtml string `json:"afterHtml,omitempty"`
	// 视频强制播放时间（激励视频有效），单位：秒
	ForceDuration int `json:"forceDuration,omitempty"`
	// 视频播放完封面图
	CompleteCoverUrl string `json:"completeCoverUrl,omitempty"`
	// 是否允许跳过：0-不允许，1-允许
	Skip int `json:"skip,omitempty"`
	// 播放多少秒后允许跳过，单位：秒
	SkipMinTime int `json:"skipMinTime,omitempty"`
	// 是否允许客户端预先加载视频广告：0=否，1=是
	Preload int `json:"preload,omitempty"`
	// 允许缓存的最长时间（单位：毫秒，激励视频有效）
	PreloadTtl int `json:"preloadTtl,omitempty"`
}

// MiniProgram 小程序广告资源
type MiniProgram struct {
	// 小程序原始id
	ID string `json:"id,omitempty"`
	// 开放平台的id
	AppId string `json:"appId,omitempty"`
	// 小程序名称
	Name string `json:"name,omitempty"`
	// 小程序路径
	Path string `json:"path,omitempty"`
}

// EventTrack 事件追踪（曝光、点击、deeplink唤醒等）
type EventTrack struct {
	// 曝光监测地址列表
	ImpTracks []string `json:"impTracks,omitempty"`
	// 点击监测地址列表
	ClkTracks []string `json:"clkTracks,omitempty"`
	// 监测App未在手机安装
	AppUninstalled []string `json:"appUninstalled,omitempty"`
	// 监测App已在手机安装
	AppInstalled []string `json:"appInstalled,omitempty"`
	// 尝试调起Deeplink
	DpTry []string `json:"dpTry,omitempty"`
	// Deeplink调用成功上报地址
	DpSuccess []string `json:"dpSuccess,omitempty"`
	// Deeplink调用失败上报地址
	DpFailed []string `json:"dpFailed,omitempty"`
	// 打开landing地址上报
	FallbackTracks []string `json:"fallbackTracks,omitempty"`
	// 下载完成监测URL列表
	DlTracks []string `json:"dlTracks,omitempty"`
	// 安装完成监测URL列表
	InstallTracks []string `json:"installTracks,omitempty"`
	// 下载开始监测URL列表
	DlStartTracks []string `json:"dlStartTracks,omitempty"`
	// 安装开始监测URL列表
	InstallStartTracks []string `json:"installStartTracks,omitempty"`
	// 是否为广点通转化监测：0-否，1-是
	GdtTracks []string `json:"gdtTracks,omitempty"`
	// 播放进度监测链接列表
	ProgressTracks ProgressTrack `json:"progressTracks,omitempty"`
	// 视频自动播放监测上报
	Autoplay []string `json:"autoplay,omitempty"`
	// 视频播放暂停或终止监测上报
	VideoPause []string `json:"videoPause,omitempty"`
	// 视频播放恢复监测上报
	VideoResume []string `json:"videoResume,omitempty"`
	// “跳过播放”监测上报
	SkipTracks []string `json:"skipTracks,omitempty"`
	// “关闭视频”监测上报
	StopTracks []string `json:"stopTracks,omitempty"`
}

// ProgressTrack 视频播放进度追踪
type ProgressTrack struct {
	// 视频开始播放上报链接
	Start []string `json:"start,omitempty"`
	// 视频播放25%上报链接
	FirstQuartile []string `json:"firstQuartile,omitempty"`
	// 视频播放50%上报链接
	MidPoint []string `json:"midPoint,omitempty"`
	// 视频播放75%上报链接
	ThirdQuartile []string `json:"thirdQuartile,omitempty"`
	// 视频播放完成即100%上报链接
	Complete []string `json:"complete,omitempty"`
}
