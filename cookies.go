package main

import (
	"github.com/valyala/fasthttp"
	"github.com/dgrr/cookiejar"
	"strings"
	"log"
)

func storeInitialCookiesAndRedirect(client *fasthttp.Client, request *fasthttp.Request) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var cookieJar *cookiejar.CookieJar = cookiejar.AcquireCookieJar()
	var NIDAuthentication bool
	defer fasthttp.ReleaseResponse(response)
	
	request.SetRequestURI("https://www.google.com/search?q=test")

	for i := 0; i < 10; i++ {
		client.Do(request, response)
		cookieJar.ReadResponse(response)
		cookieJar.FillRequest(request)
		var statusCode int = response.StatusCode()

		if (statusCode == 200) {
			if (NIDAuthentication) {
				return "SKIP"
			}
			request.Header.Add("Cookie", "SG_SS=")
			NIDAuthentication = true
			continue
		}
		if (statusCode == 302 && strings.Contains(request.Header.String(), "NID=")) {
			return string(response.Header.Peek("Location"))
		}
		log.Printf("[WARNING] Attempt %v: Failed to get inital cookies", i + 1)
		request.Header.DelAllCookies()
		cookieJar.Release()
	}
	log.Printf("[ERROR] Failed to get store inital cookie after 10 attempts")
	return ""
}


func getEnterpriseValue(client *fasthttp.Client, request *fasthttp.Request, location string) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(location)

	client.Do(request, response)
	var body string = string(response.Body())

	if (!strings.Contains(body, "data-s=")) {
		log.Printf("[ERROR] Failed to get captcha enterprise value")
		return ""
	}
	return strings.Split(strings.Split(body, `data-s="`)[1], `"`)[0]
}