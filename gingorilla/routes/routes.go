package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muhsufyan/websocket/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	go controllers.HandleMessages()
	r.LoadHTMLGlob("templates/*/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.Index)
	r.GET("/ws", controllers.HandleConnections)

	return r
}
