package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/moonrhythm/parapet"
	"github.com/moonrhythm/parapet/pkg/upstream"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	target := flag.String("target", "", "target address")
	flag.Parse()

	if *addr == "" {
		log.Fatal("-addr required")
	}
	if *target == "" {
		log.Fatal("-upstream required")
	}

	var tr http.RoundTripper
	switch {
	case strings.HasPrefix(*target, "unix://"):
		*target = strings.TrimPrefix(*target, "unix://")
		tr = &upstream.UnixTransport{
			MaxIdleConns:          5000,
			ResponseHeaderTimeout: -1,
		}
	case strings.HasPrefix(*target, "https://"):
		*target = strings.TrimPrefix(*target, "https://")
		*target = injectDefaultPort(*target, "443")
		tr = &upstream.HTTPSTransport{
			DialTimeout:           time.Second,
			MaxIdleConns:          1000,
			ResponseHeaderTimeout: -1,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	case strings.HasPrefix(*target, "h2c://"):
		*target = strings.TrimPrefix(*target, "h2c://")
		*target = injectDefaultPort(*target, "80")
		tr = &upstream.H2CTransport{}
	default:
		*target = strings.TrimPrefix(*target, "http://")
		*target = injectDefaultPort(*target, "80")
		tr = &upstream.HTTPTransport{
			DialTimeout:           time.Second,
			MaxIdleConns:          1000,
			ResponseHeaderTimeout: -1,
		}
	}

	svc := parapet.NewBackend()
	svc.Addr = *addr
	svc.H2C = true
	svc.Use(upstream.SingleHost(*target, tr))
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func injectDefaultPort(addr string, port string) string {
	if strings.Contains(addr, ":") {
		return addr
	}
	return addr + ":" + port
}
