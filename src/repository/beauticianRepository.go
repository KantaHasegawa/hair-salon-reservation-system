package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type BeauticianRepository struct {
	db *gorm.DB
}

func NewBeauticianRepository(db *gorm.DB) *BeauticianRepository {
	return &BeauticianRepository{db: db}
}

func (r *BeauticianRepository) All() ([]entity.Beautician, error) {
	beauticians := []entity.Beautician{}
	err := r.db.Find(&beauticians).Error
	if err != nil {
		return []entity.Beautician{}, fmt.Errorf("failed to All: %w", err)
	}
	return beauticians, nil
}
func (r *BeauticianRepository) Find(id string) (entity.Beautician, error) {
	beautician := entity.Beautician{}
	err := r.db.First(&beautician, "id = ?", id).Error
	if err != nil {
		return entity.Beautician{}, fmt.Errorf("failed to Find: %w", err)
	}
	return beautician, nil
}
func (r *BeauticianRepository) Create(name string, sex string, price int) error {
	fmt.Println("all")
	return nil
}
func (r *BeauticianRepository) Update(id string) error {
	fmt.Println("all")
	return nil
}
func (r *BeauticianRepository) Delete(id string) error {
	fmt.Println("all")
	return nil
}
