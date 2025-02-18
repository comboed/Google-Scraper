package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

func createCrawler() *fasthttp.Request {
	var client *fasthttp.Client = createClient()
	var request *fasthttp.Request = createRequest("GET")

	for i := 0; i < 10; i++ {
		var location string = storeInitialCookiesAndRedirect(client, request)
		if (location == "NO lOCATION") {
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

	// Split this first; then split `,5,null` then `","`` to parse full data
	var data []string = strings.Split(string(response.Body()), "WEB_RESULT_INNER")[1:]
	fmt.Println(data[0])
	fmt.Println(data[1])
}