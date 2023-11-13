package service

import (
	"context"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
	"ordering-system-backend/internal/utils"

	"github.com/pkg/errors"
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
	orderStatus, err := utils.OrderStatusConverter(domain.Open) // 預設create時為open
	if err != nil {
		return err
	}

	if len(input.OrderItems) == 0 {
		return errors.New("order items cannot be empty")
	}

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
		SeatId:           input.SeatId,
		UserId:           input.UserId,
		TotalPrice:       totalPrice,
		OrderStatus:      orderStatus,
		OrderTicketItems: orderItems,
	}

	if err := s.repo.Create(ctx, orderTicket); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) Update(ctx context.Context, storeId string, ticketId int, input UpdateOrderTicketInput) error {
	orderTicket := domain.OrderTicket{
		Id:          ticketId,
		OrderStatus: input.OrderStatus,
	}

	if err := s.repo.Update(ctx, storeId, orderTicket); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) Delete(ctx context.Context, storeId string, ticketId int) error {
	if err := s.repo.Delete(ctx, storeId, ticketId); err != nil {
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
