package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {

	var (
		signedURL string
		method    string
		help      bool
	)

	flag.StringVar(&signedURL, "url", "", "Signed URL")
	flag.StringVar(&method, "method", "", "HTTP method")
	flag.BoolVar(&help, "h", false, "help")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	URL, err := url.Parse(signedURL)
	if err != nil {
		log.Fatal(err)
	}

	host := strings.Split(URL.Host, ":")[0]

	nodes, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	for _, node := range nodes {
		// This could be less hacky with https://github.com/golang/go/issues/12503
		transport := &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				addr = node + ":" + strings.Split(addr, ":")[1]
				return dialer.DialContext(ctx, network, addr)
			},
		}

		c := &http.Client{Timeout: 10 * time.Second, Transport: transport}

		req, err := http.NewRequest(method, signedURL, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := c.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("HTTP status:", resp.StatusCode)
		for k, v := range resp.Header {
			fmt.Println(k, "=>", v)
		}

		resp.Body.Close()
	}

}
