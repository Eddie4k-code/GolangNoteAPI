package main

import (
	db "note_api/db"
	users "note_api/handlers"
	middleware "note_api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	/* Initalize and Connect to database */
	db.InitDB()

	/* All User Routes */
	server.GET("/users", users.GetUsers)
	server.POST("/register", users.PostUsers)
	server.POST("/login", users.LoginUsers)
	server.GET("/user-check", middleware.VerifyUser, users.CheckAuthenticationTest)

	/* All Note Routes */
	server.Run()

}
