package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func main() {
	var router *gin.Engine = gin.Default()
	

	router.GET("/search", func(c *gin.Context) {
		var query string = c.Query("q")
		var page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
		if (err != nil || query == "") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query or page number"})
			return
		}

		var results []SearchResult = Query(query, page)
		if (results == nil) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": results})
	})

	router.Run(":8080") // Start server on port 8080
}
