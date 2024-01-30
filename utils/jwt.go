package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// in a real app this would be an env variable
const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  "",
		"userId": "",
		// expires in 2 hours
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// this syntax checks if token.Method is of type jwt.SigningMethodHMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return errors.New("error parsing jwt token")
	}

	isValid := parsedToken.Valid

	if !isValid {
		return errors.New("invalid token")
	}

	// claims is the data on the token
	// which we set to userId and email on the token when we created it
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
