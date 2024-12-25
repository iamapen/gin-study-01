package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// route定義
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 起動
	router.Run()
}
