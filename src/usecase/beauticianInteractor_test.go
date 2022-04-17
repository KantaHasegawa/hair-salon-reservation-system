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

func (r *mockBeauticianRepository) Find(id string) (entity.Beautician, error) {
	return entity.Beautician{}, nil
}

func (r *mockBeauticianRepository) Create(name string, sex string, price int) (string, error) {
	return "", nil
}

func (r *mockBeauticianRepository) Update(id string, name string, sex string, price int) error {
	return nil
}

func (r *mockBeauticianRepository) Delete(id string) error {
	return nil
}

func TestGetBeauticians(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	_, err := i.GetBeauticians()
	assert.Nil(t, err)
}

func TestGetBeautician(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	_, err := i.GetBeautician("1")
	assert.Nil(t, err)
}

func TestNewBeautician(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	_, err := i.AddBeautician("", "", 0)
	assert.Nil(t, err)
}

func TestUpdateBeautician(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	err := i.UpdateBeautician("", "", "", 0)
	assert.Nil(t, err)
}

func TestDeleteBeautician(t *testing.T) {
	i := NewBeauticianInteractor(&mockBeauticianRepository{})
	err := i.DeleteBeautician("")
	assert.Nil(t, err)
}
