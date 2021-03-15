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

	r.Use(CORSMiddleware())

	r.Run(":8000")
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

// CORSMiddleware allows CORS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
