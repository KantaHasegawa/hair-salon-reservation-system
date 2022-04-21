package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (r *ReservationRepository) All() ([]entity.Reservation, error) {
	Reservations := []entity.Reservation{}
	err := r.db.Find(&Reservations).Error
	if err != nil {
		return []entity.Reservation{}, fmt.Errorf("failed to All: %w", err)
	}
	return Reservations, nil
}
func (r *ReservationRepository) Find(id string) (entity.Reservation, error) {
	Reservation := entity.Reservation{}
	err := r.db.First(&Reservation, "id = ?", id).Error
	if err != nil {
		return entity.Reservation{}, fmt.Errorf("failed to Find: %w", err)
	}
	return Reservation, nil
}
func (r *ReservationRepository) FindByBeauticianAndTime(customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error) {
	Reservations := []entity.Reservation{}
	err := r.db.Find(&Reservations, "customer_id = ? AND ? < end_time and ? > start_time ", customerId, startTime, endTime).Error
	if err != nil {
		return []entity.Reservation{}, fmt.Errorf("failed to All: %w", err)
	}
	return Reservations, nil
}

func (r *ReservationRepository) Create(customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error) {
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
	if err := r.db.Create(Reservation).Error; err != nil {
		return "", fmt.Errorf("failed to Create: %w", err)
	}
	return id, nil
}

func (r *ReservationRepository) Delete(id string) error {
	Reservation := entity.Reservation{}
	if err := r.db.Delete(&Reservation, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to Delete: %w", err)
	}
	return nil
}
