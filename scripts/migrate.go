package scripts

import (
	"log"

	"Go_gRPC/config"
	"Go_gRPC/internal/db"
	"Go_gRPC/internal/models" // Import your models
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Automatically migrate your models
	if err := database.AutoMigrate(
		&models.User{},
		&models.Role{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully!")
}
