package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func main() {
	proxies = openFile("./data/proxies.txt")
	var router *gin.Engine = gin.Default()

	router.GET("/search", func(context *gin.Context) {
		var query string = context.Query("q")
		var page, err = strconv.Atoi(context.DefaultQuery("page", "0"))
		
		if (err != nil || query == "" || page < 0) {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query or page number"})
			page = 0 // Set it to default page number
		}

		var results []SearchResult = Query(query, page - 1)
		if (results == nil) {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": results})
	})
	router.Run(":8080") // Start server on port 8080
}