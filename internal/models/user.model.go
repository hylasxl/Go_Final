package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(50);unique"`
	Password     string `gorm:"type:varchar(255)"`
	FirstName    string `gorm:"type:varchar(50)"`
	LastName     string `gorm:"type:varchar(50)"`
	Email        string `gorm:"type:varchar(100);unique"`
	PhoneNumber  string `gorm:"type:varchar(15);unique"`
	DateOfBirth  time.Time
	Address      string        `gorm:"type:text"`
	RoleID       uint          `gorm:"not null"`
	Reservations []Reservation `gorm:"foreignKey:UserID"`
}
