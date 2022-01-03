package controllers

import (
	"github.com/gin-gonic/gin"
	"go-conf/config"
	user "go-conf/models"
	"go-conf/validators"
	"net/http"
)

func SignUp(c *gin.Context) {
	var json validators.SignUpValidator

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := user.User{Email: c.PostForm("email"), Password: c.PostForm("password1")}
	db_con := db.GetDB()
	result := db_con.Create(&user)

	if result.Error != nil {
		c.JSON(422, gin.H{"error": result.Error})
		return
	}

	c.JSON(201, user)
	return
}
