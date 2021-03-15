package auth

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
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
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Create token
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Sign token
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	// Remove the prefix Bearer
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ValidateToken returns nil if the token is valid or else returns an error
func ValidateToken(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
