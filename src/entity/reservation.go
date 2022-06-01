package entity

import (
	"golang.org/x/net/context"
	"time"
)

type Reservation struct {
	Id           string    `json:"id"`
	CustomerId   string    `json:"customer_id"`
	BeauticianId string    `json:"beautician_id"`
	MenuId       string    `json:"menu_id"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Price        int       `json:"price"`
}

type ReservationInteractorInterface interface {
	GetReservations(ctx context.Context) ([]Reservation, error)
	GetReservation(ctx context.Context, id string) (Reservation, error)
	AddReservation(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time) (string, error)
	DeleteReservation(ctx context.Context, id string) error
}
