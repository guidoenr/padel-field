package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/padel-field/controllers"
	"net/http"
)

// ListenAndServe turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func ListenAndServe() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", index())

	turnos := router.Group("/turnos")
	{
		turnos.GET("/", indexTurnos())
		turnos.GET("/index", indexTurnos())
	}

	users := router.Group("/users")
	{
		users.GET("/", indexUsers())
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}

// indexTurnos is the main page for the turnos website
func index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Index",
		})
	}
}

// indexTurnos is the main page for the turnos website
func indexTurnos() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "turnos.html", gin.H{
			"title": "Turnos",
		})
		controllers.GetAvailableTurnos()
	}
}

// indexUsers is the main page for the users website
func indexUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "users.html", gin.H{
			"title": "Users",
		})
	}
}
