package config

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

var gormConn *gorm.DB

func get_connection() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	response := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	return response
	//return "user:password@tcp(localhost:3306)/database_name"
}

func GetDatabaseConnection() *gorm.DB {
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	conn, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_CONNECTION"))
	if err != nil {
		logrus.Fatal("Could not connect to the database")
	}

	gormConn = conn

	return gormConn
}