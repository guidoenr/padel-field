package api

import (
	"github.com/gin-gonic/gin"
	"github.com/guidoenr/padel-field/controllers"
	"net/http"
)

// ListenAndServe turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func ListenAndServe() {
	router := gin.Default()

	//router.LoadHTMLFiles("../padel-ui/public/index.html")

	router.GET("/", showIndex())

	turnos := router.Group("/turnos")
	{
		turnos.GET("/", showTurnos())
		turnos.GET("/:id", showTurnoByID())
		turnos.POST("/:id/reserve", reserveTurno())
		turnos.POST("/:id/cancel", cancelTurno())
	}

	users := router.Group("/users")
	{
		users.GET("/:id/turnos", showTurnosByOwnerId())
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}

// showTurnos is the main page for the turnos website
func showIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
		})
	}
}

// showTurnos is the main page for the turnos website
func showTurnos() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := controllers.GetAvailableTurnos()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turnos})
	}
}

// showTurnoByID is the main page for the turnos website
func showTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		turno, err := controllers.GetTurnoById(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"turno": turno})
	}
}

// showTurnoByID is the main page for the turnos website
func showTurnosByOwnerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := controllers.GetTurnosByOwnerId(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"turno": turnos})
	}
}

// reserveTurno is the main page for the turnos website
func reserveTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, check the ownerID logic?
		err := controllers.ReserveTurno(c.Param("id"), 0)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"turno": err})
	}
}

// cancelTurno is the main page for the turnos website
func cancelTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, check the ownerID logic?
		err := controllers.CancelTurno(c.Param("id"), 0)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"turno": err})
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
