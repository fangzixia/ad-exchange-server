package log

type Res struct {
	Status        int     `json:"status"`
	DispatchTimes int     `json:"dispatchTimes"`
	Duration      int     `json:"duration"`
	PlatformId    int     `json:"platformId"`
	ConfigId      int     `json:"configId"`
	FinalPrice    float64 `json:"finalPrice"`
	FloorPrice    float64 `json:"floorPrice"`
}

// AdSpace 广告位信息结构体
type AdSpace struct {
	RequestAdSpaceId string `json:"requestAdSpaceId"`
	AdSpaceId        int    `json:"adSpaceId"`
	AdFormat         int    `json:"adFormat"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	BidFloor         int    `json:"bidFloor"`
	// 其他字段可按需补充
}

// Device 设备信息结构体
type Device struct {
	MUniqueType int    `json:"mUniqueType"`
	MUniqueId   string `json:"mUniqueId"`
	Ua          string `json:"ua"`
	Ip          string `json:"ip"`
	Make        string `json:"make"`
	Model       string `json:"model"`
	Brand       string `json:"brand"`
	Os          int    `json:"os"`
	OsVersion   string `json:"osVersion"`
	// 其他字段可按需补充
}

// Req 请求数据结构体
type Req struct {
	Timestamp     int64   `json:"timestamp"`
	ServerId      string  `json:"serverId"`
	RemoteIp      string  `json:"remoteIp"`
	ChannelId     int     `json:"channelId"`
	ChannelType   string  `json:"channelType"`
	TransactionID string  `json:"transactionID"`
	AdSpace       AdSpace `json:"adSpace"`
	Device        Device  `json:"device"`
	// 其他字段（geoBlock、appBlock等）可按需补充
}

// LogData 完整日志JSON数据结构体
type LogData struct {
	Res Res `json:"res"`
	Req Req `json:"req"`
}

// LogItem 完整日志项（包含前缀标识和JSON数据）
type MediaLog struct {
	UniqueID string  `json:"unique_id"`
	Data     LogData `json:"data"`
}
