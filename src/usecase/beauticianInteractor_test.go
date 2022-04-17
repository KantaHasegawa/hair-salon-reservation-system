package usecase

import (
	"testing"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/stretchr/testify/assert"
)

type mockBeauticianRepository struct {
	repository.BeauticianRepository
}

func (r *mockBeauticianRepository) All() ([]entity.Beautician, error) {
	return []entity.Beautician{}, nil
}

func TestGetBeauticians(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	_, err := i.GetBeauticians()
	assert.Nil(t, err)
}
