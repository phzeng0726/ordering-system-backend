package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUserSeatsRoutes(api *gin.RouterGroup) {
	stores := api.Group("/stores/:store_id/seats")
	{
		stores.POST("", h.createSeat)
		stores.PATCH("/:seat_id", h.updateSeat)
		stores.DELETE("/:seat_id", h.deleteSeat)
		stores.GET("", h.getAllSeatsByStoreId)
		stores.GET("/:seat_id", h.getSeatBySeatId)
	}
}

func (h *Handler) createSeat(c *gin.Context) {
	data, err := h.services.Seats.LoadImage()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

func (h *Handler) updateSeat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) deleteSeat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getAllSeatsByStoreId(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *Handler) getSeatBySeatId(c *gin.Context) {
	storeId := c.Param("store_id")
	seatIdStr := c.Param("seat_id")
	seatId, err := strconv.Atoi(seatIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	seat, err := h.services.Seats.GetById(c.Request.Context(), storeId, seatId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, seat)
}
