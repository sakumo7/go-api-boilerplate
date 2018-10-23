package main

import (
	"fmt"
	"./models"
)

func main(){
	user := new(models.User)
	models.ExecuteMigrations()
	users := user.FetchAll()
	fmt.Println(users)
}