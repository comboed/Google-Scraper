package main

import (
	"github.com/valyala/fasthttp/fasthttpproxy"
	"github.com/valyala/fasthttp"
	"math/rand"
	"crypto/tls"
	"time"
	"net"
)

func createClient(channel chan string) *fasthttp.Client {
	return &fasthttp.Client {
		MaxConnsPerHost: 1,
		ReadBufferSize: 4096 * 3,
		DisablePathNormalizing: true,
		NoDefaultUserAgentHeader: true,
		DisableHeaderNamesNormalizing: true,
		TLSConfig: &tls.Config {
			InsecureSkipVerify: true,
			MinVersion: tls.VersionTLS13,
			MaxVersion: tls.VersionTLS13,
		},
		Dial: func(addr string) (net.Conn, error) {
			return fasthttpproxy.FasthttpHTTPDialerTimeout(proxies[rand.Intn(len(proxies))], time.Minute)(addr)
		},
	}

}

func createRequest() *fasthttp.Request {
	var request *fasthttp.Request = fasthttp.AcquireRequest()

	request.Header.SetMethod("GET")

	for key, value := range getHeaders() {
        request.Header.Set(key, value)
    }

	return request
}

func getHeaders() map[string]string {
	return map[string]string {
		"Host": "www.google.com",
        "Connection": "keep-alive",
        "rtt": "0",
        "downlink": "10",
        "sec-ch-ua": `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`,
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-full-version": `"133.0.6943.98"`,
        "sec-ch-ua-arch": `"x86"`,
        "sec-ch-ua-platform": `"Windows"`,
        "sec-ch-ua-platform-version": `"15.0.0"`,
        "sec-ch-ua-model": `""`,
        "sec-ch-ua-bitness": `"64"`,
        "sec-ch-ua-wow64": "?0",
        "sec-ch-ua-full-version-list": `"Not(A:Brand";v="99.0.0.0", "Google Chrome";v="133.0.6943.98", "Chromium";v="133.0.6943.98"`,
        "sec-ch-ua-form-factors": `"Desktop"`,
        "sec-ch-prefers-color-scheme": "dark",
        "Upgrade-Insecure-Requests": "1",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "X-Browser-Channel": "stable",
        "X-Browser-Year": "2025",
        "X-Browser-Validation": "1nAW9Rb/M8Lkk97ILDg00FWYjns=",
        "X-Browser-Copyright": "Copyright 2025 Google LLC. All rights reserved.",
        "Sec-Fetch-Site": "none",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-User": "?1",
        "Sec-Fetch-Dest": "document",
        "Accept-Language": "en-US,en;q=0.9",
	}
}