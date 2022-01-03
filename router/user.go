package router

import (
	"github.com/gin-gonic/gin"
	"go-conf/controllers"
	"go-conf/middleware"
)

func UserRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.SignUp)
	router.POST("/token", middleware.AuthMiddleware().LoginHandler)
	router.Use(middleware.AuthMiddleware().MiddlewareFunc())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}
