package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"
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

type createSeatInput struct {
	Title string `json:"title" binding:"required"`
}

type updateSeatInput struct {
	Title string `json:"title" binding:"required"`
}

// @Tags Store Seats
// @Description Create seat
// @Accept json
// @Param data body createSeatInput true "JSON data"
// @Param store_id path string true "Store Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /stores/{store_id}/seats [post]
func (h *Handler) createSeat(c *gin.Context) {
	var inp createSeatInput
	storeId := c.Param("store_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Seats.Create(c.Request.Context(), storeId, service.CreateSeatInput{
		Title: inp.Title,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Store Seats
// @Description Update seat
// @Accept json
// @Param data body updateSeatInput true "JSON data"
// @Param store_id path string true "Store Id"
// @Param seat_id path int true "Seat Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /stores/{store_id}/seats/{seat_id} [patch]
func (h *Handler) updateSeat(c *gin.Context) {
	var inp updateSeatInput
	storeId := c.Param("store_id")
	seatIdStr := c.Param("seat_id")
	seatId, err := strconv.Atoi(seatIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Seats.Update(c.Request.Context(), storeId, seatId, service.UpdateSeatInput{
		Title: inp.Title,
	}); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Store Seats
// @Description Delete seat
// @Param store_id path string true "Store Id"
// @Param seat_id path int true "Seat Id"
// @Produce json
// @Success 200 {boolean} result
// @Router /stores/{store_id}/seats/{seat_id} [delete]
func (h *Handler) deleteSeat(c *gin.Context) {
	storeId := c.Param("store_id")
	seatIdStr := c.Param("seat_id")
	seatId, err := strconv.Atoi(seatIdStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.services.Seats.Delete(c.Request.Context(), storeId, seatId); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

// @Tags Store Seats
// @Description Get all seats by store id
// @Param store_id path string true "Store Id"
// @Param seat_id path int true "Seat Id"
// @Produce json
// @Success 200 {array} domain.Seat
// @Router /stores/{store_id}/seats [get]
func (h *Handler) getAllSeatsByStoreId(c *gin.Context) {
	storeId := c.Param("store_id")

	seats, err := h.services.Seats.GetAllByStoreId(c.Request.Context(), storeId)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, seats)
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
