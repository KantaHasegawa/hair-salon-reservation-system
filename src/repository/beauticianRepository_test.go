package repository

import (
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
