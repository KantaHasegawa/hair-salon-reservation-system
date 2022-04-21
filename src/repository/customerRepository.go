package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) All() ([]entity.Customer, error) {
	Customers := []entity.Customer{}
	err := r.db.Find(&Customers).Error
	if err != nil {
		return []entity.Customer{}, fmt.Errorf("failed to All: %w", err)
	}
	return Customers, nil
}

func (r *CustomerRepository) Find(id string) (entity.Customer, error) {
	Customer := entity.Customer{}
	err := r.db.First(&Customer, "id = ?", id).Error
	if err != nil {
		return entity.Customer{}, fmt.Errorf("failed to Find: %w", err)
	}
	return Customer, nil
}

func (r *CustomerRepository) Create(name string, sex string) (string, error) {
	id := uuid.NewString()
	Customer := entity.Customer{
		Id:   id,
		Name: name,
		Sex:  sex,
	}
	if err := r.db.Create(Customer).Error; err != nil {
		return "", fmt.Errorf("failed to Create: %w", err)
	}
	return id, nil
}

func (r *CustomerRepository) Update(id string, name string, sex string) error {
	Customer := entity.Customer{}
	if err := r.db.First(&Customer, "id = ?", id).Error; err != nil {return fmt.Errorf("failed to Find: %w", err)}
	Customer.Name = name
	Customer.Sex = sex
	if err :=	r.db.Save(&Customer).Error; err != nil {return fmt.Errorf("failed to Save: %w", err) }
	return nil
}

func (r *CustomerRepository) Delete(id string) error {
	Customer := entity.Customer{}
	if err := r.db.Delete(&Customer, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to Delete: %w", err)
	}
	return nil
}
