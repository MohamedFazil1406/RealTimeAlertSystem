package main

import (
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/config"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadEnv()

	database.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
	routes.RegisterDashboardRoutes(router)

	router.Run(":" + config.GetEnv("PORT"))
}