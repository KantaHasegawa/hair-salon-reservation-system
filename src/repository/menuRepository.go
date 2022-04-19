package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) All() ([]entity.Menu, error) {
	Menus := []entity.Menu{}
	err := r.db.Find(&Menus).Error
	if err != nil {
		return []entity.Menu{}, fmt.Errorf("failed to All: %w", err)
	}
	return Menus, nil
}
func (r *MenuRepository) Find(id string) (entity.Menu, error) {
	Menu := entity.Menu{}
	err := r.db.First(&Menu, "id = ?", id).Error
	if err != nil {
		return entity.Menu{}, fmt.Errorf("failed to Find: %w", err)
	}
	return Menu, nil
}
func (r *MenuRepository) Create(name string, sex string, price int, time int) (string, error) {
	id := uuid.NewString()
	Menu := entity.Menu{
		Id: id,
		Name: name,
		Sex: sex,
		Price: price,
		Time: time,
	}
	if err := r.db.Create(Menu).Error; err != nil {return "", fmt.Errorf("failed to Create: %w", err)}
	return id, nil
}
func (r *MenuRepository) Update(id string, name string, sex string, price int, time int) error {
	Menu := entity.Menu{}
	if err := r.db.First(&Menu, "id = ?", id).Error; err != nil {return fmt.Errorf("failed to Find: %w", err)}
	Menu.Name = name
	Menu.Sex = sex
	Menu.Price = price
	if err :=	r.db.Save(&Menu).Error; err != nil {return fmt.Errorf("failed to Save: %w", err) }
	return nil
}
func (r *MenuRepository) Delete(id string) error {
	Menu := entity.Menu{}
	if err := r.db.Delete(&Menu, "id = ?", id).Error; err != nil{return fmt.Errorf("failed to Delete: %w", err)}
	return nil
}
