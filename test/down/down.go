package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func settingMigration(path string, url string) (*migrate.Migrate, error){
	m, err := migrate.New("file://" + path, url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return m, nil
}

func main(){
	log.Println("down migration start")

  m, err :=	settingMigration("db/migrations", "mysql://root:@tcp(127.0.0.1:4306)/hs_reservation_test")
	if err != nil {
		panic(err)
	}
	err = m.Down()
	if err != nil {
		panic(err)
	}
	log.Println("down migration done")
}
