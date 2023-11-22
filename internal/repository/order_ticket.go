package repository

import (
	"context"
	"errors"
	"ordering-system-backend/internal/domain"

	"gorm.io/gorm"
)

type OrderTicketsRepo struct {
	db *gorm.DB
	rt *RepoTools
}

func NewOrderTicketsRepo(db *gorm.DB, rt *RepoTools) *OrderTicketsRepo {
	return &OrderTicketsRepo{
		db: db,
		rt: rt,
	}
}

func (r *OrderTicketsRepo) Create(ctx context.Context, ticket domain.OrderTicket) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 因為有設fKey，會自動幫忙create orderItems
		if err := tx.Create(&ticket).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) Update(ctx context.Context, storeId string, ticket domain.OrderTicket) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.OrderTicket{}).Where("id = ?", ticket.Id).Updates(ticket).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) Delete(ctx context.Context, storeId string, ticketId int) error {
	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", ticketId).First(&domain.OrderTicket{}).Error; err != nil {
			return errors.New("order ticket not found")
		}

		if err := tx.Where("order_id = ?", ticketId).Delete(&domain.OrderTicketItem{}).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", ticketId).Delete(&domain.OrderTicket{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderTicketsRepo) GetAllByStoreId(ctx context.Context, storeId string) ([]domain.OrderTicket, error) {
	var orderTickets []domain.OrderTicket
	var seats []domain.Seat
	seatMap := make(map[int]string)

	db := r.db.WithContext(ctx)

	if err := db.Transaction(func(tx *gorm.DB) error {
		// 獲取商店的所有座位
		if err := tx.Where("store_id = ?", storeId).Find(&seats).Error; err != nil {
			return err
		}
		for _, seat := range seats {
			seatMap[seat.Id] = seat.Title
		}

		// 以座位查詢訂單
		if err := tx.Preload("OrderTicketItems").
			Where("seat_id IN (SELECT id FROM store_seats WHERE store_id = ?)", storeId).
			Find(&orderTickets).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return orderTickets, err
	}

	// 將seatTitle寫進orderTicket的return json structure，方便frontend使用
	for i, orderTicket := range orderTickets {
		orderTickets[i].SeatTitle = seatMap[orderTicket.SeatId]
	}

	return orderTickets, nil
}
