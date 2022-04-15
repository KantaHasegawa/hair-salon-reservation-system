package repository

import "github.com/kantahasegawa/hair-salon-reservation-system/src/entity"

type BeauticianRepository struct{}

func NewBeauticianRepository() *BeauticianRepository {
	return &BeauticianRepository{}
}

func (r *BeauticianRepository) All() ([]entity.Beautician, error)
func (r *BeauticianRepository) Find(id string) (entity.Beautician, error)
func (r *BeauticianRepository) Create(name string, sex string, price int) error
func (r *BeauticianRepository) Update(id string) error
func (r *BeauticianRepository) Delete(id string) error
