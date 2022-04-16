package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type BeauticianRepository struct{
	db *gorm.DB
}

func NewBeauticianRepository(db *gorm.DB) *BeauticianRepository {
	return &BeauticianRepository{db: db}
}

func (r *BeauticianRepository) All() ([]entity.Beautician, error) {
	beauticiansRecord := []entity.Beautician{}
	r.db.Find(&beauticiansRecord)
	fmt.Println(beauticiansRecord)
	return nil, nil
}
func (r *BeauticianRepository) Find(id string) (entity.Beautician, error) {
	fmt.Println("all")
	return entity.Beautician{}, nil
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
