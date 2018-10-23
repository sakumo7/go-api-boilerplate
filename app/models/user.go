package models

import (
	"time"
	"../config"
)

type User struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt *time.Time `json:"createdAt, omitempty"`
	UpdatedAt *time.Time `json:"updatedAt, omitempty"`
	DeletedAt *time.Time `json:"deletedAt, omitempty" sql:"index"`
	FirstName string `json:"firstname, omitempty" gorm:"not null; type:varchar(100)"`
	LastName  string `json:"lastname, omitempty" gorm:"not null; type:varchar(100)"`
	Email     string `json:"email, omitempty" gorm:"not null; type:varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) FetchAll() []User {
	db := config.GetDatabaseConnection()
	var users []User
	db.Find(&users)

	return users
}