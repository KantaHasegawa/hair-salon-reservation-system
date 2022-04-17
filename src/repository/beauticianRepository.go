package repository

import (
	"fmt"

	"github.com/google/uuid"
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
func (r *BeauticianRepository) Create(name string, sex string, price int) (string, error) {
	id := uuid.NewString()
	beautician := entity.Beautician{
		Id: id,
		Name: name,
		Sex: sex,
		Price: price,
	}
	if err := r.db.Create(beautician).Error; err != nil {return "", fmt.Errorf("failed to Create: %w", err)}
	return id, nil
}
func (r *BeauticianRepository) Update(id string) error {
	fmt.Println("all")
	return nil
}
func (r *BeauticianRepository) Delete(id string) error {
	fmt.Println("all")
	return nil
}
