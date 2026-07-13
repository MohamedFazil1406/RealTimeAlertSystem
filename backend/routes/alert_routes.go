package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
)

func RegisterAlertRoutes(router *gin.Engine) {

	handler := handlers.NewAlertHandler()

	router.POST("/alerts/configure", handler.Create)

	router.GET("/alerts", handler.GetAll)
}