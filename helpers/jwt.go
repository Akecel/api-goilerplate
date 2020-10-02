package helpers

import (
	"time"

	e "gechoplate/models"

	"github.com/dgrijalva/jwt-go"
)

// GetJWTToken create a new JWT token for users.
func GetJWTToken(res e.User) (t string, err error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = res.ID
	claims["name"] = res.LastName + res.FirstName
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte("secret"))
	return
}
