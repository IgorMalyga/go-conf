package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "go-conf/config"
	"go-conf/models"
	"go-conf/validators"
	"net/http"
	"reflect"
)

func CreateWorkspace(c *gin.Context) {
	db_con := db.GetDB()
	var json validators.WorkspaceValidator

	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := c.MustGet("User").(models.User)
	fmt.Println("var1 = ", reflect.TypeOf(user))
	//workspace := models.Workspace{Name: c.PostForm("Name"), CreatedBy: user.}
	workspace := models.Workspace{Name: c.PostForm("Name"), CreatedBy: user}
	result := db_con.Create(&workspace)

	if result.Error != nil {
		c.JSON(422, gin.H{"error": result.Error})
		return
	}

	c.JSON(201, workspace)
	return
}
