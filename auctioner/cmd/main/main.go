package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/rishabh625/socketprblm/auctioner/internal/handlers"
	. "github.com/rishabh625/socketprblm/middleware"
	"strconv"
)

var (
	port    int
	logfile string
	ver     bool
)

func init() {
	flag.IntVar(&port, "port", 8080, "The port to listen on.")
	flag.StringVar(&logfile, "logfile", "", "Location of the logfile.")
	flag.BoolVar(&ver, "version", false, "Print server version.")
}

const (
	// base HTTP paths.
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion + "/"

	//http path .
	register = apiBasePath + "register"
	listbid  = apiBasePath + "listbid"
	startbid = apiBasePath + "startbid"
)

func main() {
	fmt.Println("Starting Auctioner")
	var logger *log.Logger
	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}
	strPort := ":" + strconv.Itoa(port)
	http.Handle(register, ServiceLoader(http.HandlerFunc(handlers.RegisterBidder), RequestMetrics(logger)))
	http.Handle(startbid, ServiceLoader(http.HandlerFunc(handlers.Bid), RequestMetrics(logger)))
	http.Handle(listbid, ServiceLoader(http.HandlerFunc(handlers.ListBids), RequestMetrics(logger)))
	logger.Fatal(http.ListenAndServe(strPort, nil))
}
