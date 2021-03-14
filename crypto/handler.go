package crypto

import (
	"local/crypto-api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCryptoEndpoint returns...
func GetCryptoEndpoint(c *gin.Context) {
	err := auth.ValidateToken(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token inválido",
		})
		return
	}
	c.JSON(http.StatusOK, getCryptoBtc())
}

// UpdateCryptoEndpoint does...
func UpdateCryptoEndpoint(c *gin.Context) {
	var body updateCryptoBody
	c.Bind(&body)
	updateCurrency(body.Currency, int(body.Value*10000))

	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}
