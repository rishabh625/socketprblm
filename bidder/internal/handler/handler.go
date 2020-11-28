package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"github.com/rishabh625/socketprblm/bidder/internal/object"
	"time"
)

var Delaytime string
var RegisterID string

func Bid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		d, _ := time.ParseDuration(Delaytime)
		time.Sleep(d)
		rand.Seed(time.Now().UnixNano())
		resp := &object.StartBidResp{
			BidderID: RegisterID,
			Price:    rand.Float64(),
		}
		fmt.Println("My Bid: ", resp.Price)
		byteresp, _ := json.Marshal(resp)
		w.Write(byteresp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := object.StartBidResp{}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}
