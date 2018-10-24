package config

import (
	"os"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var gormConn *gorm.DB

func getConnection() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("SSL")
	connection_string := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, name, password)
	if ssl == ""{
		connection_string = connection_string+" sslmode=disable"	
	}
	return connection_string
}

func GetDatabaseConnection() *gorm.DB {
	if gormConn != nil && gormConn.DB() != nil && gormConn.DB().Ping() == nil {
		return gormConn
	}

	conn, err := gorm.Open("postgres", getConnection())
	if err != nil {
		fmt.Println(err)
		logrus.Fatal("Could not connect to the database")
	}

	gormConn = conn

	return gormConn
}