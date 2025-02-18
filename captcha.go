package main

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
	"strings"
	"strconv"
	"time"
	"fmt"
	"log"
)

func createCaptchaTask(query string) string {
	var request *fasthttp.Request = createRequest("POST")
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://api.capmonster.cloud/createTask")

	request.Header.Set("Content-Type", "application/json")
	request.SetBody([]byte(fmt.Sprintf(`{"clientKey": "%s","task":{"type": "RecaptchaV2TaskProxyless","websiteURL": "%s","websiteKey": "6LfwuyUTAAAAAOAmoS0fdqijC2PbbdH4kjq62Y1b"}}`, CapMonsterKey, query)))

	for i := 0; i < 10; i++ {
		fasthttp.Do(request, response)
		if taskId := fastjson.GetInt(response.Body(), "taskId"); (taskId != 0) {
			return strconv.Itoa(taskId)
		}
		log.Printf("[%v] Failed to create captcha task, retrying \n", i)
	}
	log.Printf("Failed to create captcha task after 10 attempts \n")
	return ""
}

func getCaptchaResult(taskID string) string {
	var request *fasthttp.Request = createRequest("POST")
	var response *fasthttp.Response = fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetRequestURI("https://api.capmonster.cloud/getTaskResult")
	
	request.Header.Set("Content-Type", "application/json")
	request.SetBody([]byte(fmt.Sprintf(`{"clientKey":"%s","taskId":"%s"}`, CapMonsterKey, taskID)))
	
	for i := 0 ; i < 30; i++ {
		fasthttp.Do(request, response)
		if token := fastjson.GetString(response.Body(), "solution", "gRecaptchaResponse"); (token != "") {
			return token
		}
		time.Sleep(time.Second * 2)
	}
	log.Printf("Failed to get captcha token after 15 attempts \n")
	return ""
}

func submitCaptcha(client *fasthttp.Client, request *fasthttp.Request, captchaToken, query string) bool {
	var response *fasthttp.Response = fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	request.Header.SetMethod("POST")
	request.SetRequestURI("https://www.google.com/sorry/index")

	request.SetBody([]byte(fmt.Sprintf(`g-recaptcha-response=%s&continue=%s`, captchaToken, query)))

	client.Do(request, response)

	return response.StatusCode() == 302 && strings.Contains(response.Header.String(), query)
}