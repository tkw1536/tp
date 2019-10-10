package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var redirectTarget, bindAddress string

func main() {
	// read command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: TARGET=<TARGET> tp <bindAdress>")
		os.Exit(1)
	}
	bindAddress = os.Args[1]
	redirectTarget = os.Getenv("TARGET")

	u, err := url.Parse(redirectTarget)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Proxying %s to %s\n", bindAddress, redirectTarget)

	// and start an http server
	http.ListenAndServe(bindAddress, httputil.NewSingleHostReverseProxy(u))
}
