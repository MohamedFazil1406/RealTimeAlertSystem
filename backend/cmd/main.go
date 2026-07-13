package main

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/config"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	database.ConnectDatabase()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Geofence API Running",
		})

	})

	routes.RegisterVehicleRoutes(router)
	routes.RegisterGeofenceRoutes(router)
	routes.RegisterLocationRoutes(router)
	routes.RegisterAlertRoutes(router)
	routes.RegisterViolationRoutes(router)
	routes.RegisterWebSocket(router)

	router.Run(":" + config.GetEnv("PORT"))
}