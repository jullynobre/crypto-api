package auth

import (
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type userModel struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Regex for containing one "@", one period and at least one character before and after them
var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s\.]+\.[^@\.\s]+$`)

// Regex for containing 6 characters with only numbers
var passwordRegex = regexp.MustCompile(`^[0-9]{6}$`)

func (user userModel) validate() bool {
	if emailRegex.MatchString(user.Email) && passwordRegex.MatchString(user.Password) {
		return true
	}
	return false
}

// CreateToken receives and user email and returns a valid token or and error
func CreateToken(userEmail string) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	// Create token
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Sign token
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
