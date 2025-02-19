package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/valyala/fasthttp"
)

func openFile(filename string) (slice []string) {
	var file, _ = os.Open(filename)
	var scan *bufio.Scanner = bufio.NewScanner(file)
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	file.Close()
	return slice
}

func getRandomCrawler() (*fasthttp.Client, *fasthttp.Request) {
	if len(crawlers) < MaxCrawlers {
		for {
			var client, request = createCrawler()
			if (client == nil || request == nil) {
				fmt.Println("Nooo")
				continue
			}
			crawlers[client] = request
			return client, request
		}
	}
	var clients []*fasthttp.Client
	for key := range crawlers {
		clients = append(clients, key)
	}
	var client *fasthttp.Client = clients[rand.Intn(len(clients))]
	return client, crawlers[client]
}
