package main

import (
	"flag"
	"net/http"
	. "github.com/rishabh625/socketprblm/internal/middleware"
)

const (
	// base HTTP paths.
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion + "/"

	//http path .
	startPath = apiBasePath + "signup"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	errs := make(chan error)
	http.Handle(startPath, ServiceLoader(http.HandlerFunc(handlers.Signup), RequestMetrics(logger)))

	go func() {
		errs <- http.ListenAndServe(*httpAddr, httpHandler)
	}()
}
