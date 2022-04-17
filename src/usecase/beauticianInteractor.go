package usecase

import (
	"errors"
	"fmt"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type BeauticianRepositoryInterface interface {
	All() ([]entity.Beautician, error)
	Find(id string) (entity.Beautician, error)
	Create(name string, sex string, price int) (string, error)
	Update(id string, name string, sex string, price int) error
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
	if err != nil {
		return result, fmt.Errorf("failed to BeauticianRepository.Find: %w", err)
	}
	return result, err
}

func (i *BeauticianInteractor) AddBeautician(name string, sex string, price int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("failed to AddBeautician validation(name): %w", errors.New("bad request"))
	}
	if sex == "" || (sex != "M" && sex != "F") {
		return "", fmt.Errorf("failed to AddBeautician validation(sex): %w", errors.New("bad request"))
	}
	if price < 0 {
		return "", fmt.Errorf("failed to AddBeautician validation(price): %w", errors.New("bad request"))
	}

	result, err := i.repository.Create(name, sex, price)
	if err := validateInput(name, sex, price); err != nil {
		return "", err
	}
	if err != nil {
		return result, fmt.Errorf("failed to BeauticianRepository.Create: %w", err)
	}
	return result, err
}

func (i *BeauticianInteractor) UpdateBeautician(id string, name string, sex string, price int) error {
	if err := validateInput(name, sex, price); err != nil {
		return err
	}
	err := i.repository.Update(id, name, sex, price)
	if err != nil {
		return fmt.Errorf("failed to BeauticianRepository.Update: %w", err)
	}
	return nil
}

func (i *BeauticianInteractor) DeleteBeautician(id string) error {
	err := i.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to BeauticianRepository.Delete: %w", err)
	}
	return err
}

func validateInput(name string, sex string, price int) error {
	if name == "" {
		return fmt.Errorf("failed to AddBeautician validation(name): %w", errors.New("bad request"))
	}
	if sex == "" || (sex != "M" && sex != "F") {
		return fmt.Errorf("failed to AddBeautician validation(sex): %w", errors.New("bad request"))
	}
	if price < 0 {
		return fmt.Errorf("failed to AddBeautician validation(price): %w", errors.New("bad request"))
	}
	return nil
}
