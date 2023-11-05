package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initOrderTicketRoutes(api *gin.RouterGroup) {
	tickets := api.Group("/order-tickets")
	{
		tickets.POST("", h.createTicket)
	}

	ticketWithStores := api.Group("/stores/:store_id/order-tickets")
	{
		ticketWithStores.PATCH("/:ticket_id", h.updateTicket)
		ticketWithStores.DELETE("/:ticket_id", h.deleteTicket)
		ticketWithStores.GET("", h.getAllTicketsByStoreId)
	}
}

type createTicketInput struct {
	Title       string          `json:"title" binding:"required"`
	Description string          `json:"description"`
	MenuItems   []menuItemInput `json:"menuItems" binding:"required"`
}

func (h *Handler) createTicket(c *gin.Context) {
	var inp createTicketInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fmt.Println(inp)

	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) updateTicket(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) deleteTicket(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}

func (h *Handler) getAllTicketsByStoreId(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, true)
}
