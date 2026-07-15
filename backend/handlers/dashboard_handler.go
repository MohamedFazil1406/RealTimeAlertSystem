package handlers

import (
	"net/http"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{
		service: services.NewDashboardService(),
	}
}

func (h *DashboardHandler) GetStats(c *gin.Context) {

	stats, err := h.service.GetStats()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}