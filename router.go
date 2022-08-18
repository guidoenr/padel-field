package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initialize() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "turnos.html", gin.H{
			"title": "Turnos",
		})
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
