package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/services"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
)

type ViolationHandler struct {
	service *services.ViolationService
}

func NewViolationHandler() *ViolationHandler {

	return &ViolationHandler{
		service: services.NewViolationService(),
	}
}

func (h *ViolationHandler) GetHistory(c *gin.Context) {

	start := time.Now()

	vehicleID := c.Query("vehicle_id")

	geofenceID := c.Query("geofence_id")

	startDate := c.Query("start_date")

	endDate := c.Query("end_date")

	limit := 50

	if l := c.Query("limit"); l != "" {

		n, err := strconv.Atoi(l)

		if err == nil {

			limit = n

			if limit > 500 {
				limit = 500
			}
		}
	}

	data, total, err := h.service.GetHistory(

		vehicleID,
		geofenceID,
		startDate,
		endDate,
		limit,
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{

			"error": err.Error(),

			"time_ns": utils.TimeNS(start),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{

		"violations": data,

		"total_count": total,

		"time_ns": utils.TimeNS(start),
	})
}