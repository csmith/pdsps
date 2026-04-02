package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/csmith/envflag/v2"
)

var (
	upstream = flag.String("upstream", "", "URL of the upstream PDS server")
	port    = flag.Int("port", 8080, "Port to listen on")
)

func main() {
	envflag.Parse()

	if *upstream == "" {
		log.Fatal("upstream is required")
	}

	upstreamURL, err := url.Parse(*upstream)
	if err != nil {
		log.Fatalf("invalid upstream URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(upstreamURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/xrpc/app.bsky.unspecced.getAgeAssuranceState" || r.URL.Path == "/xrpc/app.bsky.ageassurance.getState" {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Headers", "authorization,dpop,atproto-accept-labelers,atproto-proxy")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"state":{"lastInitiatedAt":"2025-07-14T14:22:43.912Z","status":"assured","access":"full"},"metadata":{"accountCreatedAt":"2022-11-17T00:35:16.391Z"}}`))
			return
		}

		proxy.ServeHTTP(w, r)
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Println("Listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
