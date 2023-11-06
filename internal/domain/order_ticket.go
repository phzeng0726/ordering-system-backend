package domain

import (
	"errors"
	"time"
)

type OrderTicket struct {
	Id               int               `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	SeatId           int               `gorm:"column:seat_id;" json:"seatId"`
	UserId           string            `gorm:"column:user_id;" json:"userId,omitempty"`
	TotalPrice       float64           `gorm:"column:total_price;" json:"totalPrice"`
	OrderStatus      string            `gorm:"column:order_status;" json:"orderStatus"`
	CreatedAt        time.Time         `gorm:"column:created_at;" json:"createdAt"`
	OrderTicketItems []OrderTicketItem `gorm:"foreignKey:OrderId;references:id" json:"orderTicketItems"`
}

type OrderTicketItem struct {
	Id           int     `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	OrderId      int     `gorm:"column:order_id;" json:"orderId"`
	ProductId    int     `gorm:"column:product_id;" json:"productId"`
	ProductName  string  `gorm:"column:product_name;" json:"productName"`
	ProductPrice float64 `gorm:"column:product_price;" json:"productPrice"`
	Quantity     int     `gorm:"column:quantity;" json:"quantity"`
}

// golang沒有enum
type OrderStatus int

const (
	Open OrderStatus = iota
	InProgress
	Cancelled
	Done
)

func OrderStatusConverter(orderStatus OrderStatus) (string, error) {
	var orderStatusStr string

	if orderStatus == Open {
		orderStatusStr = "open"
	} else if orderStatus == InProgress {
		orderStatusStr = "inProgress"
	} else if orderStatus == Cancelled {
		orderStatusStr = "cancelled"
	} else if orderStatus == Done {
		orderStatusStr = "done"
	} else {
		return orderStatusStr, errors.New("order status not available")
	}

	return orderStatusStr, nil
}
