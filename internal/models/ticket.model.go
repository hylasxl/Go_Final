package models

import (
	"gorm.io/gorm"
	"time"
)

type Ticket struct {
	gorm.Model
	ReservationID uint    `gorm:"index"`
	SeatID        uint    `gorm:"index"`
	FlightID      uint    `gorm:"index"`
	TicketNumber  string  `gorm:"type:varchar(20);unique"`
	Price         float64 `gorm:"type:decimal(10,2)"`
	IssueDate     time.Time
	PassengerID   uint        `gorm:"index"`
	Reservation   Reservation `gorm:"foreignKey:ReservationID"`
	Seat          Seat        `gorm:"foreignKey:SeatID"`
	FLight        Flight      `gorm:"foreignKey:FlightID"`
}
