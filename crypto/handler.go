package crypto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCryptoEndpoint the exchange rate
func GetCryptoEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, getCryptoBtc())
}

// UpdateCryptoEndpoint updates the currency rate
func UpdateCryptoEndpoint(c *gin.Context) {
	var body updateCryptoBody
	c.Bind(&body)
	updateCurrency(body.Currency, int(body.Value*10000))

	c.JSON(http.StatusOK, gin.H{
		"message": "Valor alterado com sucesso",
	})
}
