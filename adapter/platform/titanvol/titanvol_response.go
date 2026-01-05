package titanvol

// BidUpperResponseDTO 对应 Java 中的 BidUpperResponseDTO 主类
type AdResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Tid     string    `json:"tid"`
	SeatBid []SeatBid `json:"seatBid"` // Java 数组转 Go 切片
	BidId   string    `json:"bidId"`
}

// SeatBid 对应 Java 中 SeatBid 静态内部类
type SeatBid struct {
	Bid []Bid `json:"bid"` // Java 数组转 Go 切片
}

// Bid 对应 Java 中 Bid 静态内部类
type Bid struct {
	Id          string                 `json:"id"`
	ImpId       string                 `json:"impId"`
	Expire      int                    `json:"expire"`
	Creative    Creative               `json:"creative"`
	AdverId     string                 `json:"adverId"`
	AdverName   string                 `json:"adverName"`
	Vocation    int                    `json:"vocation"`
	Price       int                    `json:"price"`
	Nurl        string                 `json:"nurl"`
	Lurl        string                 `json:"lurl"`
	ClickAction int                    `json:"clickAction"`
	Landing     string                 `json:"landing"`
	Source      string                 `json:"source"`
	SourceLogo  string                 `json:"sourceLogo"`
	DeepLink    string                 `json:"deepLink"`
	DeepUlink   string                 `json:"deepUlink"`
	Isgdt       int                    `json:"isgdt"`
	App         AppInfo                `json:"app"`
	MiniProgram MiniProgram            `json:"miniProgram"` // Java 中仅声明未定义字段，先空结构体
	Trackings   []Tracking             `json:"trackings"`   // Java List 转 Go 切片
	Ext         map[string]interface{} `json:"ext"`         // Java Map<String, Object> 转 Go map
}

// Event 对应 Java 中 Event 静态内部类（含常量+字段）
type Event struct {
	EventKey   int    `json:"event_key"` // Java 下划线字段，JSON 标签保持一致
	EventValue string `json:"event_value"`
}

// Creative 对应 Java 中 Creative 静态内部类
type Creative struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cta         string   `json:"cta"`
	Rating      float32  `json:"rating"` // Java float 对应 Go float32
	Icon        Image    `json:"icon"`
	Images      []Image  `json:"Images"` // Java 数组转切片，JSON 标签与 Java 字段名（Images）一致
	ImageMode   int      `json:"imageMode"`
	Video       Video    `json:"video"`
	Mimes       []string `json:"mimes"` // Java List<String> 转 Go 字符串切片
}

// Image 对应 Java 中 Image 静态内部类
type Image struct {
	Width  int      `json:"width"`
	Height int      `json:"height"`
	Url    string   `json:"url"`
	Size   int      `json:"size"`
	Ratio  float32  `json:"ratio"` // Java float 对应 Go float32
	Mimes  []string `json:"mimes"` // Java 数组转 Go 字符串切片
}

// Video 对应 Java 中 Video 静态内部类
type Video struct {
	Type         int             `json:"type"`
	Title        string          `json:"title"`
	Desc         string          `json:"desc"`
	Width        int             `json:"width"`
	Height       int             `json:"height"`
	Url          string          `json:"url"`
	CoverUrl     string          `json:"coverUrl"`
	CoverWidth   int             `json:"coverWidth"`
	CoverHeight  int             `json:"coverHeight"`
	EndCoverUrl  string          `json:"endCoverUrl"`
	Duration     int             `json:"duration"`
	Size         int             `json:"size"`
	Resolution   string          `json:"resolution"`
	MinDuration  int             `json:"minDuration"`
	Afterhtml    string          `json:"afterhtml"`
	AfterBtnText string          `json:"afterBtnText"`
	ValidTime    int             `json:"validTime"`
	AfterBtnUrl  string          `json:"afterBtnUrl"`
	PlayType     int             `json:"playType"`
	Prefetch     int             `json:"prefetch"`
	Mimes        []string        `json:"mimes"`     // Java 数组转 Go 字符串切片
	Trackings    []PointTracking `json:"trackings"` // Java List 转 Go 切片
}

// Tracking 对应 Java 中 Tracking 静态内部类
type Tracking struct {
	EventType int      `json:"eventType"`
	Urls      []string `json:"urls"` // Java 数组转 Go 字符串切片
}

// PointTracking 对应 Java 中 PointTracking 静态内部类
type PointTracking struct {
	Ts   int      `json:"ts"`
	Urls []string `json:"urls"` // Java 数组转 Go 字符串切片
}

// App 对应 Java 中 App 静态内部类（字段按 Java 完整映射）
type AppInfo struct {
	Name              string   `json:"name"`
	Bundle            string   `json:"bundle"`
	ItunesId          string   `json:"itunesId"`
	MarketUrl         string   `json:"marketUrl"`
	Relation          int      `json:"relation"`
	IconUrl           string   `json:"iconUrl"`
	IconW             int      `json:"iconW"`
	IconH             int      `json:"iconH"`
	Version           string   `json:"version"`
	VersionCode       int      `json:"versionCode"`
	UpdateTime        string   `json:"updateTime"`
	QuickUrl          string   `json:"quickUrl"`
	DownUrl           string   `json:"downUrl"`
	Md5               string   `json:"md5"`
	Developer         string   `json:"developer"`
	Intro             string   `json:"intro"`
	IntroLink         string   `json:"introLink"`
	PrivatePolicy     string   `json:"privatePolicy"`
	PrivatePolicyLink string   `json:"privatePolicyLink"`
	Permissions       string   `json:"permissions"`
	PermissionsLink   string   `json:"permissionsLink"`
	Size              int      `json:"size"`
	Rating            float32  `json:"rating"` // Java float 对应 Go float32
	Downloads         int      `json:"downloads"`
	Comments          int      `json:"comments"`
	Snapshots         []string `json:"snapshots"` // Java 数组转 Go 字符串切片
	Desc              string   `json:"desc"`
	Tag               []string `json:"tag"` // Java 数组转 Go 字符串切片
	Icp               string   `json:"icp"`
}

type MiniProgram struct {
	Id    string `json:"id"`
	Path  string `json:"path"`
	Name  string `json:"name"`
	AppId string `json:"appId"`
}
