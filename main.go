package main

import (
	db "note_api/db"
	users "note_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	/* Initalize and Connect to database */
	db.InitDB()

	/* All User Routes */
	server.GET("/users", users.GetUsers)
	server.POST("/users", users.PostUsers)
	server.POST("/login", users.LoginUsers)

	/* All Note Routes */
	server.Run()

}
