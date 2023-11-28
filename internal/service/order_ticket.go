package service

import (
	"context"
	"errors"
	"fmt"
	"ordering-system-backend/internal/domain"
	"ordering-system-backend/internal/repository"
	"ordering-system-backend/internal/utils"
	firebase_fcm "ordering-system-backend/pkg/firebase_core/notification"
	"strings"

	"firebase.google.com/go/messaging"
)

type OrderTicketsService struct {
	orderRepo repository.OrderTickets
	fcmRepo   repository.FCMTokens
	seatRepo  repository.Seats
}

func NewOrderTicketsService(orderRepo repository.OrderTickets, fcmRepo repository.FCMTokens, seatRepo repository.Seats) *OrderTicketsService {
	return &OrderTicketsService{
		orderRepo: orderRepo,
		fcmRepo:   fcmRepo,
		seatRepo:  seatRepo,
	}
}

func (s *OrderTicketsService) pushFirebaseNotification(userType int, deviceTokens []string, notification *messaging.Notification) error {
	// Push FCM Token
	fcmClient, err := firebase_fcm.Init(userType)
	if err != nil {
		return err
	}

	if err := firebase_fcm.SendNotification(fcmClient, deviceTokens, notification); err != nil {
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
		SeatId:           &input.SeatId,
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

	seat, err := s.seatRepo.GetSeatBySeatId(ctx, input.SeatId)
	if err != nil {
		return err
	}

	// 以 FCM 通知商家刷新頁面，因為一個store user有好幾個store，所以要特別抓出是哪個store收到更新，且如果Device Tokens為空的時候不報錯
	if err := s.pushFirebaseNotification(0, deviceTokens, &messaging.Notification{
		Title: "NEW_ORDER_TICKET",
		Body:  seat.StoreId,
	}); err != nil {
		if strings.Contains(err.Error(), "tokens must not be nil or empty") {
			fmt.Println("empty device tokens")
			return nil
		}
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

	// 以 ticketId 撈出userId，userId到fcm_token撈出device tokens
	deviceTokens, err := s.fcmRepo.GetAllByTicketId(ctx, ticketId)
	if err != nil {
		return err
	}

	// 以 FCM 通知客戶刷新頁面，且如果Device Tokens為空的時候不報錯
	if err := s.pushFirebaseNotification(1, deviceTokens, &messaging.Notification{
		Title: "UPDATE_ORDER_TICKET",
		Body:  fmt.Sprintf("%d", ticketId),
	}); err != nil {
		if strings.Contains(err.Error(), "tokens must not be nil or empty") {
			fmt.Println("empty device tokens")
			return nil
		}
		return err
	}

	return nil
}

func (s *OrderTicketsService) Delete(ctx context.Context, storeId string, ticketId int) error {
	// 以 ticketId 撈出userId，userId到fcm_token撈出device tokens
	deviceTokens, err := s.fcmRepo.GetAllByTicketId(ctx, ticketId)
	if err != nil {
		return err
	}

	// 先撈再刪
	if err := s.orderRepo.Delete(ctx, storeId, ticketId); err != nil {
		return err
	}

	// 以 FCM 通知客戶刷新頁面，且如果Device Tokens為空的時候不報錯
	if err := s.pushFirebaseNotification(1, deviceTokens, &messaging.Notification{
		Title: "DELETE_ORDER_TICKET",
		Body:  fmt.Sprintf("%d", ticketId),
	}); err != nil {
		if strings.Contains(err.Error(), "tokens must not be nil or empty") {
			fmt.Println("empty device tokens")
			return nil
		}
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

func (s *OrderTicketsService) GetAllByUserId(ctx context.Context, userId string) ([]domain.OrderTicket, error) {
	orderTickets, err := s.orderRepo.GetAllByUserId(ctx, userId)
	if err != nil {
		return orderTickets, err
	}

	for i, orderTicket := range orderTickets {
		if orderTicket.Seat != nil {
			_seat := *orderTicket.Seat
			orderTickets[i].StoreName = _seat.Store.Name
		}

	}

	return orderTickets, nil
}
