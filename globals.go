package main

import "github.com/valyala/fasthttp"

const (
	CapSolverKey string = "CAP-B367787B6A0E5DEEE4186194F7C81372"
	MaxCrawlers int = 1
	MaxCount int = 27
)

type SearchResult struct {
	URL string `json:"url"`
	Description string `json:"description"`
}

var (
	proxies []string
	crawlers map[*fasthttp.Client]*fasthttp.Request
)