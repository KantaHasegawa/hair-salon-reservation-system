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

func InitializeReservationController() *controller.ReservationController {
	db := database.NewDatabaseHandler()
	reservationRepository := repository.NewReservationRepository(db)
	beauticianRepository := repository.NewBeauticianRepository(db)
	menuRepository := repository.NewMenuRepository(db)
	reservationInteractor := usecase.NewReservationInteractor(reservationRepository, beauticianRepository, menuRepository)
	reservationController := controller.NewReservationController(reservationInteractor)
	return reservationController
}

func InitializeBeauticianRepository() *repository.BeauticianRepository {
	db := database.NewDatabaseHandler()
	beauticianRepository := repository.NewBeauticianRepository(db)
	return beauticianRepository
}

func InitializeMenuRepository() *repository.MenuRepository {
	db := database.NewDatabaseHandler()
	menuRepository := repository.NewMenuRepository(db)
	return menuRepository
}

func InitializeReservationRepository() *repository.ReservationRepository {
	db := database.NewDatabaseHandler()
	reservationRepository := repository.NewReservationRepository(db)
	return reservationRepository
}
