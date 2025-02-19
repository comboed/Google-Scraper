package main

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
	"log"
)

func createCrawler() (*fasthttp.Client, *fasthttp.Request) {
	var client *fasthttp.Client = createClient()
	var request *fasthttp.Request = createRequest("GET")

	for i := 0; i < 10; i++ {
		var location string = preAuthorizeIP(client, request)
		if (location == "") {
			log.Printf("[ERROR] Failed to store IP authorization cookies %s", strings.Repeat(" ", 35))
			break
		}

		var enterpriseValue string = getEnterpriseValue(client, request, location)
		if (enterpriseValue == "") {
			log.Printf("[ERROR] Failed to get captcha enterprise value %s \n", strings.Repeat(" ", 35))
			continue
		}

		var taskId string = createCaptchaTask(location, enterpriseValue)
		if (taskId == "") {
			log.Printf("[ERROR] Failed to create captcha task %s \n", strings.Repeat(" ", 35))
			continue
		}

		var token string = getCaptchaResult(taskId)
		if (token == "") {
			log.Printf("[ERROR] Failed to get captcha token after 30 attempts %s \n", strings.Repeat(" ", 35))
			continue
		}
		
		var abuseUrl string = submitCaptcha(client, token, location)
		if (abuseUrl == "") {
			continue
		}

		request.Header.Add("Cookie", strings.Split(strings.Split(abuseUrl, "google_abuse=")[1], ";")[0])
		return client, request
	}
	return nil, nil
}

func Query(query string, page int) []SearchResult {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	for i := 0; i < 10; i++ {
		var client, request = getRandomCrawler()
		if (request == nil) {
			continue
		}

		request.SetRequestURI("https://www.google.com/search?q=" + query + "&start=" + strconv.Itoa(9 * page))
		if err := client.DoTimeout(request, response, time.Second * 5); (err != nil) {
			log.Println("[ERROR] Failed to fetch results:", err)
			delete(crawlers, client)
			continue
		}

		var body string = string(response.Body())
		if (strings.Contains(body, "SPDX-License-Identifier: Apache-2.0")) {
			delete(crawlers, client)
			continue
		}

		if !strings.Contains(body, `"WEB_RESULT_INNER",["`) {
			continue
		}

		var results []SearchResult
		for _, section := range strings.Split(body, `"WEB_RESULT_INNER",null,"BLUR",0,0,0,null,"`)[1:] {
			var parts []string = strings.Split(section, `"],["`)
			if len(parts) < 2 {
				continue
			}

			results = append(results, SearchResult{
				URL: parts[0],
				Description: strings.Split(parts[1], `","data:image`)[0],
			})
		}
		return results
	}
	return nil
}