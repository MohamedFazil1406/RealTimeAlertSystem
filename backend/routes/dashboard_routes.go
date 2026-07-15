package routes

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(router *gin.Engine) {

	dashboardHandler := handlers.NewDashboardHandler()

router.GET("/dashboard", dashboardHandler.GetStats)
}