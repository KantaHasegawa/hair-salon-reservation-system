package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/reservation"
	"github.com/stretchr/testify/assert"
)

func TestReservationAll(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler())
	result, _ := r.All()
	assert.Equal(t, len(reservation.Seed), len(result))
}

func TestReservationFind(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find("1")
	fmt.Print(result)
	assert.Equal(t, reservation.Seed[0].Price, result.Price)
}

func TestReservationCreate(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler().Begin())
	_, err := r.Create("1", "1", "1", time.Now(), time.Now().Add(1 * time.Hour), 5000)
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(reservation.Seed)+1, len(result))
	r.db.Rollback()
}

func TestReservationDelete(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Delete("3")
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(reservation.Seed)-1, len(result))
	r.db.Rollback()
}
