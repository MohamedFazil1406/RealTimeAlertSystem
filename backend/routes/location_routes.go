package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
)

func RegisterLocationRoutes(router *gin.Engine) {

	handler := handlers.NewLocationHandler()

	router.POST("/vehicles/location", handler.UpdateLocation)

	router.GET("/vehicles/location/:vehicle_id", handler.GetCurrentLocation)
}