package main

import (
	// "os"

	"github.com/gin-gonic/gin"
)

func main() {
	serve := gin.Default()
	serve.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"title": "Efficax Counter",
			"status": "ok",
			"code": 200,
		})
	})
}