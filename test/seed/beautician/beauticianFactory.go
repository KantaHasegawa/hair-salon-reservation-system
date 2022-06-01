package beautician

import (
	"fmt"
	"log"

	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
	"gorm.io/gorm"
)

var Seed = []entity.Beautician{
	{Id: "1", Name: "One", Sex: "M", Price: 1},
	{Id: "2", Name: "Two", Sex: "F", Price: 2},
	{Id: "3", Name: "Three", Sex: "M", Price: 3},
}

func Factory(db *gorm.DB) error {
	var err error
	for _, value := range Seed {
		err = db.Create(&value).Error
	}
	if err != nil {
		return fmt.Errorf("failed to beautician factory: %w", err)
	}
	log.Println("beautician factory done")
	return nil
}
