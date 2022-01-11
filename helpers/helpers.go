package helpers

import (
	"strconv"
	"time"
	"wishlists/database"
	"wishlists/models"

	"github.com/dgrijalva/jwt-go/v4"
)

const secretKey string = "w1shl1st_4pp_s3cr3t"

func CreateUserToken(userID uint, expiration time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(userID)),
		ExpiresAt: jwt.At(expiration),
	})

	secret := []byte(secretKey)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserFromToken(tokenString string) (models.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return models.User{}, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	return user, nil
}