package v1

import (
	"net/http"
	"ordering-system-backend/internal/domain"
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

func (h *Handler) createSeat(c *gin.Context) {
	var inp createSeatInput
	storeId := c.Param("store_id")

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	seat := domain.Seat{
		StoreId: storeId,
		Title:   inp.Title,
	}

	if err := h.services.Seats.Create(c.Request.Context(), seat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

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

	seat := domain.Seat{
		Id:      seatId,
		StoreId: storeId,
		Title:   inp.Title,
	}

	if err := h.services.Seats.Update(c.Request.Context(), seat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, true)
}

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
