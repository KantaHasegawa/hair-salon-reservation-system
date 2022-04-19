// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
)

// Injectors from wire.go:

func InitializeBeauticianController() *controller.BeauticianController {
	db := database.NewDatabaseHandler()
	beauticianRepository := repository.NewBeauticianRepository(db)
	beauticianInteractor := usecase.NewBeauticianInteractor(beauticianRepository)
	beauticianController := controller.NewBeauticianController(beauticianInteractor)
	return beauticianController
}

func InitializeMenuController() *controller.MenuController {
	db := database.NewDatabaseHandler()
	menuRepository := repository.NewMenuRepository(db)
	menuInteractor := usecase.NewMenuInteractor(menuRepository)
	menuController := controller.NewMenuController(menuInteractor)
	return menuController
}
