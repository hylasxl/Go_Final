package models

import (
	"gorm.io/gorm"
	"time"
)

type Reservation struct {
	gorm.Model
	UserID            uint `gorm:"index"`
	FlightID          uint `gorm:"index"`
	ReservationDate   time.Time
	ReservationStatus string `gorm:"type:enum('Confirmed','Canceled','Pending')"`
	PassportID        int
	TotalPrice        float64 `gorm:"type:decimal(10,2)"`
	User              User    `gorm:"foreignKey:UserID"`
	Flight            Flight  `gorm:"foreignKey:FlightID"`
}
