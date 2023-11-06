package v1

import (
	"net/http"
	"ordering-system-backend/internal/service"

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

type createOrderTicketInput struct {
	SeatId     int                          `json:"seatId" binding:"required"`
	UserId     string                       `json:"userId" binding:"required"`
	TotalPrice *float64                     `json:"totalPrice" binding:"required"`
	OrderItems []createOrderTicketItemInput `json:"orderItems" binding:"required,dive,required"`
}

type createOrderTicketItemInput struct {
	ProductId    int      `json:"productId" binding:"required"`
	ProductName  string   `json:"productName" binding:"required"`
	ProductPrice *float64 `json:"productPrice" binding:"required"`
	Quantity     *int     `json:"quantity" binding:"required"`
}

func (h *Handler) createTicket(c *gin.Context) {
	var inp createOrderTicketInput

	if err := c.BindJSON(&inp); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var orderItems []service.OrderTicketItemInput
	for _, oi := range inp.OrderItems {
		orderItems = append(orderItems, service.OrderTicketItemInput{
			ProductId:    oi.ProductId,
			ProductName:  oi.ProductName,
			ProductPrice: *oi.ProductPrice,
			Quantity:     *oi.Quantity,
		})
	}

	err := h.services.OrderTickets.Create(c.Request.Context(), service.CreateOrderTicketInput{
		SeatId:     inp.SeatId,
		UserId:     inp.UserId,
		TotalPrice: *inp.TotalPrice,
		OrderItems: orderItems,
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
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
