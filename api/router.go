package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/guidoenr/padel-field/api/controllers"
	"github.com/guidoenr/padel-field/logger"
	"github.com/guidoenr/padel-field/models"
	"net/http"
	"time"
)

const (
	SecretKey = "secret"
)

// ListenAndServe turns on the gin-gonic server and initialize the entire REST-API
// routes and endpoints
func ListenAndServe() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", showIndex())

	turnos := router.Group("/turnos")
	users := router.Group("/users")
	auth := router.Group("/auth")

	{
		// turnos
		turnos.GET("/", showTurnos())
		turnos.GET("/:id", showTurnoByID())
		turnos.GET("/day/:day", showTurnosByDay())
		turnos.POST("/:id/reserve", reserveTurno())
		turnos.POST("/:id/cancel", cancelTurno())

		// users
		users.GET("/:id/turnos", showTurnosByOwnerId())

		// auth
		auth.POST("/register", register())
		auth.POST("/login", login())
		auth.POST("/logout", logout())
		auth.GET("/user", userGet())
	}

	router.Run()
}

// -------------------------- CONTROLLERS

// showIndex is the main page for the turnos website
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
		if err.Err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turnos})
	}
}

// showTurnos is the main page for the turnos website
func showTurnosByDay() gin.HandlerFunc {
	return func(c *gin.Context) {
		day := c.Param("day")
		turnos, err := controllers.GetAvailableTurnosByDay(day)
		if err.Err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turnos})
	}
}

// showTurnoByID is the main page for the turnos website
func showTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		turno, err := controllers.GetTurnoById(c.Param("id"))
		if err.Err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turno})
	}
}

// showTurnosByOwnerId is the main page for the turnos website
func showTurnosByOwnerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := controllers.GetTurnosByOwnerId(c.Param("id"))
		if err.Err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turnos})
	}
}

// reserveTurno is the main page for the turnos website
func reserveTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, check the ownerID logic?
		err := controllers.ReserveTurno(c.Param("id"), 0)
		if err.Err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"ok": "reserved turno"})
	}
}

// cancelTurno is the main page for the turnos website
func cancelTurno() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO, check the ownerID logic?
		err := controllers.CancelTurno(c.Param("id"), 0)
		if err.Err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"ok": "canceled turno"})
	}
}

// -------------------------- login/register

// login
func login() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User
		err := c.BindJSON(&user)

		var cookie *http.Cookie

		// loggin and storing the cookie
		cookie, reqErr := controllers.Login(&user)
		switch reqErr.StatusCode {
		case -1:
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		case 3: // username does not exist
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		case 4: // wrong password
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
		case 5: // wrong email
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
		default:
			c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
			c.IndentedJSON(http.StatusOK, gin.H{"ok": "Logged in"})
			logger.Loginfo.Printf("user '%s' logged in", user.Username)
		}

	}
}

// register
func register() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newUser models.User

		// unmarshal body data into user struct
		err := c.BindJSON(&newUser)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "can not unmarshal request body"})
			return
		}
		reqErr := controllers.Register(&newUser)
		switch reqErr.StatusCode {
		case -1:
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		case 1: // existing username
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "the username already exist"})
		case 2: // existing email
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "the email already exist"})
		default:
			c.IndentedJSON(http.StatusOK, gin.H{"ok": "User registered"})
			logger.Loginfo.Printf("user '%s' registered", newUser.Username)
		}

	}
}

// userGet
func userGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, err.Error())
			return
		}

		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			return
		}

		claims := token.Claims.(*jwt.StandardClaims)
		user, _ := controllers.GetUserById(claims.Issuer)

		c.IndentedJSON(http.StatusOK, gin.H{"userGet": user})
	}
}

func logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := http.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour), // this expires the cookie because set the expire time at the past
			HttpOnly: true,
		}
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.IndentedJSON(http.StatusOK, gin.H{"ok": "logged out"})
	}

}
