package usecase

import (
	"errors"
	"fmt"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

type MenuRepositoryInterface interface {
	All() ([]entity.Menu, error)
	Find(id string) (entity.Menu, error)
	Create(name string, sex string, price int, time int) (string, error)
	Update(id string, name string, sex string, price int, time int) error
	Delete(id string) error
}

type MenuInteractor struct {
	repository MenuRepositoryInterface
}

func NewMenuInteractor(repository MenuRepositoryInterface) *MenuInteractor {
	return &MenuInteractor{repository}
}

func (i *MenuInteractor) GetMenus() ([]entity.Menu, error) {
	result, err := i.repository.All()
	if err != nil {
		return result, fmt.Errorf("failed to MenuRepository.All: %w", err)
	}
	return result, nil
}

func (i *MenuInteractor) GetMenu(id string) (entity.Menu, error) {
	result, err := i.repository.Find(id)
	if err != nil {
		return result, fmt.Errorf("failed to MenuRepository.Find: %w", err)
	}
	return result, err
}

func (i *MenuInteractor) AddMenu(name string, sex string, price int, time int) (string, error) {
	result, err := i.repository.Create(name, sex, price, time)
	if err := validateMenuInput(name, sex, price, time); err != nil {
		return "", err
	}
	if err != nil {
		return result, fmt.Errorf("failed to MenuRepository.Create: %w", err)
	}
	return result, err
}

func (i *MenuInteractor) UpdateMenu(id string, name string, sex string, price int, time int) error {
	if err := validateMenuInput(name, sex, price, time); err != nil {
		return err
	}
	err := i.repository.Update(id, name, sex, price, time)
	if err != nil {
		return fmt.Errorf("failed to MenuRepository.Update: %w", err)
	}
	return nil
}

func (i *MenuInteractor) DeleteMenu(id string) error {
	err := i.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to MenuRepository.Delete: %w", err)
	}
	return err
}

func validateMenuInput(name string, sex string, price int, time int) error {
	const MIN_TIME = 30
	const MAX_TIME = 720
	if name == "" {
		return fmt.Errorf("failed to AddMenu validation(name): %w", errors.New("bad request"))
	}
	if sex == "" || (sex != "M" && sex != "F") {
		return fmt.Errorf("failed to AddMenu validation(sex): %w", errors.New("bad request"))
	}
	if price < 0 {
		return fmt.Errorf("failed to AddMenu validation(price): %w", errors.New("bad request"))
	}
	if time < MIN_TIME || MAX_TIME < time {
		return fmt.Errorf("failed to AddMenu validation(time): %w", errors.New("bad request"))
	}
	return nil
}
