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
	router.LoadHTMLGlob("static/html/*")
	router.Static("/static", "./static")

	router.GET("/", index())
	router.GET("/login", login())
	router.GET("/register", register())

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
		turnos, err := controllers.GetAvailableTurnos()
		c.HTML(http.StatusOK, "turnos.html", gin.H{
			"title":           "Turnos",
			"availableTurnos": turnos,
		})
		if err != nil {
			c.String(500, "ERROR GETTING TURNOS")
		}
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

// ---------------------- LOGIN - REGISTER
// login
func login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	}
}

// register
func register() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "register",
		})
	}
}
