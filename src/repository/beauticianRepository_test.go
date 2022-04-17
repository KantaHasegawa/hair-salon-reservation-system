package repository

import (
	"fmt"
	"testing"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/beautician"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	r := NewBeauticianRepository(database.NewTestDatabaseHandler())
	result, _ := r.All()
	assert.Equal(t, len(beautician.Seed), len(result))
}

func TestFind(t *testing.T) {
	r := NewBeauticianRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find("1")
	fmt.Print(result)
	assert.Equal(t, beautician.Seed[0].Name, result.Name)
}

func TestCreate(t *testing.T) {
	r := NewBeauticianRepository(database.NewTestDatabaseHandler().Begin())
	_, err := r.Create("Bob", "M", 10000)
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(beautician.Seed) + 1, len(result))
	r.db.Rollback()
}

func TestUpdate(t *testing.T) {
	r := NewBeauticianRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Update("1", "update", "M", 10000)
	assert.Nil(t, err)
	result, err := r.Find("1")
	assert.Nil(t, err)
	assert.Equal(t, "update", result.Name)
	r.db.Rollback()
}

func TestDelete(t *testing.T){
	r := NewBeauticianRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Delete("3")
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(beautician.Seed) - 1, len(result))
	r.db.Rollback()
}
