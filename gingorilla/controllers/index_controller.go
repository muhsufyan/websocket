package controllers

import (
	"net/http"

	"github.com/muhsufyan/websocket/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients

var broadcast = make(chan models.Message) // broadcast channel

// Configure the upgrader. convert http ke websocket
var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}

func Index(c *gin.Context) {

	c.HTML(200, "index.html", nil)
}
