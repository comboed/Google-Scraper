package main

import "fmt"

func main() {
	task := createCaptchaTask("https://www.google.com/sorry/index?continue=https://www.google.com/search%3Fq%3Dtest%26oq%3Dtest%26gs_lcrp%3DEgZjaHJvbWUqBggAEEUYOzIGCAAQRRg7MgYIARBFGDsyBwgCEAAYjwIyBwgDEAAYjwIyBwgEEAAYjwIyBggFEEUYPDIGCAYQRRg8MgYIBxBFGDzSAQc3MzFqMGoxqAIAsAIA%26sourceid%3Dchrome%26ie%3DUTF-8&q=EgS8fl46GOGpy70GIjDdtZLTQRKy79ARWlVNp9maOeW9o9782Ep9dX8TIQr61oUo9X5UqWkQv6YDMGpe1rYyAXJKGVNPUlJZX0FCVVNJVkVfTkVUX01FU1NBR0VaAUM")
	token := getCaptchaResult(task)
	fmt.Println(token)
}