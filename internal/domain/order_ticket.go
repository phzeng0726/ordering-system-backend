package domain

import (
	"time"
)

type OrderTicket struct {
	Id               int               `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	SeatId           *int              `gorm:"column:seat_id;" json:"seatId,omitempty"` // 如果Seat被Store刪掉的話，就會變空的(為了段開連接)
	SeatTitle        string            `gorm:"-" json:"seatTitle,omitempty"`
	StoreName        string            `gorm:"-" json:"storeName"` // GetAllByUserId時抓的資料欄位
	UserId           string            `gorm:"column:user_id;" json:"userId,omitempty"`
	TotalPrice       float64           `gorm:"column:total_price;" json:"totalPrice"`
	OrderStatus      string            `gorm:"column:order_status;" json:"orderStatus"`
	CreatedAt        time.Time         `gorm:"column:created_at;" json:"createdAt"`
	OrderTicketItems []OrderTicketItem `gorm:"foreignKey:OrderId;references:id" json:"orderItems"`
	Seat             *Seat             `gorm:"foreignKey:SeatId;references:id" json:"-"`
}

type OrderTicketItem struct {
	Id           int     `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	OrderId      int     `gorm:"column:order_id;" json:"-"`
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
