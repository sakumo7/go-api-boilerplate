package models

import (
	"../core/database"
)

func ExecuteMigrations(){
	db := database.GetDatabaseConnection()
	db.AutoMigrate(&User{})
}