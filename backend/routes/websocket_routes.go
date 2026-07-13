package routes

import (
	ws "github.com/MohamedFazil1406/RealTimeAlertSystem/websocket"

	"github.com/gin-gonic/gin"
)

func RegisterWebSocket(router *gin.Engine) {

	router.GET("/ws/alerts", ws.HandleWebSocket)
}