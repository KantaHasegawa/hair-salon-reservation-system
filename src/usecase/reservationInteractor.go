package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type ReservationRepositoryInterface interface {
	All(ctx context.Context) ([]entity.Reservation, error)
	Find(ctx context.Context, id string) (entity.Reservation, error)
	FindByBeauticianAndTime(ctx context.Context, customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error)
	FindByCustomerAndTime(ctx context.Context, customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error)
	Create(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error)
	Delete(ctx context.Context, id string) error
}

type CustomerRepositoryInterface interface {
	All() ([]entity.Customer, error)
	Find(id string) (entity.Customer, error)
	Create(name string, sex string) (string, error)
	Update(id string, name string, sex string) error
	Delete(id string) error
}

type ReservationInteractor struct {
	tx                    database.Transaction
	reservationRepository ReservationRepositoryInterface
	beauticianRepository  BeauticianRepositoryInterface
	menuRepository        MenuRepositoryInterface
	customerRepository    CustomerRepositoryInterface
}

func NewReservationInteractor(tx database.Transaction, reservationRepository ReservationRepositoryInterface, beauticianRepository BeauticianRepositoryInterface, menuRepository MenuRepositoryInterface, customerRepository CustomerRepositoryInterface) *ReservationInteractor {
	return &ReservationInteractor{tx, reservationRepository, beauticianRepository, menuRepository, customerRepository}
}

func (i *ReservationInteractor) GetReservations(ctx context.Context) ([]entity.Reservation, error) {
	result, err := i.reservationRepository.All(ctx)
	if err != nil {
		return result, fmt.Errorf("failed to ReservationRepository.All: %w", err)
	}
	return result, nil
}

func (i *ReservationInteractor) GetReservation(ctx context.Context, id string) (entity.Reservation, error) {
	result, err := i.reservationRepository.Find(ctx, id)
	if err != nil {
		return result, fmt.Errorf("failed to ReservationRepository.Find: %w", err)
	}
	return result, err
}

func (i *ReservationInteractor) AddReservation(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time) (string, error) {
	if err := validateReservationInput(customerId, beauticianId, menuId, startTime); err != nil {
		return "", err
	}
	result, err := i.tx.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		_, err := i.customerRepository.Find(customerId)
		if err != nil {
			return "", fmt.Errorf("failed to CustomerRepository.Find: %w", err)
		}
		beautician, err := i.beauticianRepository.Find(beauticianId)
		if err != nil {
			return "", fmt.Errorf("failed to BeauticianRepository.Find: %w", err)
		}
		menu, err := i.menuRepository.Find(menuId)
		if err != nil {
			return "", fmt.Errorf("failed to MenuRepository.Find: %w", err)
		}

		price := beautician.Price + menu.Price
		endTime := startTime.Add(time.Duration(menu.Time) * time.Minute)

		duplicateBeauticianReservation, err := i.reservationRepository.FindByBeauticianAndTime(ctx, beauticianId, startTime, endTime)
		if err != nil {
			return "", fmt.Errorf("failed to ReservationRepository.FindByBeauticianAndTime: %w", err)
		}
		if len(duplicateBeauticianReservation) != 0 {
			return "", fmt.Errorf("failed to AddReservation validation(time duplicate): %w", errors.New("bad request"))
		}

		duplicateCustomerReservation, err := i.reservationRepository.FindByCustomerAndTime(ctx, customerId, startTime, endTime)
		if err != nil {
			return "", fmt.Errorf("failed to ReservationRepository.FindByCustomerAndTime: %w", err)
		}
		if len(duplicateCustomerReservation) != 0 {
			return "", fmt.Errorf("failed to AddReservation validation(time duplicate): %w", errors.New("bad request"))
		}

		result, err := i.reservationRepository.Create(ctx, customerId, beauticianId, menuId, startTime, endTime, price)

		if err != nil {
			return result, fmt.Errorf("failed to ReservationRepository.Create: %w", err)
		}
		return result, err
	})

	s, ok := result.(string)
	if !ok {
		return "", err
	}
	return s, err
}

func (i *ReservationInteractor) DeleteReservation(ctx context.Context, id string) error {
	err := i.reservationRepository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to ReservationRepository.Delete: %w", err)
	}
	return err
}

func validateReservationInput(customerId string, beauticianId string, menuId string, startTime time.Time) error {
	return nil
}
