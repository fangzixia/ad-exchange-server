package adlink

type AdRequest struct {
	RequestID string `json:"requestID"`
	AdSlot    Adslot `json:"adslot"`
	Site      Site   `json:"site"`
	App       App    `json:"app"`
	Device    Device `json:"device"`
	User      User   `json:"user"`
	At        uint   `json:"at"`
	TMax      uint   `json:"tmax"`
}

type Adslot struct{}

type Site struct{}

type App struct{}

type Device struct{}

type User struct{}
