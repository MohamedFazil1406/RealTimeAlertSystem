package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
)

func RegisterViolationRoutes(router *gin.Engine) {

	handler := handlers.NewViolationHandler()

	router.GET("/violations/history", handler.GetHistory)
}