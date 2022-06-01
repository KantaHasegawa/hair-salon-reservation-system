package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/reservation"
	"github.com/stretchr/testify/assert"
)

func TestReservationAll(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler())
	result, _ := r.All(context.Background())
	assert.Equal(t, len(reservation.Seed), len(result))
}

func TestReservationFind(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find(context.Background(), "1")
	fmt.Print(result)
	assert.Equal(t, reservation.Seed[0].Price, result.Price)
}

func TestReservationCreate(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler().Begin())
	ctx := context.Background()
	_, err := r.Create(ctx, "1", "1", "1", time.Now(), time.Now().Add(1*time.Hour), 5000)
	assert.Nil(t, err)
	result, err := r.All(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(reservation.Seed)+1, len(result))
	r.db.Rollback()
}

func TestReservationDelete(t *testing.T) {
	r := NewReservationRepository(database.NewTestDatabaseHandler().Begin())
	ctx := context.Background()
	err := r.Delete(ctx, "3")
	assert.Nil(t, err)
	result, err := r.All(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(reservation.Seed)-1, len(result))
	r.db.Rollback()
}
