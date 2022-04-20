package customer

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

var Seed = []entity.Customer{
	{Id: "1", Name: "One", Sex: "M"},
	{Id: "2", Name: "Two", Sex: "F"},
	{Id: "3", Name: "Three", Sex: "M"},
}

func Factory(db *gorm.DB) error {
	var err error
	for _, value := range Seed {
		err = db.Create(&value).Error
	}
	if err != nil {
		return fmt.Errorf("failed to Customer factory: %w", err)
	}
	log.Println("Customer factory done")
	return nil
}
