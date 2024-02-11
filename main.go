package main

import (
	users "note_api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	/* All User Routes */
	server.GET("/users", users.GetUsers)
	server.POST("/users", users.PostUsers)
	server.POST("/login", users.LoginUsers)

	/* All Note Routes */

	server.Run()

}
