package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
)

func RegisterVehicleRoutes(router *gin.Engine) {

	handler := handlers.NewVehicleHandler()

	router.POST("/vehicles", handler.CreateVehicle)

	router.GET("/vehicles", handler.GetVehicles)
}