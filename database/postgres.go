package database

import (
	"assigment-2/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "terserah"
	dbname   = "assignment-2"
	db       *gorm.DB
	err      error
)

func NewPostgres() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if !db.HasTable(&models.Orders{}) && !db.HasTable(&models.Items{}) {
		db.AutoMigrate(&models.Orders{}, &models.Items{})
	}

	return db
}
