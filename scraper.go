package main

import (
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"log"
	_"fmt"
)

func createCrawler() *fasthttp.Request {
	var client *fasthttp.Client = createClient()
	var request *fasthttp.Request = createRequest("GET")

	for i := 0; i < 10; i++ {
		var location string = preAuthorizeIP(client, request)
		if (location == "NO LOCATION") {
			return request
		}
		if (location == "") {
			log.Printf("[ERROR] Failed to get store IP authorization cookies %s", strings.Repeat(" ", 35))
			continue
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
		return request
	}
	return nil
}

func Query(query string, page int) []SearchResult {
	var client *fasthttp.Client = createClient()
	var request *fasthttp.Request = createCrawler()
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://www.google.com/search?q=" + query + "&start=" + strconv.Itoa(9 * page))

	client.Do(request, response)
	var body string = string(response.Body())
	
	var results []SearchResult
	if (strings.Contains(body, `"WEB_RESULT_INNER",["`)) {
		var sections []string = strings.Split(body, `"WEB_RESULT_INNER",null,"BLUR",0,0,0,null,"`)[1:]

		for i := range sections {
			var parts []string = strings.Split(sections[i], `"],["`)
			results = append(results, SearchResult{ URL: parts[0], Description: strings.Split(parts[1], `","data:image`)[0],})
		}
	}
	return results
}