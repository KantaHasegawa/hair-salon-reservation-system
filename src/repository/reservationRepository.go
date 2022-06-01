package repository

import (
	"context"
	"fmt"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"time"

	"github.com/google/uuid"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) All(ctx context.Context) ([]entity.Reservation, error) {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	Reservations := []entity.Reservation{}
	err := tx.Find(&Reservations).Error
	if err != nil {
		return []entity.Reservation{}, fmt.Errorf("failed to All: %w", err)
	}
	return Reservations, nil
}
func (r *ReservationRepository) Find(ctx context.Context, id string) (entity.Reservation, error) {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	Reservation := entity.Reservation{}
	err := tx.First(&Reservation, "id = ?", id).Error
	if err != nil {
		return entity.Reservation{}, fmt.Errorf("failed to Find: %w", err)
	}
	return Reservation, nil
}

func (r *ReservationRepository) FindByBeauticianAndTime(ctx context.Context, beauticianId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error) {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	Reservations := []entity.Reservation{}
	err := tx.Find(&Reservations, "beautician_id = ? AND ? < end_time and ? > start_time ", beauticianId, startTime, endTime).Error
	if err != nil {
		return []entity.Reservation{}, fmt.Errorf("failed to All: %w", err)
	}
	return Reservations, nil
}

func (r *ReservationRepository) FindByCustomerAndTime(ctx context.Context, customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error) {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	Reservations := []entity.Reservation{}
	err := tx.Find(&Reservations, "customer_id = ? AND ? < end_time and ? > start_time ", customerId, startTime, endTime).Error
	if err != nil {
		return []entity.Reservation{}, fmt.Errorf("failed to All: %w", err)
	}
	return Reservations, nil
}

func (r *ReservationRepository) Create(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error) {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	id := uuid.NewString()
	Reservation := entity.Reservation{
		Id:           id,
		CustomerId:   customerId,
		BeauticianId: beauticianId,
		MenuId:       menuId,
		StartTime:    startTime,
		EndTime:      endTime,
		Price:        price,
	}
	if err := tx.Create(Reservation).Error; err != nil {
		return "", fmt.Errorf("failed to Create: %w", err)
	}
	return id, nil
}

func (r *ReservationRepository) Delete(ctx context.Context, id string) error {
	tx, ok := database.GetTx(ctx)
	if !ok {
		tx = r.db
	}
	Reservation := entity.Reservation{}
	if err := tx.Delete(&Reservation, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to Delete: %w", err)
	}
	return nil
}
