package adlink

type AdResponse struct {
	ID      string     `json:"id"`
	Code    uint8      `json:"code"`
	SeatBid []*SeatBid `json:"seatBid"`
	BidID   string     `json:"bidID"`
}

type SeatBid struct{}
