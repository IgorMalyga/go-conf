package router

import (
	"github.com/gin-gonic/gin"
	"go-conf/controllers"
)

func UserRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)

	return router
}
