package handlers

import (
	"net/http"
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	service *services.LocationService
}

func NewLocationHandler() *LocationHandler {
	return &LocationHandler{
		service: services.NewLocationService(),
	}
}

func (h *LocationHandler) UpdateLocation(c *gin.Context) {

	start := time.Now()

	var req dto.UpdateLocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	err := h.service.UpdateLocation(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicle_id": req.VehicleID,
		"location_updated": true,
		"time_ns": utils.TimeNS(start),
	})
}

func (h *LocationHandler) GetCurrentLocation(c *gin.Context) {

	start := time.Now()

	vehicleID := c.Param("vehicle_id")

	data, err := h.service.GetCurrentLocation(vehicleID)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	data["time_ns"] = utils.TimeNS(start)

	c.JSON(http.StatusOK, data)
}