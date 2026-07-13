package routes

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterGeofenceRoutes(router *gin.Engine) {

	handler := handlers.NewGeofenceHandler()

	router.POST("/geofences", handler.Create)

	router.GET("/geofences", handler.GetAll)
}