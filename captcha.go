package main

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"net/url"
	"strings"
	"time"
	"fmt"
)

func createCaptchaTask(location, enterpriseValue string) string {
	var request *fasthttp.Request = createRequest("POST")
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://api.capsolver.com/createTask")
	request.Header.Set("Content-Type", "application/json")
	request.SetBody([]byte(fmt.Sprintf(`{"clientKey":"%s","task":{"type":"ReCaptchaV2Task","websiteURL":"%s","websiteKey":"6LfwuyUTAAAAAOAmoS0fdqijC2PbbdH4kjq62Y1b","enterprisePayload":{"s":"%s"}}}`, CapSolverKey, location, enterpriseValue)))

	fasthttp.Do(request, response)
	
	if taskId := fastjson.GetString(response.Body(), "taskId"); (taskId != "") {
		return taskId
	}
	return ""
}

func getCaptchaResult(taskID string) string {
	var request *fasthttp.Request = createRequest("POST")
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://api.capsolver.com/getTaskResult")
	request.Header.Set("Content-Type", "application/json")
	request.SetBody([]byte(fmt.Sprintf(`{"clientKey":"%s","taskId":"%s"}`, CapSolverKey, taskID)))

	for i := 0 ; i < 30; i++ {
		fasthttp.Do(request, response)

		if token := fastjson.GetString(response.Body(), "solution", "gRecaptchaResponse"); (token != "") {
			return token
		}
		time.Sleep(time.Second * 2)
	}
	return ""
}

func submitCaptcha(client *fasthttp.Client, token, location string) string {
	var request *fasthttp.Request = createRequest("POST")
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	var parts []string = strings.Split(location, "&q=")
	defer fasthttp.ReleaseResponse(response)
	
	request.SetRequestURI("https://www.google.com/sorry/index")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBody([]byte(fmt.Sprintf("g-recaptcha-response=%s&q=%s&continue=%s", token, parts[1], strings.Split(parts[0], "continue=")[1])))

	client.Do(request, response)

	if (response.StatusCode() != 302) {
		return ""
	}
	var unescapedUrl, _ = url.QueryUnescape(string(response.Header.Peek("Location")))
	return unescapedUrl
}