package handlers

import (
	"net/http"
	models "note_api/models"
	utils "note_api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/* Retrieve all Users */
func GetUsers(ctx *gin.Context) {

	users, err := models.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

/* Register a User */
func PostUsers(ctx *gin.Context) {

	var newUser models.User

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
		return
	}

	newUser = models.User{Username: newUser.Username, Password: string(hashedPassword)}

	newUser.CreateUser()

	//Generate Token
	jwtToken, err := utils.GenerateJWT(utils.UserData{ID: newUser.ID, Username: newUser.Username})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to generate JWT."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jwtToken": jwtToken})

}

/* Login a User */
func LoginUsers(ctx *gin.Context) {

	var loginUser models.User

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userFound := models.FindUser(loginUser)

	if userFound == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Does not Exist!"})
		return
	}

	//Compare hashed password with entered password..
	err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginUser.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or Password"})
		return
	}

	//Generate a JWT
	jwtToken, err := utils.GenerateJWT(utils.UserData{ID: userFound.ID, Username: userFound.Username})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Unable to generate JWT")
		return
	}

	//Send JWT back to user..
	ctx.JSON(http.StatusOK, gin.H{"jwtToken": jwtToken})

}

func CheckAuthenticationTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Authenticated"})
}
