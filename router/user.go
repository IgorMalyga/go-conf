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

	return router
}
