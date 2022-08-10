package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var redirectTarget, rootTarget, bindAddress string

func main() {
	// read command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: TARGET=<TARGET> tp <bindAdress>")
		os.Exit(1)
	}
	bindAddress = os.Args[1]
	redirectTarget = os.Getenv("TARGET")
	rootTarget = os.Getenv("ROOT")

	// parse the redirect target
	u, err := url.Parse(redirectTarget)
	if err != nil {
		panic(err)
	}

	// create a handler, and override the root if requested
	var handler http.Handler = httputil.NewSingleHostReverseProxy(u)
	if rootTarget != "" {
		handler = RootDirector{
			Target:  rootTarget,
			Handler: handler,
		}
	}

	// and start an http server
	fmt.Printf("Proxying %s to %s\n", bindAddress, redirectTarget)
	http.ListenAndServe(bindAddress, handler)
}

// RootDirector is an [http.Handler] that wraps an underlying [http.Handler], but overrides the root URL
type RootDirector struct {
	// Target is the target for the root directory
	Target string

	// Handler is the original [http.Handler] being wrapped
	Handler http.Handler
}

func (ro RootDirector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// root URL performs a redirect
	if r.URL.Path == "" || r.URL.Path == "/" {
		http.Redirect(w, r, ro.Target, http.StatusFound)
		return
	}

	// everything else serves normally
	ro.Handler.ServeHTTP(w, r)
}
