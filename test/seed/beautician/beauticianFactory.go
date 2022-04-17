package beautician

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

func Factory(db *gorm.DB) error {
	var seed = []entity.Beautician{
		{Id: "1", Name: "One", Sex: "M", Price: 1},
		{Id: "2", Name: "Two", Sex: "F", Price: 2},
		{Id: "3", Name: "Three", Sex: "M", Price: 3},
	}
	var err error
	for _, value := range seed {
		err = db.Create(&value).Error
	}
	if err != nil {
		return fmt.Errorf("failed to beautician factory: %w", err)
	}
	log.Println("beautician factory done")
	return nil
}
