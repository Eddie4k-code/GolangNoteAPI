package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Shape of data that will be returned back to client once a JWT is generated.
type UserData struct {
	ID       int
	Username string
}

// Generates a JSON Web Token to authenticate a logged in user.
func GenerateJWT(userData UserData) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        time.Now().Add(1 * time.Minute).Unix(),
		"authorized": true,
		"username":   userData.Username,
		"ID":         userData.ID,
	})

	var sampleSecretKey = []byte("SecretYouShouldHide")

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
