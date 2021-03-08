package auth

import "regexp"

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
