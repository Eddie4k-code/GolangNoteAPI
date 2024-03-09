package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Shape of data that will be returned back to client once a JWT is generated.
type UserData struct {
	ID       primitive.ObjectID
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

/* Used to verify a json web token */
func VerifyJwt(tokenString string) (*jwt.Token, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte("SecretYouShouldHide"), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
