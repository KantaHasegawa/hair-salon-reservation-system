package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
	"github.com/stretchr/testify/assert"
)

type mockReservationInteractor struct {
	usecase.ReservationInteractor
}

var loc, _ = time.LoadLocation("Asia/Tokyo")
var startTime = time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
var endTime = startTime.Add(1 * time.Hour)

func (i *mockReservationInteractor) GetReservations(ctx context.Context) ([]entity.Reservation, error) {
	result := []entity.Reservation{{Id: "1", CustomerId: "1", BeauticianId: "1", MenuId: "1", StartTime: startTime, EndTime: endTime, Price: 5000},
		{Id: "2", CustomerId: "2", BeauticianId: "2", MenuId: "2", StartTime: startTime, EndTime: endTime, Price: 5000}}
	return result, nil
}

func (i *mockReservationInteractor) GetReservation(ctx context.Context, id string) (entity.Reservation, error) {
	result := entity.Reservation{Id: "1", CustomerId: "1", BeauticianId: "1", MenuId: "1", StartTime: startTime, EndTime: endTime, Price: 5000}
	return result, nil
}

func (i *mockReservationInteractor) AddReservation(ctx context.Context, customerId string, beauticianId string, menuId string, startTime time.Time) (string, error) {
	return "id", nil
}

func (i *mockReservationInteractor) DeleteReservation(ctx context.Context, id string) error {
	return nil
}

func TestReservationIndexHndler(t *testing.T) {
	expect := "{\"result\":[{\"id\":\"1\",\"customer_id\":\"1\",\"beautician_id\":\"1\",\"menu_id\":\"1\",\"start_time\":\"2000-01-01T00:00:00+09:00\",\"end_time\":\"2000-01-01T01:00:00+09:00\",\"price\":5000},{\"id\":\"2\",\"customer_id\":\"2\",\"beautician_id\":\"2\",\"menu_id\":\"2\",\"start_time\":\"2000-01-01T00:00:00+09:00\",\"end_time\":\"2000-01-01T01:00:00+09:00\",\"price\":5000}]}"
	controller := NewReservationController(&mockReservationInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Reservations",
		nil,
	)
	controller.IndexHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestReservationShowHandler(t *testing.T) {
	expect := "{\"result\":{\"id\":\"1\",\"customer_id\":\"1\",\"beautician_id\":\"1\",\"menu_id\":\"1\",\"start_time\":\"2000-01-01T00:00:00+09:00\",\"end_time\":\"2000-01-01T01:00:00+09:00\",\"price\":5000}}"
	controller := NewReservationController(&mockReservationInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Reservation/1",
		nil,
	)
	controller.ShowHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestReservationNewHandler(t *testing.T) {
	expect := "{\"id\":\"id\"}"
	jsonStr := `{"name":"test","sex":"M","price":1000}`
	controller := NewReservationController(&mockReservationInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodPost,
		"/Reservation",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	c.Header("Content-Type", "application/json")
	controller.NewHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}

func TestReservationDeleteHandler(t *testing.T) {
	expect := "{\"message\":\"success\"}"
	controller := NewReservationController(&mockReservationInteractor{})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(
		http.MethodGet,
		"/Reservation/1",
		nil,
	)
	controller.DeleteHandler(c)
	assert.Equal(t, 200, response.Code)
	assert.Equal(t, expect, response.Body.String())
}
