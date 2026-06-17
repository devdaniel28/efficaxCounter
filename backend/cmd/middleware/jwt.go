package middleware

import (
	"efficaxcounter/cmd/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtEnv = os.Getenv("SIGNATUREJWT")
var jwtScret = []byte(jwtEnv)

func GerenerateJwt(user_id int, email string,) (string, error) {
	userModel := models.User{}
	claims := jwt.MapClaims{
		userModel.Email: user_id,
		userModel.Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtScret)
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return  "", err
	}

	return string(hashed), nil
}