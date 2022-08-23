package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListenAndServe turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func ListenAndServe() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", index())

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}

// index is the main page for the website
func index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Turnos",
		})
	}
}
