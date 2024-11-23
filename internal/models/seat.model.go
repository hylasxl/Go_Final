package models

import (
	"gorm.io/gorm"
)

type Seat struct {
	gorm.Model
	FlightID    uint   `gorm:"index"`
	SeatNumber  string `gorm:"type:varchar(10)"`
	SeatClass   string `gorm:"type:enum('Economy','Business','First')"`
	IsAvailable bool   `gorm:"default:true"`
	Price       float64
}

func (seat *Seat) BeforeCreate(tx *gorm.DB) (err error) {
	// Convert bool to smallint (1 or 0)
	if seat.IsAvailable {
		seat.IsAvailable = true // True is stored as 1
	} else {
		seat.IsAvailable = false // False is stored as 0
	}
	return nil
}
