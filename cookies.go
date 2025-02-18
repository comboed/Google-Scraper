package main

import (
	// "fmt"
	"github.com/valyala/fasthttp"
	"github.com/dgrr/cookiejar"
	"strings"
	"log"
)

func preAuthorizeIP(client *fasthttp.Client, request *fasthttp.Request) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var cookieJar *cookiejar.CookieJar = cookiejar.AcquireCookieJar()
	defer fasthttp.ReleaseResponse(response)
	
	request.SetRequestURI("https://www.google.com/search?q=test")
	request.Header.Set("Cookie", "SG_SS=")

	for i := 0; i < 10; i++ {
		client.Do(request, response)
		cookieJar.ReadResponse(response)
		cookieJar.FillRequest(request)

		if (strings.Contains(string(response.Body()), "302 Moved") && strings.Contains(request.Header.String(), "NID=")) {
			return string(response.Header.Peek("Location"))
		}
		log.Printf("[WARNING] Attempt %v: Failed to get inital cookies \n %s", i + 1, strings.Repeat(" ", 35))
		request.Header.DelAllCookies()
		cookieJar.Release()
	}
	log.Printf("[ERROR] Failed to get store inital cookie after 10 attempts %s", strings.Repeat(" ", 35))
	return ""
}


func getEnterpriseValue(client *fasthttp.Client, request *fasthttp.Request, location string) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(location)

	client.Do(request, response)
	var body string = string(response.Body())

	if (!strings.Contains(body, "data-s=")) {
		log.Printf("[ERROR] Failed to get captcha enterprise value %s \n", strings.Repeat(" ", 35))
		return ""
	}
	return strings.Split(strings.Split(body, `data-s="`)[1], `"`)[0]
}