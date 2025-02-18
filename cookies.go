package main

import (
	"github.com/valyala/fasthttp"
	"github.com/dgrr/cookiejar"
	"strings"
	"log"
	"os"
)

func storeInitialCookies(client *fasthttp.Client, request *fasthttp.Request) bool {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var cookieJar *cookiejar.CookieJar = cookiejar.AcquireCookieJar()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://www.google.com/search")

	for i := 0; i < 10; i++ {
		client.Do(request, response)
		cookieJar.ReadResponse(response)
		cookieJar.FillRequest(request)
		var header string = request.Header.String()
		
		if strings.Contains(header, "AEC=") && strings.Contains(header, "NID=") {
			return response.StatusCode() != 302 // Will tell us if we have a captcha
		}
		log.Printf("[%d] Failed to get initial cookies, retrying", i)
		request.Header.DelAllCookies()
		cookieJar.Release()
	}
	log.Printf("Failed to get store inital cookie after 10 attempts")
	os.Exit(0)
	return false
}

func getExemptionCookie(client *fasthttp.Client, request *fasthttp.Request, query string) bool {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var cookieJar *cookiejar.CookieJar = cookiejar.AcquireCookieJar()
	defer fasthttp.ReleaseResponse(response)

	request.Header.SetMethod("GET")
	request.SetRequestURI(query)

	client.Do(request, response)
	cookieJar.ReadResponse(response)
	cookieJar.ReadRequest(request)

	return strings.Contains(request.Header.String(), "GOOGLE_ABUSE_EXEMPTION")
}