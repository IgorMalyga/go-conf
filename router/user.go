package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-conf/controllers"
	"go-conf/middleware"
	user "go-conf/models"
)

func UserRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.SignUp)
	router.POST("/token", middleware.AuthMiddleware().LoginHandler)
	router.Use(middleware.AuthMiddleware().MiddlewareFunc())
	router.Use(middleware.RequestUser())
	router.GET("/ping", func(c *gin.Context) {

		us := c.MustGet("User").(user.User)
		fmt.Printf("User: %v", us)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
