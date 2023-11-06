package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
)

type OrderTicketsService struct {
	repo repository.OrderTickets
}

func NewOrderTicketsService(repo repository.OrderTickets) *OrderTicketsService {
	return &OrderTicketsService{repo: repo}
}

func (s *OrderTicketsService) Create(ctx context.Context, input CreateOrderTicketInput) error {
	var orderItems []domain.OrderTicketItem
	var totalPrice float64

	for _, oi := range input.OrderItems {
		orderItems = append(
			orderItems,
			domain.OrderTicketItem{
				ProductId:    oi.ProductId,
				ProductName:  oi.ProductName,
				ProductPrice: oi.ProductPrice,
				Quantity:     oi.Quantity,
			},
		)
		totalPrice += oi.ProductPrice * float64(oi.Quantity)
	}

	orderTicket := domain.OrderTicket{
		Id:               0,
		SeatId:           input.SeatId,
		UserId:           input.UserId,
		TotalPrice:       totalPrice,
		OrderStatus:      "open",
		OrderTicketItems: orderItems,
	}

	if err := s.repo.Create(ctx, orderTicket); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error) {
	orderTickets, err := s.repo.GetAllByStoreId(ctx, storeId)
	if err != nil {
		return orderTickets, err
	}

	return orderTickets, nil
}
