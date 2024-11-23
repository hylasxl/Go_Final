package models

import (
	"gorm.io/gorm"
)

type Airport struct {
	gorm.Model
	AirportCode string   `gorm:"type:varchar(10);unique"`
	AirportName string   `gorm:"type:varchar(100)"`
	City        string   `gorm:"type:varchar(50)"`
	Country     string   `gorm:"type:varchar(50)"`
	Departures  []Flight `gorm:"foreignKey:DepartureAirportID"`
	Arrivals    []Flight `gorm:"foreignKey:ArrivalAirportID"`
}
