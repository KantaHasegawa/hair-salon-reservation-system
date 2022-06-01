package usecase

import (
	"context"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"testing"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/stretchr/testify/assert"
)

type mockCustomerRepository struct {
	repository.CustomerRepository
}

func (r *mockCustomerRepository) Find(id string) (entity.Customer, error) {
	return entity.Customer{}, nil
}

type mockReservationRepository struct {
	repository.ReservationRepository
}

func (r *mockReservationRepository) All(ctx context.Context) ([]entity.Reservation, error) {
	return []entity.Reservation{}, nil
}

func (r *mockReservationRepository) Find(ctx context.Context, id string) (entity.Reservation, error) {
	return entity.Reservation{}, nil
}

func (r *mockReservationRepository) FindByBeauticianAndTime(ctx context.Context, customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error) {
	return []entity.Reservation{}, nil
}

func (r *mockReservationRepository) FindByCustomerAndTime(ctx context.Context, customerId string, startTime time.Time, endTime time.Time) ([]entity.Reservation, error) {
	return []entity.Reservation{}, nil
}

func (r *mockReservationRepository) Create(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error) {
	return "", nil
}

func (r *mockReservationRepository) Delete(ctx context.Context, id string) error {
	return nil
}

var loc, _ = time.LoadLocation("Asia/Tokyo")
var startTime = time.Date(2000, 1, 1, 0, 0, 0, 0, loc)

func TestGetReservations(t *testing.T) {
	i := NewReservationInteractor(database.NewTransaction(database.NewTestDatabaseHandler()), &mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{}, &mockCustomerRepository{})
	_, err := i.GetReservations(context.Background())
	assert.Nil(t, err)
}

func TestGetReservation(t *testing.T) {
	i := NewReservationInteractor(database.NewTransaction(database.NewTestDatabaseHandler()), &mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{}, &mockCustomerRepository{})
	_, err := i.GetReservation(context.Background(), "1")
	assert.Nil(t, err)
}

func TestNewReservation(t *testing.T) {
	i := NewReservationInteractor(database.NewTransaction(database.NewTestDatabaseHandler()), &mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{}, &mockCustomerRepository{})
	_, err := i.AddReservation(context.Background(), "1", "1", "1", startTime)
	assert.Nil(t, err)
}

func TestDeleteReservation(t *testing.T) {
	i := NewReservationInteractor(database.NewTransaction(database.NewTestDatabaseHandler()), &mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{}, &mockCustomerRepository{})
	err := i.DeleteReservation(context.Background(), "")
	assert.Nil(t, err)
}
