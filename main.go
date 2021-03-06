package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/upstream"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	upstreamAddr := flag.String("upstream", "", "upstream address")
	isHTTPS := flag.Bool("https", false, "https upstream")
	flag.Parse()

	if *addr == "" {
		log.Fatal("-addr required")
	}
	if *upstreamAddr == "" {
		log.Fatal("-upstream required")
	}

	var tr http.RoundTripper
	if *isHTTPS {
		tr = &upstream.HTTPSTransport{
			DialTimeout:           time.Second,
			MaxIdleConns:          1000,
			ResponseHeaderTimeout: -1,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	} else {
		tr = &upstream.HTTPTransport{
			DialTimeout:           time.Second,
			MaxIdleConns:          1000,
			ResponseHeaderTimeout: -1,
		}
	}

	svc := parapet.NewBackend()
	svc.Addr = *addr
	svc.H2C = true
	svc.Use(upstream.SingleHost(*upstreamAddr, tr))
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
