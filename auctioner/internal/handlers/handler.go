package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rishabh625/socketprblm/auctioner/internal/auction"
	"github.com/rishabh625/socketprblm/auctioner/internal/object"
)

func RegisterBidder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		w.WriteHeader(http.StatusOK)
		var reqobj object.RegisterBidderReq
		var resp *object.RegisterBidderResp
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			resp = &object.RegisterBidderResp{}
		} else {
			resp = &object.RegisterBidderResp{
				BidderID: auction.RegisterBidder(reqobj.BidURL),
			}
		}
		byteresp, _ := json.Marshal(resp)
		w.Write(byteresp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := object.RegisterBidderResp{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}

func Bid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		var reqobj object.StartBidReq
		var resp *object.StartBidResp
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			resp = &object.StartBidResp{}
		} else {
			bidderid, price := auction.StartBid(reqobj.BidID)
			if bidderid == "" {
				w.WriteHeader(http.StatusNoContent)
			}
			resp = &object.StartBidResp{
				BidderID: bidderid,
				Price:    price,
			}
		}
		byteresp, _ := json.Marshal(resp)
		w.Write(byteresp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := object.StartBidResp{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}

func ListBids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		var reqobj object.ListAuctions
		var resp *object.ListAuctionsResp
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			resp = &object.ListAuctionsResp{}
		} else {
			resp = &object.ListAuctionsResp{
				Bids: auction.ListBids(),
			}
		}
		byteresp, _ := json.Marshal(resp)
		w.Write(byteresp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := object.ListAuctionsResp{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}
