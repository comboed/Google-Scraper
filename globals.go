package main

import "github.com/valyala/fasthttp"

const (
	CapSolverKey string = "capsolver-key"
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