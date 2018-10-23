package main

import (
	"log"
	"net/http"
	//"os"

	//"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"./config"
	"./models"
	"./server"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	logrus.Fatal("Error loading .env file")
	//}
	
	db := config.GetDatabaseConnection()
	defer db.Close()
	logrus.Info("Version is ", "1.0")
	logrus.Info("Starting Server on http://localhost:8000")
	models.ExecuteMigrations()
	server := server.NewServer()
	log.Fatal(http.ListenAndServe(":8000", server))
}