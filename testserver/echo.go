package main

import (
	"fmt"
	"net/http"
	"os"
)

var redirectTarget, bindAddress string

func main() {
	// read command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: echo <bindAdress>")
		os.Exit(1)
	}
	bindAddress = os.Args[1]
	fmt.Printf("Running echoserver on %s\n", bindAddress)

	// and start an http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path + "\n"))
	})
	http.ListenAndServe(bindAddress, nil)
}
