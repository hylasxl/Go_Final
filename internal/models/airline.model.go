package models

import (
	"gorm.io/gorm"
)

type Airline struct {
	gorm.Model
	AirlineName string   `gorm:"type:varchar(100)"`
	AirlineCode string   `gorm:"type:varchar(10);unique"`
	Country     string   `gorm:"type:varchar(50)"`
	Flights     []Flight `gorm:"foreignKey:AirlineID"`
}
