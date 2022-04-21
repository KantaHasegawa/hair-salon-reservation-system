package repository

import (
	"fmt"
	"testing"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/test/seed/customer"
	"github.com/stretchr/testify/assert"
)

func TestCustomerAll(t *testing.T) {
	r := NewCustomerRepository(database.NewTestDatabaseHandler())
	result, _ := r.All()
	assert.Equal(t, len(customer.Seed), len(result))
}

func TestCustomerFind(t *testing.T) {
	r := NewCustomerRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find("1")
	fmt.Print(result)
	assert.Equal(t, customer.Seed[0].Name, result.Name)
}

func TestCustomerCreate(t *testing.T) {
	r := NewCustomerRepository(database.NewTestDatabaseHandler().Begin())
	_, err := r.Create("Bob", "M")
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(customer.Seed) + 1, len(result))
	r.db.Rollback()
}

func TestCustomerUpdate(t *testing.T) {
	r := NewCustomerRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Update("1", "update", "M")
	assert.Nil(t, err)
	result, err := r.Find("1")
	assert.Nil(t, err)
	assert.Equal(t, "update", result.Name)
	r.db.Rollback()
}

func TestCustomerDelete(t *testing.T){
	r := NewCustomerRepository(database.NewTestDatabaseHandler().Begin())
	err := r.Delete("3")
	assert.Nil(t, err)
	result, err := r.All()
	assert.Nil(t, err)
	assert.Equal(t, len(customer.Seed) - 1, len(result))
	r.db.Rollback()
}
