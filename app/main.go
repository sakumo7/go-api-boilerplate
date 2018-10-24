package main

import (
	"log"
	"net/http"
	"os"
	"github.com/sirupsen/logrus"
	"./settings"
	"./models"
	"./server"
)

func main() {
	db := settings.GetDatabaseConnection()
	port := os.Getenv("APP_PORT")
	
	defer db.Close()
	logrus.Info("Version is ", "1.0")
	logrus.Info("Starting Server on http://localhost:"+port)
	
	models.ExecuteMigrations()
	server := server.NewServer()
	log.Fatal(http.ListenAndServe(":"+port, server))
}