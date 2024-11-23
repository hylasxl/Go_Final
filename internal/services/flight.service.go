package services

import (
	"Go_gRPC/internal/models"
	"Go_gRPC/pb/flightpb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

type FlightService struct {
	flightpb.UnimplementedFlightServiceServer
	DB    *gorm.DB
	Redis *redis.Client
}

func NewFlightService(db *gorm.DB) *FlightService {
	return &FlightService{
		DB: db,
	}
}

func (svc *FlightService) CreateFlight(ctx context.Context, req *flightpb.CreateFlightRequest) (*flightpb.CreateFlightResponse, error) {
	if req.AvailableSeats < 0 {
		return &flightpb.CreateFlightResponse{
			Error: "Invalid number of seats",
		}, nil
	}

	if req.ArrivalAirportID == req.DepartureAirportID {
		return &flightpb.CreateFlightResponse{
			Error: "Duplicate airport",
		}, nil
	}

	departureTime := req.DepartureTime.AsTime()
	arrivalTime := req.ArrivalTime.AsTime()
	currentTime := time.Now()
	duration := arrivalTime.Sub(departureTime)
	durationInMinutes := int(duration.Minutes())

	if departureTime.Before(currentTime) {
		return &flightpb.CreateFlightResponse{
			Error: "Departure time cannot be in the past",
		}, nil
	}

	if arrivalTime.Before(currentTime) {
		return &flightpb.CreateFlightResponse{
			Error: "Arrival time cannot be in the past",
		}, nil
	}

	if arrivalTime.Before(departureTime) {
		return &flightpb.CreateFlightResponse{
			Error: "Arrival time cannot be before departure time",
		}, nil
	}

	tx := svc.DB.Begin()

	var departureAirport models.Airport
	if err := tx.First(&departureAirport, "id = ?", req.DepartureAirportID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.CreateFlightResponse{
				Error: "Departure airport not found",
			}, status.Errorf(codes.NotFound, "Departure airport with ID %d does not exist.", req.DepartureAirportID)
		}
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Failed to query departure airport: %v", err)
	}

	var arrivalAirport models.Airport
	if err := tx.First(&arrivalAirport, "id = ?", req.ArrivalAirportID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.CreateFlightResponse{
				Error: "Arrival airport not found",
			}, status.Errorf(codes.NotFound, "Arrival airport with ID %d does not exist.", req.ArrivalAirportID)
		}
	}

	var airline models.Airline
	if err := tx.First(&airline, "id = ?", req.AirlineID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.CreateFlightResponse{
				Error: "Airline not found",
			}, status.Errorf(codes.NotFound, "Airline with ID %d does not exist.", req.AirlineID)
		}
	}

	var flightNumber string
	for {
		flightNumber = GenerateFlightNumber()

		var existingFlightNumber models.Flight
		if err := tx.First(&existingFlightNumber, "flight_number = ?", flightNumber).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}
			log.Printf("Error checking for existing flight number: %v", err)
			return &flightpb.CreateFlightResponse{
				Error: "Database error",
			}, status.Errorf(codes.Internal, "Database error: %v", err)
		}
		// If a record exists, continue the loop to generate a new flight number
	}

	flight := &models.Flight{
		FlightNumber:         flightNumber,
		AirlineID:            uint(req.AirlineID),
		DepartureAirportID:   uint(req.DepartureAirportID),
		ArrivalAirportID:     uint(req.ArrivalAirportID),
		DepartureTime:        departureTime,
		ArrivalTime:          arrivalTime,
		Duration:             time.Duration(durationInMinutes),
		Status:               "Scheduled",
		AvailableSeats:       int(req.AvailableSeats),
		UpdatedDepartureTime: departureTime,
		UpdatedArrivalTime:   arrivalTime,
		RescheduleReason:     "",
	}

	if err := tx.Create(flight).Error; err != nil {
		tx.Rollback()
		return &flightpb.CreateFlightResponse{
			Error: "Failed to create flight",
		}, status.Errorf(codes.Internal, "Failed to create flight: %v", err)
	}

	if err := svc.createSeatsForFlight(tx, flight.ID, flight.AvailableSeats, durationInMinutes); err != nil {
		tx.Rollback()
		return &flightpb.CreateFlightResponse{
			Error: "Failed to create seats",
		}, status.Errorf(codes.Internal, "Failed to create seats: %v", err)
	}

	departureAirportF := &flightpb.AirportF{
		AirportCode: departureAirport.AirportCode,
		AirportName: departureAirport.AirportName,
		Country:     departureAirport.Country,
		City:        departureAirport.City,
	}

	arrivalAirportF := &flightpb.AirportF{
		AirportCode: arrivalAirport.AirportCode,
		AirportName: arrivalAirport.AirportName,
		Country:     arrivalAirport.Country,
		City:        arrivalAirport.City,
	}

	arlineF := &flightpb.AirlineF{
		AirlineCode: airline.AirlineCode,
		AirlineName: airline.AirlineName,
		Country:     airline.Country,
	}

	res := &flightpb.CreateFlightResponse{
		FlightNumber:     flight.FlightNumber,
		DepartureAirport: departureAirportF,
		ArrivalAirport:   arrivalAirportF,
		DepartureTime:    timestamppb.New(flight.DepartureTime),
		ArrivalTime:      timestamppb.New(flight.ArrivalTime),
		AvailableSeats:   uint32(flight.AvailableSeats),
		Duration:         uint32(flight.Duration),
		Status:           flight.Status,
		Airline:          arlineF,
		Error:            "",
	}

	tx.Commit()

	return res, nil
}

func GenerateFlightNumber() string {
	rand.Seed(time.Now().UnixNano())
	prefix := string([]rune{
		rune('A' + rand.Intn(26)),
		rune('A' + rand.Intn(26)),
	})
	postfix := rand.Intn(10000)
	return fmt.Sprintf("%s%04d", prefix, postfix)
}

func (svc *FlightService) createSeatsForFlight(tx *gorm.DB, flightID uint, availableSeats int, durationInMinutes int) error {

	var baseRate float64
	switch {
	case durationInMinutes < 180:
		baseRate = 0.15
	case durationInMinutes <= 360:
		baseRate = 0.12
	default:
		baseRate = 0.08
	}

	economyMultiplier := 1.0
	businessMultiplier := 3.0
	firstMultiplier := 6.0

	economySeats := int(0.7 * float64(availableSeats))
	businessSeats := int(0.2 * float64(availableSeats))
	firstClassSeats := availableSeats - economySeats - businessSeats // Remainder for first class

	type Seat2 struct {
		gorm.Model
		FlightID    uint   `gorm:"index"`
		SeatNumber  string `gorm:"type:varchar(10)"`
		SeatClass   string `gorm:"type:enum('Economy','Business','First')"`
		IsAvailable int    `gorm:"default:1"`
		Price       float64
	}

	var seats []Seat2

	prefixes := []string{"A", "B", "C", "D"}
	for i := 0; i < economySeats; i++ {
		seats = append(seats, Seat2{
			FlightID:    flightID,
			SeatNumber:  fmt.Sprintf("%s%d", prefixes[i%len(prefixes)], (i/len(prefixes))+1),
			SeatClass:   "Economy",
			Price:       baseRate * economyMultiplier * float64(durationInMinutes),
			IsAvailable: 1,
		})
	}

	for i := 0; i < businessSeats; i++ {
		seats = append(seats, Seat2{
			FlightID:    flightID,
			SeatNumber:  fmt.Sprintf("BS%d", i+1),
			SeatClass:   "Business",
			Price:       baseRate * businessMultiplier * float64(durationInMinutes),
			IsAvailable: 1,
		})
	}

	for i := 0; i < firstClassSeats; i++ {
		seats = append(seats, Seat2{
			FlightID:    flightID,
			SeatNumber:  fmt.Sprintf("FI%d", i+1),
			SeatClass:   "First",
			Price:       baseRate * firstMultiplier * float64(durationInMinutes),
			IsAvailable: 1,
		})
	}

	if err := tx.Model(&models.Seat{}).Create(&seats).Error; err != nil {
		return err
	}

	return nil
}

func (svc *FlightService) RescheduleFlight(ctx context.Context, req *flightpb.RescheduleFlightRequest) (*flightpb.RescheduleFlightResponse, error) {
	tx := svc.DB.Begin()
	var flightStatus = "Scheduled"
	var flight models.Flight
	if err := tx.First(&flight, "id = ?", req.FlightID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.RescheduleFlightResponse{
				Error: "Flight not found",
			}, status.Errorf(codes.NotFound, "Flight with ID %d does not exist.", req.FlightID)
		}
	}
	if flight.Status == "Canceled" {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Flight is canceled",
		}, status.Errorf(codes.FailedPrecondition, "Flight with ID %d is canceled.", req.FlightID)
	}

	if flight.DepartureTime.Before(time.Now()) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Cannot reschedule a flight that has already departed",
		}, status.Errorf(codes.FailedPrecondition, "Flight departure time %v is in the past.", flight.DepartureTime)
	}

	if flight.ArrivalTime.Before(time.Now()) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Cannot reschedule a flight that has already arrived",
		}, status.Errorf(codes.FailedPrecondition, "Flight arrival time %v is in the past.", flight.ArrivalTime)
	}

	if req.UpdatedDepartureTime.AsTime().Before(time.Now()) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Updated departure time is in the past",
		}, status.Errorf(codes.FailedPrecondition, "Updated departure time %v is in the past.", req.UpdatedDepartureTime)
	}

	if req.UpdatedArrivalTime.AsTime().Before(time.Now()) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Updated arrival time is in the past",
		}, status.Errorf(codes.FailedPrecondition, "Updated arrival time %v is in the past.", req.UpdatedArrivalTime)
	}

	if req.UpdatedDepartureTime.AsTime().Equal(flight.DepartureTime) && req.UpdatedArrivalTime.AsTime().Equal(flight.ArrivalTime) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Updated departure and arrival times are the same",
		}, status.Errorf(codes.FailedPrecondition, "Updated departure and arrival times are the same.")
	}

	if req.UpdatedDepartureTime.AsTime().After(req.UpdatedArrivalTime.AsTime()) {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Updated departure time is after updated arrival time",
		}, status.Errorf(codes.FailedPrecondition, "Updated departure time %v is after updated arrival time %v.", req.UpdatedDepartureTime, req.UpdatedArrivalTime)
	}

	var departureAirport models.Airport
	if err := tx.First(&departureAirport, "id = ?", flight.DepartureAirportID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.RescheduleFlightResponse{
				Error: "Departure airport not found",
			}, status.Errorf(codes.NotFound, "Departure airport with ID %d does not exist.", flight.DepartureAirportID)
		}
	}

	var arrivalAirport models.Airport
	if err := tx.First(&arrivalAirport, "id = ?", flight.ArrivalAirportID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.RescheduleFlightResponse{
				Error: "Arrival airport not found",
			}, status.Errorf(codes.NotFound, "Arrival airport with ID %d does not exist.", flight.ArrivalAirportID)
		}
	}

	var airline models.Airline
	if err := tx.First(&airline, "id = ?", flight.AirlineID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.RescheduleFlightResponse{
				Error: "Airline not found",
			}, status.Errorf(codes.NotFound, "Airline with ID %d does not exist.", flight.AirlineID)
		}
	}

	if req.UpdatedDepartureTime.AsTime().After(flight.DepartureTime) {
		flightStatus = "Delayed"
	}

	updatedDuration := req.UpdatedArrivalTime.AsTime().Sub(req.UpdatedDepartureTime.AsTime())

	flight.UpdatedDepartureTime = req.UpdatedDepartureTime.AsTime()
	flight.UpdatedArrivalTime = req.UpdatedArrivalTime.AsTime()
	flight.RescheduleReason = req.Reason
	flight.Status = flightStatus
	flight.Duration = time.Duration(int(updatedDuration.Minutes()))

	if err := tx.Save(&flight).Error; err != nil {
		tx.Rollback()
		return &flightpb.RescheduleFlightResponse{
			Error: "Error rescheduling flight",
		}, status.Errorf(codes.Internal, "Error rescheduling flight: %v", err)
	}

	departureAirportF := &flightpb.AirportF{
		AirportCode: departureAirport.AirportCode,
		AirportName: departureAirport.AirportName,
		Country:     departureAirport.Country,
		City:        departureAirport.City,
	}

	arrivalAirportF := &flightpb.AirportF{
		AirportCode: arrivalAirport.AirportCode,
		AirportName: arrivalAirport.AirportName,
		Country:     arrivalAirport.Country,
		City:        arrivalAirport.City,
	}

	airlineF := &flightpb.AirlineF{
		AirlineCode: airline.AirlineCode,
		AirlineName: airline.AirlineName,
		Country:     airline.Country,
	}

	tx.Commit()

	res := &flightpb.RescheduleFlightResponse{
		FlightNumber:         flight.FlightNumber,
		DepartureAirport:     departureAirportF,
		ArrivalAirport:       arrivalAirportF,
		Airline:              airlineF,
		DepartureTime:        timestamppb.New(flight.DepartureTime),
		ArrivalTime:          timestamppb.New(flight.ArrivalTime),
		UpdatedDepartureTime: timestamppb.New(flight.UpdatedDepartureTime),
		UpdatedArrivalTime:   timestamppb.New(flight.UpdatedArrivalTime),
		AvailableSeats:       uint32(flight.AvailableSeats),
		Reason:               flight.RescheduleReason,
		Status:               flight.Status,
		Duration:             uint32(flight.Duration),
	}
	return res, nil
}

func (svc *FlightService) GetFlightSeats(ctx context.Context, req *flightpb.GetFlightSeatRequest) (*flightpb.GetFlightSeatResponse, error) {

	tx := svc.DB.Begin()

	var flight models.Flight
	if err := tx.First(&flight, "id = ?", req.FlightID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return &flightpb.GetFlightSeatResponse{
				Error: fmt.Sprintf("Flight with ID %d not found", req.FlightID),
			}, status.Errorf(codes.NotFound, "Flight with ID %d does not exist.", req.FlightID)
		}
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "Error retrieving flight: %v", err)
	}

	var seats []models.Seat
	if err := tx.Find(&seats, "flight_id = ?", req.FlightID).Error; err != nil {
		tx.Rollback()
		return &flightpb.GetFlightSeatResponse{
			Error: "Error getting flight seats",
		}, status.Errorf(codes.Internal, "Error getting flight seats: %v", err)
	}

	var seatsF []*flightpb.SeatF
	for _, seat := range seats {
		seatsF = append(seatsF, &flightpb.SeatF{
			SeatNumber:  seat.SeatNumber,
			SeatClass:   seat.SeatClass,
			IsAvailable: seat.IsAvailable,
			Price:       seat.Price,
			ID:          uint32(seat.ID),
		})
	}

	tx.Commit()

	res := &flightpb.GetFlightSeatResponse{
		Seats: seatsF,
	}
	return res, nil
}

func (svc *FlightService) GetFlights(ctx context.Context, req *flightpb.GetFlightRequest) (*flightpb.GetFlightResponse, error) {
	// Define a cache key for the flights
	cacheKey := "flights_list"

	// Check the cache first
	cachedFlights, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		// Cache hit, return the cached response
		var flightResponses []*flightpb.RescheduleFlightResponse
		if err := json.Unmarshal([]byte(cachedFlights), &flightResponses); err == nil {
			return &flightpb.GetFlightResponse{Flights: flightResponses}, nil
		}
	}

	// Cache miss, query the database
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	var flights []models.Flight
	if err := tx.Order("created_at DESC").Find(&flights).Error; err != nil {
		return &flightpb.GetFlightResponse{
			Error: "Error getting flights",
		}, status.Errorf(codes.Internal, "Error getting flights: %v", err)
	}

	// Preload related entities for optimization
	var departureAirports, arrivalAirports []models.Airport
	var airlines []models.Airline

	depIDs := uniqueFlightIDs(flights, func(f models.Flight) uint32 { return uint32(f.DepartureAirportID) })
	arrIDs := uniqueFlightIDs(flights, func(f models.Flight) uint32 { return uint32(f.ArrivalAirportID) })
	airlineIDs := uniqueFlightIDs(flights, func(f models.Flight) uint32 { return uint32(f.AirlineID) })

	if err := tx.Where("id IN ?", depIDs).Find(&departureAirports).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error loading departure airports: %v", err)
	}
	if err := tx.Where("id IN ?", arrIDs).Find(&arrivalAirports).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error loading arrival airports: %v", err)
	}
	if err := tx.Where("id IN ?", airlineIDs).Find(&airlines).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error loading airlines: %v", err)
	}

	// Build lookup maps for efficient access
	depMap := airportMap(departureAirports)
	arrMap := airportMap(arrivalAirports)
	airlineMap := airlineMap(airlines)

	var flightResponses []*flightpb.RescheduleFlightResponse

	for _, flight := range flights {
		departureAirport := depMap[uint32(flight.DepartureAirportID)]
		arrivalAirport := arrMap[uint32(flight.ArrivalAirportID)]
		airline := airlineMap[uint32(flight.AirlineID)]

		if departureAirport == nil || arrivalAirport == nil || airline == nil {
			return nil, status.Errorf(codes.Internal, "Inconsistent data for flight ID %d", flight.ID)
		}

		flightResponse := &flightpb.RescheduleFlightResponse{
			FlightNumber: flight.FlightNumber,
			DepartureAirport: &flightpb.AirportF{
				AirportCode: departureAirport.AirportCode,
				AirportName: departureAirport.AirportName,
				Country:     departureAirport.Country,
				City:        departureAirport.City,
			},
			ArrivalAirport: &flightpb.AirportF{
				AirportCode: arrivalAirport.AirportCode,
				AirportName: arrivalAirport.AirportName,
				Country:     arrivalAirport.Country,
				City:        arrivalAirport.City,
			},
			Airline: &flightpb.AirlineF{
				AirlineCode: airline.AirlineCode,
				AirlineName: airline.AirlineName,
				Country:     airline.Country,
			},
			DepartureTime: timestamppb.New(flight.DepartureTime),
			ArrivalTime:   timestamppb.New(flight.ArrivalTime),
			Duration:      uint32(flight.Duration),
			Status:        flight.Status,
			ID:            uint32(flight.ID),
		}
		flightResponses = append(flightResponses, flightResponse)
	}

	tx.Commit()

	// Cache the result in Redis
	cachedData, _ := json.Marshal(flightResponses)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	return &flightpb.GetFlightResponse{Flights: flightResponses}, nil
}

func uniqueFlightIDs(flights []models.Flight, keyFunc func(f models.Flight) uint32) []uint32 {
	idSet := make(map[uint32]struct{})
	for _, flight := range flights {
		idSet[keyFunc(flight)] = struct{}{}
	}

	ids := make([]uint32, 0, len(idSet))
	for id := range idSet {
		ids = append(ids, id)
	}
	return ids
}

func airportMap(airports []models.Airport) map[uint32]*models.Airport {
	airportMap := make(map[uint32]*models.Airport)
	for i := range airports {
		airportMap[uint32(airports[i].ID)] = &airports[i]
	}
	return airportMap
}

func airlineMap(airlines []models.Airline) map[uint32]*models.Airline {
	airlineMap := make(map[uint32]*models.Airline)
	for i := range airlines {
		airlineMap[uint32(airlines[i].ID)] = &airlines[i]
	}
	return airlineMap
}
