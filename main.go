package main

import (
	"local/crypto-api/auth"

	"local/crypto-api/crypto"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", auth.LoginView)
	r.GET("/api/crypto/btc", crypto.GetCryptoView)
	r.POST("/api/crypto/btc", crypto.UpdateCryptoView)

	r.Run()
}
