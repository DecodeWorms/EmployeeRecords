package driver

import (
	"log"

	"github.com/go-gorm/postgres"
	"gorm.io/gorm"
)

var err error

var db *gorm.DB

func ConnDb() *gorm.DB {
	dsn := "host = localhost dbname = employeerecord user = Oluwashola sslmode = disable port = 5432"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
