package entity

import "time"

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
	GetReservations() ([]Reservation, error)
	GetReservation(id string) (Reservation, error)
	AddReservation(customerId string, beauticianId string, menuId string, startTime time.Time) (string, error)
	DeleteReservation(id string) error
}
