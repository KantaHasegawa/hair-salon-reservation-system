package reservation

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kantahasegawa/hair-salon-reservation-system/src/entity"
)

var time1 = time.Now()
var time2 = time.Now().Add(1 * time.Hour)
var time3 = time.Now().Add(2 * time.Hour)
var time4 = time.Now().Add(3 * time.Hour)
var time5 = time.Now().Add(4 * time.Hour)
var time6 = time.Now().Add(5 * time.Hour)

var Seed = []entity.Reservation{
	{Id: "1", CustomerId: "1", BeauticianId: "1", MenuId: "1", StartTime: time1, EndTime: time2, Price: 5000},
	{Id: "2", CustomerId: "2", BeauticianId: "2", MenuId: "2", StartTime: time3, EndTime: time4, Price: 5000},
	{Id: "3", CustomerId: "3", BeauticianId: "3", MenuId: "3", StartTime: time5, EndTime: time6, Price: 5000},
}

func Factory(db *gorm.DB) error {
	var err error
	for _, value := range Seed {
		err = db.Create(&value).Error
	}
	if err != nil {
		return fmt.Errorf("failed to Reservation factory: %w", err)
	}
	log.Println("Reservation factory done")
	return nil
}
