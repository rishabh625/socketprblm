package object

type RegisterBidderReq struct {
	BidURL string `json:"url"`
}

type RegisterBidderResp struct {
	BidderID string `json:"bidder_id,omitempty"`
}

type StartBidResp struct {
	BidderID string  `json:"bidder_id,omitempty"`
	Price    float64 `json:"price,omitempty"`
}
