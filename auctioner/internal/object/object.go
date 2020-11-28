package object

type RegisterBidderReq struct {
	BidURL string `json:"url"`
}

type RegisterBidderResp struct {
	BidderID string `json:"bidder_id,omitempty"`
}

type ListAuctions struct {
}

type ListAuctionsResp struct {
	Bids []string `json:"bidders"`
}

type StartBidReq struct {
	BidID string `json:"auction_id"`
}

type StartBidResp struct {
	BidderID string  `json:"bidder_id,omitempty"`
	Price    float64 `json:"price,omitempty"`
}
