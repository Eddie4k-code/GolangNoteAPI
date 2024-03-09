package middleware

import (
	utils "note_api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

/* Middleware to ensure user is authenticated with verified jwt token */
func VerifyUser(ctx *gin.Context) {

	grabHeader := ctx.GetHeader("Authorization")

	grabbedToken := strings.Split(grabHeader, "Bearer ")[1]

	_, err := utils.VerifyJwt(grabbedToken)

	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid JWT Token"})
		ctx.Abort()
		return
	}

	ctx.Next()

}
