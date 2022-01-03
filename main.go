package main

import (
	"fmt"
	"go-conf/config"
	"go-conf/router"
)

func main() {
	var db_con = db.GetDB()
	db.Migrate()
	router := router.UserRouter()
	fmt.Println(db_con)
	router.Run(":8080")
}
