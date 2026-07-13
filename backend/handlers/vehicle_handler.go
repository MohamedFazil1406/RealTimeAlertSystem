package handlers

import (
	"net/http"
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
	"github.com/gin-gonic/gin"
)

type VehicleHandler struct {
	service *services.VehicleService
}

func NewVehicleHandler() *VehicleHandler {
	return &VehicleHandler{
		service: services.NewVehicleService(),
	}
}

func (h *VehicleHandler) CreateVehicle(c *gin.Context) {

	start := time.Now()

	var req dto.CreateVehicleRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	vehicle, err := h.service.CreateVehicle(req)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": vehicle.ID,
		"vehicle_number": vehicle.VehicleNumber,
		"status": vehicle.Status,
		"time_ns": utils.TimeNS(start),
	})
}

func (h *VehicleHandler) GetVehicles(c *gin.Context) {

	start := time.Now()

	vehicles, err := h.service.GetVehicles()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vehicles": vehicles,
		"time_ns": utils.TimeNS(start),
	})
}