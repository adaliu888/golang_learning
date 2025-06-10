package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB
var err error

func DBIint() *gorm.DB {
	log.Printf("connecting to database...")
	// replace with your own database credentials

	dsn := `root:test1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`
	DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")

	// Migrate the schema

	//
	return DBConnect

}
