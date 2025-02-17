package main

func main() {
	client := createClient()
	request := createRequest()

	storeInitialCookies(client, request)
}