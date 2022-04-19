//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/controller"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/database"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/repository"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/usecase"
)

func InitializeBeauticianController() *controller.BeauticianController {
	wire.Build(controller.NewBeauticianController, usecase.NewBeauticianInteractor, repository.NewBeauticianRepository, database.NewDatabaseHandler,
		wire.Bind(new(usecase.BeauticianRepositoryInterface), new(*repository.BeauticianRepository)),
		wire.Bind(new(entity.BeauticianInteractorInterface), new(*usecase.BeauticianInteractor)),
	)
	return &controller.BeauticianController{}
}

func InitializeMenuController() *controller.MenuController {
	wire.Build(controller.NewMenuController, usecase.NewMenuInteractor, repository.NewMenuRepository, database.NewDatabaseHandler,
		wire.Bind(new(usecase.MenuRepositoryInterface), new(*repository.MenuRepository)),
		wire.Bind(new(entity.MenuInteractorInterface), new(*usecase.MenuInteractor)),
	)
	return &controller.MenuController{}
}
