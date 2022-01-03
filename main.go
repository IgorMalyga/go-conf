package main

import (
	"go-conf/config"
	"go-conf/router"
)

func main() {
	db.Migrate()
	router := router.UserRouter()
	router.Run(":8080")
}
