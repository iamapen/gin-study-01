package main

import (
	"gin-study-01/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// route定義
	router.GET("/ping", controller.Ping)

	// 起動
	router.Run()
}
