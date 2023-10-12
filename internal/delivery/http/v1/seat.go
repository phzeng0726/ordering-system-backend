package v1

import (
	"net/http"

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
	c.IndentedJSON(http.StatusOK, nil)
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
	c.IndentedJSON(http.StatusOK, nil)
}
