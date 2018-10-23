package models

import (
	"../config"
)

func ExecuteMigrations(){
	db := config.GetDatabaseConnection()
	db.AutoMigrate(&User{})
}