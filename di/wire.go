//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
)

func InitializeBeauticianController(db *gorm.DB) *controller.BeauticianController {
	wire.Build(controller.NewBeauticianController, usecase.NewBeauticianInteractor, repository.NewBeauticianRepository,
		wire.Bind(new(usecase.BeauticianRepositoryInterface), new(*repository.BeauticianRepository)),
		wire.Bind(new(entity.BeauticianInteractorInterface), new(*usecase.BeauticianInteractor)),
	)
	return &controller.BeauticianController{}
}

func InitializeMenuController(db *gorm.DB) *controller.MenuController {
	wire.Build(controller.NewMenuController, usecase.NewMenuInteractor, repository.NewMenuRepository,
		wire.Bind(new(usecase.MenuRepositoryInterface), new(*repository.MenuRepository)),
		wire.Bind(new(entity.MenuInteractorInterface), new(*usecase.MenuInteractor)),
	)
	return &controller.MenuController{}
}

func InitializeReservationController(db *gorm.DB) *controller.ReservationController {
	wire.Build(controller.NewReservationController, usecase.NewReservationInteractor, repository.NewReservationRepository, repository.NewMenuRepository, repository.NewBeauticianRepository, repository.NewCustomerRepository,
		wire.Bind(new(usecase.ReservationRepositoryInterface), new(*repository.ReservationRepository)),
		wire.Bind(new(usecase.MenuRepositoryInterface), new(*repository.MenuRepository)),
		wire.Bind(new(usecase.BeauticianRepositoryInterface), new(*repository.BeauticianRepository)),
		wire.Bind(new(usecase.CustomerRepositoryInterface), new(*repository.CustomerRepository)),
		wire.Bind(new(entity.ReservationInteractorInterface), new(*usecase.ReservationInteractor)),
	)
	return &controller.ReservationController{}
}

func InitializeBeauticianRepository() *repository.BeauticianRepository {
	wire.Build(repository.NewBeauticianRepository, database.NewDatabaseHandler)
	return &repository.BeauticianRepository{}
}

func InitializeMenuRepository() *repository.MenuRepository {
	wire.Build(repository.NewMenuRepository, database.NewDatabaseHandler)
	return &repository.MenuRepository{}
}

func InitializeReservationRepository() *repository.ReservationRepository {
	wire.Build(repository.NewReservationRepository, database.NewDatabaseHandler)
	return &repository.ReservationRepository{}
}
