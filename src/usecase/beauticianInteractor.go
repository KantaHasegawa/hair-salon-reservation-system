package usecase

import (
	"fmt"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type BeauticianRepositoryInterface interface {
	All() ([]entity.Beautician, error)
	Find(id string) (entity.Beautician, error)
	Create(name string, sex string, price int) error
	Update(id string) error
	Delete(id string) error
}

type BeauticianInteractor struct {
	repository BeauticianRepositoryInterface
}

func NewBeauticianInteractor(repository BeauticianRepositoryInterface) *BeauticianInteractor {
	return &BeauticianInteractor{repository}
}

func (i *BeauticianInteractor) GetBeauticians() ([]entity.Beautician, error) {
	result, err := i.repository.All()
	if err != nil {
		return result, fmt.Errorf("failed to BeauticianRepository.All: %w", err)
	}
	return result, nil
}

func (i *BeauticianInteractor) GetBeautician(id string) (entity.Beautician, error) {
	result, err := i.repository.Find(id)
	return result, err
}

func (i *BeauticianInteractor) AddBeautician(name string, sex string, price int) error {
	err := i.repository.Create(name, sex, price)
	return err
}

func (i *BeauticianInteractor) UpdateBeautician(id string) error {
	err := i.repository.Update(id)
	return err
}

func (i *BeauticianInteractor) DeleteBeautician(id string) error {
	err := i.repository.Delete(id)
	return err
}
