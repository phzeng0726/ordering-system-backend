package utils

import (
	"errors"
	"ordering-system-backend/internal/domain"
	"time"
)

func DateTimeConverter(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeStr)
}
func TimeConverter(timeStr string) (time.Time, error) {
	return time.Parse("15:04:05", timeStr)
}

func OrderStatusConverter(orderStatus domain.OrderStatus) (string, error) {
	var orderStatusStr string

	if orderStatus == domain.Open {
		orderStatusStr = "open"
	} else if orderStatus == domain.InProgress {
		orderStatusStr = "inProgress"
	} else if orderStatus == domain.Cancelled {
		orderStatusStr = "cancelled"
	} else if orderStatus == domain.Done {
		orderStatusStr = "done"
	} else {
		return orderStatusStr, errors.New("order status not available")
	}

	return orderStatusStr, nil
}
