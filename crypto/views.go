package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCryptoView returns...
func GetCryptoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}

// UpdateCryptoView does...
func UpdateCryptoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "retorno...",
	})
}
