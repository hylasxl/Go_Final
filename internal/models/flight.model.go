package models

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	gorm.Model
	FlightNumber         string `gorm:"type:varchar(10);unique"`
	AirlineID            uint   `gorm:"index"`
	DepartureAirportID   uint   `gorm:"index"`
	ArrivalAirportID     uint   `gorm:"index"`
	DepartureTime        time.Time
	ArrivalTime          time.Time
	Duration             time.Duration
	Status               string `gorm:"type:enum('Scheduled','Delayed','Canceled');default:'Scheduled'"`
	AvailableSeats       int
	UpdatedDepartureTime time.Time     `default:"null"`
	UpdatedArrivalTime   time.Time     `default:"null"`
	RescheduleReason     string        `gorm:"type:text" default:""`
	Airline              Airline       `gorm:"foreignKey:AirlineID"`
	DepartureAirport     Airport       `gorm:"foreignKey:DepartureAirportID"`
	ArrivalAirport       Airport       `gorm:"foreignKey:ArrivalAirportID"`
	Reservations         []Reservation `gorm:"foreignKey:FlightID"`
}
