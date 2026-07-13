package handlers

import (
	"net/http"
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	service *services.AlertService
}

func NewAlertHandler() *AlertHandler {
	return &AlertHandler{
		service: services.NewAlertService(),
	}
}

func (h *AlertHandler) Create(c *gin.Context) {

	start := time.Now()

	var req dto.CreateAlertRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	alert, err := h.service.Create(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"alert_id": alert.ID,
		"geofence_id": alert.GeofenceID,
		"vehicle_id": alert.VehicleID,
		"event_type": alert.EventType,
		"status": alert.Status,
		"time_ns": utils.TimeNS(start),
	})
}

func (h *AlertHandler) GetAll(c *gin.Context) {

	start := time.Now()

	alerts, err := h.service.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"alerts": alerts,
		"time_ns": utils.TimeNS(start),
	})
}