package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//LoginView returns...
func LoginView(c *gin.Context) {
	var user userModel

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Campos Inválidos",
		})
		return
	}

	if user.validate() {
		c.JSON(http.StatusOK, gin.H{
			"token": "some-token",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Email e/ou senha inválido(s)",
	})

}
