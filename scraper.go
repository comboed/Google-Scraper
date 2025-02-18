package main

import (
	"github.com/valyala/fasthttp"
	"strings"
	"strconv"
	"fmt"
	"os"
)

func createCrawler() *fasthttp.Request {
	var client *fasthttp.Client = createClient()
	var request *fasthttp.Request = createRequest("GET")

	for i := 0; i < 10; i++ {
		var location string = storeInitialCookiesAndRedirect(client, request)
		if (location == "SKIP") {
			return request
		}
		if (location == "") {
			continue
		}

		var enterpriseValue string = getEnterpriseValue(client, request, location)
		if (enterpriseValue == "") {
			continue
		}

		var taskId string = createCaptchaTask(location, enterpriseValue)
		if (taskId == "") {
			continue
		}

		var token string = getCaptchaResult(taskId)
		if (token == "") {
			continue
		}
		
		var abuseUrl string = submitCaptcha(client, token, location)
		if (abuseUrl == "") {
			continue
		}

		request.Header.Add("Cookie", strings.Split(strings.Split(abuseUrl, "google_abuse=")[1], ";")[0])
		return request
	}
	os.Exit(0)
	return nil
}

func run(client *fasthttp.Client, channel chan string) {
	var request *fasthttp.Request = createCrawler()
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var count int = 9
	
	//var query string = <- channel
	request.SetRequestURI("https://www.google.com/search?q=" + "test" + "&start=" + strconv.Itoa(count))
	client.Do(request, response)
	fmt.Println(response.String())
		
}