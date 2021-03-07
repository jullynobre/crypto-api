package main

import (
	"local/crypto-api/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", auth.LoginView)
	r.Run()
}
