package main

import (
	"github.com/SpoonBuoy/waba/controllers"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-gonic/gin"
)

var (
	wabaService    = service.NewWabaService()
	chatController = controllers.NewChatController(wabaService)
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hook/verify", chatController.Verify)
	r.POST("hook/verify", chatController.Verify)
	r.Run(":9000")
}
