package services

import (
	"Go_gRPC/internal/models"
	"Go_gRPC/pb/airlinepb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"strconv"
)

type AirlineService struct {
	airlinepb.UnimplementedAirlineServiceServer
	DB    *gorm.DB
	Redis *redis.Client
}

func NewAirlineService(db *gorm.DB, redisClient *redis.Client) *AirlineService {
	return &AirlineService{
		DB:    db,
		Redis: redisClient,
	}
}

func (svc *AirlineService) CreateAirline(ctx context.Context, req *airlinepb.CreateArlineRequest) (*airlinepb.AirlineResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	// Check for existing airline with the same code or name
	var existingAirline airlinepb.Airline
	if err := tx.Where("airline_code = ?", req.AirlineCode).Or("airline_name = ?", req.AirlineName).First(&existingAirline).Error; err == nil {
		return &airlinepb.AirlineResponse{
			Error: fmt.Sprintf("Airline with code '%s' or name '%s' already exists", req.AirlineCode, req.AirlineName),
		}, nil
	}

	var maxId uint
	tx.Model(&airlinepb.Airline{}).Pluck("id", &maxId)

	newAirline := &airlinepb.Airline{
		AirlineName: req.AirlineName,
		AirlineCode: req.AirlineCode,
		Country:     req.Country,
		Id:          strconv.Itoa(int(maxId + 1)),
	}

	// Create the airline
	if err := tx.Create(newAirline).Error; err != nil {
		return &airlinepb.AirlineResponse{
			Error: fmt.Sprintf("Failed to create airline: %v", err),
		}, nil
	}

	tx.Commit()
	return &airlinepb.AirlineResponse{}, nil
}

func (svc *AirlineService) GetAirline(ctx context.Context, req *airlinepb.GetAirlineRequest) (*airlinepb.GetAirlineResponse, error) {
	cacheKey := "airline:" + req.AirlineId

	// Check the cache first
	cachedAirline, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var airline airlinepb.Airline
		if err := json.Unmarshal([]byte(cachedAirline), &airline); err == nil {
			return &airlinepb.GetAirlineResponse{Airline: &airline}, nil
		}
	}

	// Cache miss, query the database
	var airline airlinepb.Airline
	if err := svc.DB.Where("id = ?", req.AirlineId).First(&airline).Error; err != nil {
		return nil, fmt.Errorf("airline with ID '%s' not found", req.AirlineId)
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(airline)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	res := &airlinepb.GetAirlineResponse{
		Airline: &airlinepb.Airline{
			AirlineName: airline.AirlineName,
			AirlineCode: airline.AirlineCode,
			Country:     airline.Country,
			Id:          req.AirlineId,
		},
	}

	return res, nil
}

func (svc *AirlineService) GetListAirline(ctx context.Context, req *airlinepb.GetListAirlineRequest) (*airlinepb.GetListAirlineResponse, error) {
	cacheKey := "airlines_list"

	// Check the cache first
	cachedAirlines, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var airlineRes []*airlinepb.Airline
		if err := json.Unmarshal([]byte(cachedAirlines), &airlineRes); err == nil {
			return &airlinepb.GetListAirlineResponse{Airlines: airlineRes}, nil
		}
	}

	// Cache miss, query the database
	var airlines []airlinepb.Airline
	if err := svc.DB.Find(&airlines).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve airline list: %v", err)
	}

	var airlineRes []*airlinepb.Airline
	for _, airline := range airlines {
		airlineRes = append(airlineRes, &airlinepb.Airline{
			AirlineName: airline.AirlineName,
			AirlineCode: airline.AirlineCode,
			Country:     airline.Country,
			Id:          airline.Id,
		})
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(airlineRes)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	return &airlinepb.GetListAirlineResponse{
		Airlines: airlineRes,
	}, nil
}

func (svc *AirlineService) UpdateAirline(ctx context.Context, req *airlinepb.UpdateAirlineRequest) (*airlinepb.AirlineResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	// Find the airline to update
	var existingAirline airlinepb.Airline
	if err := tx.Where("id = ?", req.AirlineId).First(&existingAirline).Error; err != nil {
		return &airlinepb.AirlineResponse{
			Error: fmt.Sprintf("Airline with ID '%s' not found", req.AirlineId),
		}, nil
	}

	// Update airline details
	if err := tx.Model(&existingAirline).Updates(&airlinepb.Airline{
		AirlineName: req.AirlineName,
		AirlineCode: req.AirlineCode,
		Country:     req.Country,
	}).Error; err != nil {
		return &airlinepb.AirlineResponse{
			Error: "Failed to update airline",
		}, nil
	}

	tx.Commit()

	// Invalidate the cache for the updated airline
	cacheKey := "airline:" + req.AirlineId
	svc.Redis.Del(ctx, cacheKey)

	return &airlinepb.AirlineResponse{}, nil
}

func (svc *AirlineService) DeleteAirline(ctx context.Context, req *airlinepb.DeleteAirlineRequest) (*airlinepb.AirlineResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	// Check if the airline exists
	var existingAirline airlinepb.Airline
	if err := tx.Where("id = ?", req.AirlineId).First(&existingAirline).Error; err != nil {
		return &airlinepb.AirlineResponse{
			Error: fmt.Sprintf("Airline with ID '%s' not found", req.AirlineId),
		}, nil
	}

	// Mark the airline as deleted by setting a deleted_at timestamp
	if err := tx.Exec("UPDATE airlines SET deleted_at = NOW() WHERE id = ?", req.AirlineId).Error; err != nil {
		return &airlinepb.AirlineResponse{
			Error: "Failed to delete airline",
		}, nil
	}

	tx.Commit()

	// Invalidate the cache for the deleted airline
	cacheKey := "airline:" + req.AirlineId
	svc.Redis.Del(ctx, cacheKey)

	return &airlinepb.AirlineResponse{}, nil
}

func (svc *AirlineService) AirlineGetFlights(ctx context.Context, request *airlinepb.AirlineGetFlightRequest) (*airlinepb.AirlineGetFlightResponse, error) {
	cacheKey := "airline_flights:" + request.AirlineId

	// Check the cache first
	cachedFlights, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var flightResponses []*airlinepb.FlightAL
		if err := json.Unmarshal([]byte(cachedFlights), &flightResponses); err == nil {
			return &airlinepb.AirlineGetFlightResponse{Flights: flightResponses}, nil
		}
	}

	// Cache miss, fetch the flights from the database
	var flights []models.Flight
	if err := svc.DB.Order("created_at DESC").Where("airline_id = ?", request.AirlineId).Find(&flights).Error; err != nil {
		return &airlinepb.AirlineGetFlightResponse{
			Error: fmt.Sprintf("Failed to retrieve airline flights: %v", err),
		}, nil
	}

	var flightResponses []*airlinepb.FlightAL
	for _, flight := range flights {
		flightResponses = append(flightResponses, &airlinepb.FlightAL{
			FlightNumber:         flight.FlightNumber,
			DepartureAirportID:   uint32(flight.DepartureAirportID),
			ArrivalAirportID:     uint32(flight.ArrivalAirportID),
			DepartureTime:        timestamppb.New(flight.DepartureTime),
			ArrivalTime:          timestamppb.New(flight.ArrivalTime),
			AvailableSeats:       uint32(flight.AvailableSeats),
			Duration:             uint32(flight.Duration),
			Status:               flight.Status,
			UpdatedDepartureTime: timestamppb.New(flight.UpdatedDepartureTime),
			UpdatedArrivalTime:   timestamppb.New(flight.UpdatedArrivalTime),
			Reason:               flight.RescheduleReason,
		})
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(flightResponses)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	return &airlinepb.AirlineGetFlightResponse{
		Flights: flightResponses,
	}, nil
}
