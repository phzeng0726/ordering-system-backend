package service

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
	"ordering-system-backend/internal/utils"
	"ordering-system-backend/pkg/notification"
)

type OrderTicketsService struct {
	orderRepo repository.OrderTickets
	fcmRepo   repository.FCMTokens
}

func NewOrderTicketsService(orderRepo repository.OrderTickets, fcmRepo repository.FCMTokens) *OrderTicketsService {
	return &OrderTicketsService{orderRepo: orderRepo, fcmRepo: fcmRepo}
}

func (s *OrderTicketsService) pushFirebaseNotification(deviceTokens []string) error {
	// Push FCM Token
	fcmClient, err := notification.Init()
	if err != nil {
		return err
	}
	err = notification.SendPushNotification(fcmClient, deviceTokens)

	if err != nil {
		return err
	}

	return nil
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

	// 建立 OrderTicket
	if err := s.orderRepo.Create(ctx, orderTicket); err != nil {
		return err
	}

	// 撈取該SeatId的商家，並獲取該商家的所有Device Token
	deviceTokens, err := s.fcmRepo.GetAllBySeatId(ctx, input.SeatId)
	if err != nil {
		return err
	}

	// 以 FCM 通知刷新頁面
	if err := s.pushFirebaseNotification(deviceTokens); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) Update(ctx context.Context, storeId string, ticketId int, input UpdateOrderTicketInput) error {
	orderTicket := domain.OrderTicket{
		Id:          ticketId,
		OrderStatus: input.OrderStatus,
	}

	if err := s.orderRepo.Update(ctx, storeId, orderTicket); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) Delete(ctx context.Context, storeId string, ticketId int) error {
	if err := s.orderRepo.Delete(ctx, storeId, ticketId); err != nil {
		return err
	}

	return nil
}

func (s *OrderTicketsService) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error) {
	orderTickets, err := s.orderRepo.GetAllByStoreId(ctx, storeId)
	if err != nil {
		return orderTickets, err
	}

	return orderTickets, nil
}
