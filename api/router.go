package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers2 "github.com/guidoenr/padel-field/api/controllers"
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
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", showIndex())

	turnos := router.Group("/turnos")
	users := router.Group("/users")
	auth := router.Group("/auth")

	// turnos
	{
		turnos.GET("/", showTurnos())
		turnos.GET("/:id", showTurnoByID())
		turnos.GET("/day/:day", showTurnosByDay())
		turnos.POST("/:id/reserve", reserveTurno())
		turnos.POST("/:id/cancel", cancelTurno())
	}

	// users
	{
		users.GET("/:id/turnos", showTurnosByOwnerId())
	}

	// auth
	{
		auth.POST("/register", register())
		auth.POST("/login", login())
		auth.POST("/logout", logout())
		auth.GET("/user", userGet())
	}
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
		turnos, err := controllers2.GetAvailableTurnos()
		if err != nil {
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
		turnos, err := controllers2.GetAvailableTurnosByDay(day)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turnos})
	}
}

// showTurnoByID is the main page for the turnos website
func showTurnoByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		turno, err := controllers2.GetTurnoById(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"data": turno})
	}
}

// showTurnosByOwnerId is the main page for the turnos website
func showTurnosByOwnerId() gin.HandlerFunc {
	return func(c *gin.Context) {
		turnos, err := controllers2.GetTurnosByOwnerId(c.Param("id"))
		if err != nil {
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
		err := controllers2.ReserveTurno(c.Param("id"), 0)
		if err != nil {
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
		err := controllers2.CancelTurno(c.Param("id"), 0)
		if err != nil {
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
		user := models.User{
			Username: c.Request.PostFormValue("username"),
			Password: c.Request.PostFormValue("password"),
		}
		var cookie *http.Cookie

		cookie, err := controllers2.Login(&user)
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"ok": "Logged in"})
	}
}

// register
func register() gin.HandlerFunc {
	return func(c *gin.Context) {

		logger.Loginfo.Println(c.Request.Body)
		for _, l := range c.Request.PostForm {
			logger.Loginfo.Printf("post form: %s", l)
		}

		newUser := models.User{
			Firstname: c.Request.PostFormValue("name"),
			Lastname:  c.Request.PostFormValue("surname"),
			Username:  c.Request.PostFormValue("username"),
			Email:     c.Request.PostFormValue("email"),
			Password:  c.Request.PostFormValue("password"),
			Phone:     c.Request.PostFormValue("phone"),
		}

		logger.Loginfo.Println(newUser.String())
		err := controllers2.Register(&newUser)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"ok": "User registered"})
	}
}

// login
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
		user, _ := controllers2.GetUserById(claims.Issuer)

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
