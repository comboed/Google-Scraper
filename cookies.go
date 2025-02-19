package main

import (
	"fmt"
	"strings"

	"github.com/dgrr/cookiejar"
	"github.com/valyala/fasthttp"
)

func preAuthorizeIP(client *fasthttp.Client, request *fasthttp.Request) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var cookieJar *cookiejar.CookieJar = cookiejar.AcquireCookieJar()
	defer fasthttp.ReleaseResponse(response)
	
	request.SetRequestURI("https://www.google.com/search?q=test")
	request.Header.Set("Cookie", "SG_SS=") // Forces captcha

	client.Do(request, response)
	cookieJar.ReadResponse(response)
	cookieJar.FillRequest(request)

	if (!strings.Contains(string(response.Body()), "302 Moved")) {
		fmt.Println(string(response.Body()))
		return ""
	}
	return string(response.Header.Peek("Location"))
}

func getEnterpriseValue(client *fasthttp.Client, request *fasthttp.Request, location string) string {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI(location)

	client.Do(request, response)
	var body string = string(response.Body())

	if (!strings.Contains(body, "data-s=")) {
		return ""
	}
	return strings.Split(strings.Split(body, `data-s="`)[1], `"`)[0]
}