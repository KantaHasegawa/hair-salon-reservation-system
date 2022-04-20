package repository

import (
	"fmt"
	"testing"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/menu"
	"github.com/stretchr/testify/assert"
)

func TestMenuAll(t *testing.T) {
	r := NewMenuRepository(database.NewTestDatabaseHandler())
	result, _ := r.All()
	assert.Equal(t, len(menu.Seed), len(result))
}

func TestMenuFind(t *testing.T) {
	r := NewMenuRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find("1")
	fmt.Print(result)
	assert.Equal(t, menu.Seed[0].Name, result.Name)
}

func TestMenuCreate(t *testing.T) {
	r := NewMenuRepository(database.NewTestDatabaseHandler().Begin())
	_, err := r.Create("パーマ", "M", 10000, 60)
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(menu.Seed) + 1, len(result))
	r.db.Rollback()
}

func TestMenuUpdate(t *testing.T) {
	r := NewMenuRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Update("1", "update", "M", 10000, 60)
	assert.Nil(t, err)
	result, err := r.Find("1")
	assert.Nil(t, err)
	assert.Equal(t, "update", result.Name)
	r.db.Rollback()
}

func TestMenuDelete(t *testing.T){
	r := NewMenuRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Delete("3")
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(menu.Seed) - 1, len(result))
	r.db.Rollback()
}
