package main

import (
	"local/crypto-api/auth"
	"net/http"

	"local/crypto-api/crypto"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", auth.LoginEndpoint)

	cryptoGroup := r.Group("/api/crypto")
	cryptoGroup.Use(AuthMidleware())
	{
		cryptoGroup.GET("/btc", crypto.GetCryptoEndpoint)
		cryptoGroup.POST("/btc", crypto.UpdateCryptoEndpoint)
	}

	r.Run()
}

// AuthMidleware checks if the used token is valid
func AuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token inválido",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
