package db

import (
	models2 "Go_gRPC/internal/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models2.User{},
		&models2.Airline{},
		&models2.Airport{},
		&models2.Flight{},
		&models2.Reservation{},
		&models2.Ticket{},
		&models2.Seat{},
	); err != nil {
		return err
	}
	print("Migrate successfully\n")
	return nil
}
