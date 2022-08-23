package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/padel-field/controllers"
	"net/http"
)

func ListenAndServe() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Turnos",
		})
		controllers.GetAvailableTurnos()
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
