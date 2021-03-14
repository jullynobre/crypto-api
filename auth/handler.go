package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginEndpoint recieves a resquest context and returns an access token if the request is valid
func LoginEndpoint(c *gin.Context) {
	var user userModel

	// Verify if the fields structure are valid
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Campos Inválidos",
		})
		return
	}
	isUserValid := user.validate()

	// Verify if the user is valid
	if !isUserValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email e/ou senha inválido(s)",
		})
		return
	}

	// Create an access token and returns it
	token, err := CreateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao criar token",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}
