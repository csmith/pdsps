package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/csmith/envflag/v2"
	"github.com/csmith/slogflags"
)

var (
	upstream = flag.String("upstream", "", "URL of the upstream PDS server")
	port     = flag.Int("port", 8080, "Port to listen on")
)

func main() {
	envflag.Parse()
	_ = slogflags.Logger(slogflags.WithSetDefault(true))

	if *upstream == "" {
		slog.Error("upstream is required")
		os.Exit(1)
	}

	upstreamURL, err := url.Parse(*upstream)
	if err != nil {
		slog.Error("Invalid upstream URL", "error", err)
		os.Exit(1)
	}

	proxy := httputil.NewSingleHostReverseProxy(upstreamURL)

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/xrpc/app.bsky.unspecced.getAgeAssuranceState" ||
				r.URL.Path == "/xrpc/app.bsky.ageassurance.getState" {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set(
					"Access-Control-Allow-Headers",
					"authorization,dpop,atproto-accept-labelers,atproto-proxy,x-atproto-device-id,x-atproto-session-id",
				)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write(
					[]byte(
						`{"state":{"lastInitiatedAt":"2025-07-14T14:22:43.912Z","status":"assured","access":"full"},` +
							`"metadata":{"accountCreatedAt":"2022-11-17T00:35:16.391Z"}}`,
					),
				)
				return
			}

			proxy.ServeHTTP(w, r)
		},
	)

	addr := fmt.Sprintf(":%d", *port)
	slog.Info("Listening", "addr", addr)
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
