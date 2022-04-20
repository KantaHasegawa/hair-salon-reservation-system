package usecase

import (
	"testing"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/stretchr/testify/assert"
)

type mockMenuRepository struct {
	repository.MenuRepository
}

func (r *mockMenuRepository) All() ([]entity.Menu, error) {
	return []entity.Menu{}, nil
}

func (r *mockMenuRepository) Find(id string) (entity.Menu, error) {
	return entity.Menu{}, nil
}

func (r *mockMenuRepository) Create(name string, sex string, price int, time int) (string, error) {
	return "", nil
}

func (r *mockMenuRepository) Update(id string, name string, sex string, price int, time int) error {
	return nil
}

func (r *mockMenuRepository) Delete(id string) error {
	return nil
}

func TestGetMenus(t *testing.T) {
	i := NewMenuInteractor(&mockMenuRepository{})
	_, err := i.GetMenus()
	assert.Nil(t, err)
}

func TestGetMenu(t *testing.T) {
	i := NewMenuInteractor(&mockMenuRepository{})
	_, err := i.GetMenu("1")
	assert.Nil(t, err)
}

func TestNewMenu(t *testing.T) {
	i := NewMenuInteractor(&mockMenuRepository{})
	_, err := i.AddMenu("", "F", 0, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(name): bad request")
	_, err = i.AddMenu("test", "male", 0, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(sex): bad request")
	_, err = i.AddMenu("test", "F", -5, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(price): bad request")
	_, err = i.AddMenu("test", "F", 0, 0)
	assert.Nil(t, err)
}

func TestUpdateMenu(t *testing.T) {
	i := NewMenuInteractor(&mockMenuRepository{})
	err := i.UpdateMenu("1", "", "F", 0, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(name): bad request")
	err = i.UpdateMenu("1", "test", "male", 0, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(sex): bad request")
	err = i.UpdateMenu("1", "test", "F", -5, 0)
	assert.EqualError(t, err, "failed to AddMenu validation(price): bad request")
	err = i.UpdateMenu("1", "test", "F", 0, 0)
	assert.Nil(t, err)
}

func TestDeleteMenu(t *testing.T) {
	i := NewMenuInteractor(&mockMenuRepository{})
	err := i.DeleteMenu("")
	assert.Nil(t, err)
}
