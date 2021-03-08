package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCryptoEndpoint returns...
func GetCryptoEndpoint(c *gin.Context) {
	getCryptoBtc()
	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}

// UpdateCryptoEndpoint does...
func UpdateCryptoEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}
