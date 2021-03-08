package main

import (
	"local/crypto-api/auth"

	"local/crypto-api/crypto"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/api/login", auth.LoginEndpoint)
	r.GET("/api/crypto/btc", crypto.GetCryptoEndpoint)
	r.POST("/api/crypto/btc", crypto.UpdateCryptoEndpoint)

	r.Run()
}
