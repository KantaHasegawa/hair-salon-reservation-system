package usecase

import (
	"fmt"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type ReservationRepositoryInterface interface {
	All() ([]entity.Reservation, error)
	Find(id string) (entity.Reservation, error)
	Create(customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error)
	Delete(id string) error
}

type ReservationInteractor struct {
	repository ReservationRepositoryInterface
}

func NewReservationInteractor(repository ReservationRepositoryInterface) *ReservationInteractor {
	return &ReservationInteractor{repository}
}

func (i *ReservationInteractor) GetReservations() ([]entity.Reservation, error) {
	result, err := i.repository.All()
	if err != nil {
		return result, fmt.Errorf("failed to ReservationRepository.All: %w", err)
	}
	return result, nil
}

func (i *ReservationInteractor) GetReservation(id string) (entity.Reservation, error) {
	result, err := i.repository.Find(id)
	if err != nil {
		return result, fmt.Errorf("failed to ReservationRepository.Find: %w", err)
	}
	return result, err
}

func (i *ReservationInteractor) AddReservation(customerId string, beauticianId string, menuId string, startTime time.Time) (string, error) {

	endTime := time.Now()
	price := 5000

	result, err := i.repository.Create(customerId, beauticianId, menuId, startTime, endTime, price)
	if err := validateReservationInput(customerId, beauticianId, menuId, startTime); err != nil {
		return "", err
	}
	if err != nil {
		return result, fmt.Errorf("failed to ReservationRepository.Create: %w", err)
	}
	return result, err
}

func (i *ReservationInteractor) DeleteReservation(id string) error {
	err := i.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to ReservationRepository.Delete: %w", err)
	}
	return err
}

func validateReservationInput(customerId string, beauticianId string, menuId string, startTime time.Time) error {
	return nil
}
