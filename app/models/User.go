package models

import (
	"../config"
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
	FirstName string     `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string     `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     string     `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) FetchAll() []map[string]interface{} {
	db := config.GetDatabaseConnection()

	var users []User
	db.Find(&users)
	var usersMap []map[string]interface{}

	for index := range users {
		usersMap = append(usersMap, users[index].ConvertToMap())
	}

	return usersMap
}

func (u *User) FetchById() error {
	db := config.GetDatabaseConnection()

	if err := db.Where("id = ?", u.ID).Find(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}

func (u *User) Save() error {
	db := config.GetDatabaseConnection()

	if db.NewRecord(u) {
		if err := db.Create(&u).Error; err != nil {
			return errors.New("Could not create user")
		}
	} else {
		if err := db.Save(&u).Error; err != nil {
			return errors.New("Could not update user")
		}
	}

	return nil
}

func (u *User) Delete() error {
	db := config.GetDatabaseConnection()

	if err := db.Delete(&u).Error; err != nil {
		return errors.New("Could not find the user")
	}

	return nil
}

func (u *User) ConvertToMap() map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(u)
	json.Unmarshal(inrec, &inInterface)

	return inInterface
}
