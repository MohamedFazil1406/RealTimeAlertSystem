package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
)

type GeofenceHandler struct {
	service *services.GeofenceService
}

func NewGeofenceHandler() *GeofenceHandler {
	return &GeofenceHandler{
		service: services.NewGeofenceService(),
	}
}

func (h *GeofenceHandler) Create(c *gin.Context) {

	start := time.Now()

	var req dto.CreateGeofenceRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	if !utils.ValidCategory(req.Category) {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid category",
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	geo, err := h.service.Create(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": geo.ID,
		"name": geo.Name,
		"status": geo.Status,
		"time_ns": utils.TimeNS(start),
	})
}

func (h *GeofenceHandler) GetAll(c *gin.Context) {

	start := time.Now()

	data, err := h.service.GetAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"geofences": data,
		"time_ns": utils.TimeNS(start),
	})
}