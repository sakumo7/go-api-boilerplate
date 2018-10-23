package main

import (
	"fmt"
	"./models"
)

func main(){
	user := new(models.User)
	users := user.FetchAll()
	fmt.Println(users)
}