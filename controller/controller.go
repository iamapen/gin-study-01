package controller

/**
 * Controller は、引数にgin.Contextを取る関数である。
 */
import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
