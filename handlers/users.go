package users

import (
	"net/http"
	user "note_api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* Retrieve all Users */
func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, user.Users)
}

/* Register a User */
func PostUsers(ctx *gin.Context) {

	var newUser user.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that both username and password are present
	if newUser.Username == "" || newUser.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username and Password are required"})
		return
	}

	//Hash the plain text password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newUser = user.User{Username: newUser.Username, Password: string(hashedPassword)}

	newUser.CreateUser()

	ctx.JSON(http.StatusCreated, newUser)

}

/* Login a User */
func LoginUsers(ctx *gin.Context) {

	var loginUser user.User

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFound := user.FindUser(loginUser)

	if userFound == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Does not Exist!"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginUser.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	ctx.JSON(http.StatusOK, "OK")

}
