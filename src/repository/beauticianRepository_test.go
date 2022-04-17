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

func TestFind(t *testing.T){
	r := NewBeauticianRepository(database.NewTestDatabaseHandler())
	result, _ := r.Find("1")
	fmt.Print(result)
	assert.Equal(t, beautician.Seed[0].Name, result.Name)
}
