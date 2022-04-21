package usecase

import (
	"testing"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/stretchr/testify/assert"
)

type mockReservationRepository struct {
	repository.ReservationRepository
}

func (r *mockReservationRepository) All() ([]entity.Reservation, error) {
	return []entity.Reservation{}, nil
}

func (r *mockReservationRepository) Find(id string) (entity.Reservation, error) {
	return entity.Reservation{}, nil
}

func (r *mockReservationRepository) Create(customerId string, beauticianId string, menuId string, startTime time.Time, endTime time.Time, price int) (string, error) {
	return "", nil
}

func (r *mockReservationRepository) Delete(id string) error {
	return nil
}

var loc, _ = time.LoadLocation("Asia/Tokyo")
var startTime = time.Date(2000, 1, 1, 0, 0, 0, 0, loc)

func TestGetReservations(t *testing.T) {
	i := NewReservationInteractor(&mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{})
	_, err := i.GetReservations()
	assert.Nil(t, err)
}

func TestGetReservation(t *testing.T) {
	i := NewReservationInteractor(&mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{})
	_, err := i.GetReservation("1")
	assert.Nil(t, err)
}

func TestNewReservation(t *testing.T) {
	i := NewReservationInteractor(&mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{})
	_, err := i.AddReservation("1", "1", "1", startTime)
	assert.Nil(t, err)
}

func TestDeleteReservation(t *testing.T) {
	i := NewReservationInteractor(&mockReservationRepository{}, &mockBeauticianRepository{}, &mockMenuRepository{})
	err := i.DeleteReservation("")
	assert.Nil(t, err)
}
