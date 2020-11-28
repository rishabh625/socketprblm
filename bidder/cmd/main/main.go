package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	handlers "github.com/rishabh625/socketprblm/bidder/internal/handler"
	object "github.com/rishabh625/socketprblm/bidder/internal/object"
	. "github.com/rishabh625/socketprblm/middleware"
	"strconv"
)

var (
	port         int
	logfile      string
	ver          bool
	bidPath      string
	delayms      string
	host         string
	auctionerurl string
	logger       *log.Logger
)

func init() {
	flag.IntVar(&port, "port", 9080, "The port to listen on.")
	flag.StringVar(&host, "host", "localhost", "Domain")
	flag.StringVar(&logfile, "logfile", "", "Location of the logfile.")
	flag.StringVar(&bidPath, "bidPath", "/bid", "BidPath")
	flag.StringVar(&delayms, "delay.time", "100ms", "Delay Time in ms")
	flag.StringVar(&auctionerurl, "URL", "http://localhost:8080/api/v1/register", "Auctiner URL")
}

func main() {
	fmt.Println("Starting Bidder")

	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}
	handlers.Delaytime = delayms
	strPort := ":" + strconv.Itoa(port)
	http.Handle(bidPath, ServiceLoader(http.HandlerFunc(handlers.Bid), RequestMetrics(logger)))
	RegisterToAuctioner()
	logger.Fatal(http.ListenAndServe(strPort, nil))
}

func RegisterToAuctioner() {
	reqBody, err := json.Marshal(map[string]string{
		"url": fmt.Sprintf("%s:%d%s", host, port, bidPath),
	})

	if err != nil {
		print(err)
	}
	resp, err := http.Post(auctionerurl,
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil || resp == nil {
		panic(err)
	}
	defer resp.Body.Close()
	var regresp object.RegisterBidderResp
	if err := json.NewDecoder(resp.Body).Decode(&regresp); err != nil {
		panic(err)
	} else {
		logger.Printf("Id obtained: %s", regresp.BidderID)
		handlers.RegisterID = regresp.BidderID
	}
}
