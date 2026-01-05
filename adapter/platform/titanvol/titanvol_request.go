package titanvol

type AdRequest struct {
	Tid        string       `json:"tid"`
	Timestamp  int64        `json:"timestamp"`
	Impression []Impression `json:"impression"`
	Device     *Device      `json:"device,omitempty"`
	App        *App         `json:"app,omitempty"`
	User       *User        `json:"user,omitempty"`
}

// Impression 嵌套结构体（原Java静态内部类）
type Impression struct {
	Id           string   `json:"id"`
	Aid          string   `json:"aid"`
	AdFormat     int      `json:"adFormat"`
	BidFloor     int      `json:"bidFloor"`
	Bidcur       string   `json:"bidcur"`
	BidType      int      `json:"bidType"`
	W            int      `json:"w"`
	H            int      `json:"h"`
	Wmax         int      `json:"wmax"`
	Hmax         int      `json:"hmax"`
	Wmin         int      `json:"wmin"`
	Hmin         int      `json:"hmin"`
	Mimes        []string `json:"mimes,omitempty"` // Java String[] -> Go []string
	Skip         int      `json:"skip"`
	SkipAfter    int      `json:"skipAfter"`
	VideoType    int      `json:"videoType"`
	SupportWx    int      `json:"supportWx"`
	SupportVideo int      `json:"supportVideo"`
	Support302   int      `json:"support302"`
	SupportDown  int      `json:"supportDown"`
	Https        int      `json:"https"`
	Deeplink     int      `json:"deeplink"`
	BlockCat     []string `json:"blockCat,omitempty"`
	Size         int      `json:"size"`
}

// Geo 嵌套结构体
type Geo struct {
	Lat           float32 `json:"lat"` // Java float -> Go float32
	Lon           float32 `json:"lon"`
	Type          int     `json:"type"`
	Country       string  `json:"country"`
	Province      string  `json:"province"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	CoordTime     int     `json:"coordTime"`
	LocationAccur float32 `json:"locationAccur"`
}

// Caid 嵌套结构体
type Caid struct {
	Caid    string `json:"caid"`
	Version string `json:"version"`
}

// Device 嵌套结构体
type Device struct {
	Ua                string   `json:"ua"`
	Geo               *Geo     `json:"geo,omitempty"`
	Ip                string   `json:"ip"`
	Ipv6              string   `json:"ipv6"`
	DeviceType        int      `json:"deviceType"`
	Make              string   `json:"make"`
	Model             string   `json:"model"`
	Os                int      `json:"os"`
	Osv               string   `json:"osv"`
	Carrier           int      `json:"carrier"`
	SerialNo          string   `json:"serialno"`
	Meid              string   `json:"meid"`
	ConnectionType    int      `json:"connectionType"`
	Imsi              string   `json:"imsi"`
	Imei              string   `json:"imei"`
	ImeiMd5           string   `json:"imeiMd5"`
	ImeiSha1          string   `json:"imeiSha1"`
	Oaid              string   `json:"oaid"`
	OaidMd5           string   `json:"oaidMd5"`
	Adid              string   `json:"adid"`
	AdidMd5           string   `json:"adidMd5"`
	Idfa              string   `json:"idfa"`
	IdfaMd5           string   `json:"idfaMd5"`
	IdfaSha1          string   `json:"idfaSha1"`
	Idfv              string   `json:"idfv"`
	IdfvMd5           string   `json:"idfvMd5"`
	Mac               string   `json:"mac"`
	MacMd5            string   `json:"macMd5"`
	MacSha1           string   `json:"macSha1"`
	W                 int      `json:"w"`
	H                 int      `json:"h"`
	Dpi               string   `json:"dpi"`
	Ppi               string   `json:"ppi"`
	Density           float32  `json:"density"`
	Inch              float32  `json:"inch"`
	CaidList          []Caid   `json:"caidList,omitempty"`
	Aaid              string   `json:"aaid"`
	BootMark          string   `json:"bootMark"`
	UpdateMark        string   `json:"updateMark"`
	DeviceRomVer      string   `json:"deviceRomVer"`
	Hmscore           string   `json:"hmscore"`
	Paid              string   `json:"paid"`
	Paid14            string   `json:"paid_1_4"` // 原字段paid_1_4，Go字段驼峰+JSON tag
	ScreenOrientation int      `json:"screenOrientation"`
	Applist           []string `json:"applist,omitempty"`
	HmsVer            string   `json:"hmsVer"`
	HwagVer           string   `json:"hwagVer"`
	DeviceName        string   `json:"deviceName"`
	DeviceNameMd5     string   `json:"deviceNameMd5"`
	TimeZone          string   `json:"timeZone"`
	LocalTzName       string   `json:"localTzName"`
	AppStoreVer       string   `json:"appStoreVer"`
	AppStoreUrl       string   `json:"appStoreUrl"`
	ApiLevel          string   `json:"apiLevel"`
	Ssid              string   `json:"ssid"`
	Bssid             string   `json:"bssid"`
	MiuiVer           string   `json:"miuiVer"`
	AuthStatus        int      `json:"authStatus"`
	Openudid          string   `json:"openudid"`
	BootTime          string   `json:"bootTime"`
	BirthTime         string   `json:"birthTime"`
	UpdateTime        string   `json:"updateTime"`
	ElapseTime        int      `json:"elapseTime"`
	SysCompileTime    string   `json:"sysCompileTime"`
	MemorySize        string   `json:"memorySize"`
	DiskSize          string   `json:"diskSize"`
	DiskFreeSpace     string   `json:"diskFreeSpace"`
	BatteryStatus     int      `json:"batteryStatus"`
	BatteryPower      int      `json:"batteryPower"`
	CpuNum            int      `json:"cpuNum"`
	CpuFreq           float32  `json:"cpuFreq"`
	HwModel           string   `json:"hwModel"`
	HwName            string   `json:"hwName"`
	HwMachine         string   `json:"hwMachine"`
	HwVersion         string   `json:"hwVersion"`
	AndroidSha1       string   `json:"androidSha1"`
}

// App 嵌套结构体
type App struct {
	Name        string `json:"name"`
	Bundle      string `json:"bundle"`
	Version     string `json:"version"`
	StoreUrl    string `json:"storeUrl"`
	Cat         string `json:"cat"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Paid        int    `json:"paid"`
}

// User 嵌套结构体
type User struct {
	UserId   string `json:"userId"`
	Tags     string `json:"tags"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age"`
	Keywords string `json:"keywords"`
	Yob      int    `json:"yob"`
}
