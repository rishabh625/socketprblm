package auction

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/rishabh625/socketprblm/auctioner/libs/object"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var bidmap map[string]map[string]float64
var registeredbidder map[string]string
var bidderidcount int

func init() {
	bidmap = make(map[string]map[string]float64)
	registeredbidder = make(map[string]string)
}

func RegisterBidder(url string) string {
	bidderidcount++
	count := strconv.Itoa(bidderidcount)
	key := fmt.Sprintf("%s%s", strconv.FormatInt(time.Now().UnixNano(), 10), RandStringRunes(rand.Intn(20))+count)
	bidderid := fmt.Sprintf("%x", md5.Sum([]byte(key)))
	registeredbidder[bidderid] = fmt.Sprintf("%s", url)
	return bidderid
}

func ListBids() []string {
	var list []string
	for indx, url := range registeredbidder {
		list = append(list, fmt.Sprintf("Id:%s, URL:%s", indx, url))
	}
	return list
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func StartBid(auctionid string) (string, float64) {
	if bidmap[auctionid] == nil {
		bidmap[auctionid] = make(map[string]float64)
	}
	for id, url := range registeredbidder {
		go Query(url, auctionid, id)
	}
	time.Sleep(200 * time.Millisecond)
	var respbidderID string
	var price float64
	for bidderID, bid := range bidmap[auctionid] {
		if bid > price {
			respbidderID = bidderID
			price = bid
		}
	}
	return respbidderID, price
}

func Query(url, auctionID, bidderID string) {
	resp, err := http.Get("http://" + url + "?auction_id=" + auctionID)
	if err != nil || resp == nil {
		delete(registeredbidder, bidderID)
		return
	}
	defer resp.Body.Close()
	var bidresp object.StartBidResp
	if err := json.NewDecoder(resp.Body).Decode(&bidresp); err != nil {
		delete(registeredbidder, bidderID)
	} else {
		if bidderID == bidresp.BidderID {
			bidmap[auctionID][bidderID] = bidresp.Price
		}
	}
}
