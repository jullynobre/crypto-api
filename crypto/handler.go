package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCryptoEndpoint returns...
func GetCryptoEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, getCryptoBtc())
}

// UpdateCryptoEndpoint does...
func UpdateCryptoEndpoint(c *gin.Context) {
	updateCurrency("BRL", 60000)

	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}
