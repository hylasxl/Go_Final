package services

import (
	"Go_gRPC/internal/models"
	"Go_gRPC/pb/airportpb"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"strconv"
)

type AirportService struct {
	airportpb.UnimplementedAirportServiceServer
	DB    *gorm.DB
	Redis *redis.Client
}

func NewAirportService(db *gorm.DB, redisClient *redis.Client) *AirportService {
	return &AirportService{
		DB:    db,
		Redis: redisClient,
	}
}

func (svc *AirportService) GetAirport(ctx context.Context, req *airportpb.GetAirportRequest) (*airportpb.GetAirportResponse, error) {
	cacheKey := "airport:" + req.AirportId

	// Check the cache first
	cachedAirport, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var airport airportpb.Airport
		if err := json.Unmarshal([]byte(cachedAirport), &airport); err == nil {
			return &airportpb.GetAirportResponse{Airport: &airport}, nil
		}
	}

	// Cache miss, query the database
	var airport airportpb.Airport
	if err := svc.DB.Where("id = ?", req.AirportId).First(&airport).Error; err != nil {
		return nil, err
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(airport)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	airportRes := &airportpb.Airport{
		AirportCode: airport.AirportCode,
		AirportName: airport.AirportName,
		Country:     airport.Country,
		City:        airport.City,
		Id:          req.AirportId,
	}
	res := &airportpb.GetAirportResponse{
		Airport: airportRes,
	}

	return res, nil
}

func (svc *AirportService) GetListAirports(ctx context.Context, req *airportpb.GetListAirportRequest) (*airportpb.GetListAirportResponse, error) {
	cacheKey := "airports_list"

	// Check the cache first
	cachedAirports, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var airportRes []*airportpb.Airport
		if err := json.Unmarshal([]byte(cachedAirports), &airportRes); err == nil {
			return &airportpb.GetListAirportResponse{Airports: airportRes}, nil
		}
	}

	// Cache miss, query the database
	var airports []airportpb.Airport
	if err := svc.DB.Find(&airports).Error; err != nil {
		return nil, err
	}

	var airportRes []*airportpb.Airport
	for _, airport := range airports {
		airportRes = append(airportRes, &airportpb.Airport{
			AirportCode: airport.AirportCode,
			AirportName: airport.AirportName,
			Country:     airport.Country,
			City:        airport.City,
			Id:          airport.Id,
		})
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(airportRes)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	return &airportpb.GetListAirportResponse{
		Airports: airportRes,
	}, nil
}

func (svc *AirportService) CreateAirport(ctx context.Context, req *airportpb.CreateAirportRequest) (*airportpb.AirportResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	var existingAirport airportpb.Airport
	if err := tx.Where("airport_code = ?", req.AirportCode).Or("airport_name = ?", req.AirportName).First(&existingAirport).Error; err == nil {
		return nil, fmt.Errorf("airport with code '%s' or name '%s' already exists", req.AirportCode, req.AirportName)
	}

	var maxId uint
	tx.Model(&airportpb.Airport{}).Pluck("id", &maxId)

	newAirport := &airportpb.Airport{
		AirportCode: req.AirportCode,
		AirportName: req.AirportName,
		Country:     req.Country,
		City:        req.City,
		Id:          strconv.Itoa(int(maxId + 1)),
	}

	if err := tx.Create(newAirport).Error; err != nil {
		return nil, fmt.Errorf("failed to create airport: %v", err)
	}

	tx.Commit()

	// Invalidate the airports list cache
	svc.Redis.Del(ctx, "airports_list")

	return &airportpb.AirportResponse{}, nil
}

func (svc *AirportService) UpdateAirport(ctx context.Context, req *airportpb.UpdateAirportRequest) (*airportpb.AirportResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	// Check if the airport exists
	var existingAirport airportpb.Airport
	if err := tx.Where("id = ?", req.AirportId).First(&existingAirport).Error; err != nil {
		return &airportpb.AirportResponse{
			Error: "Airport with id " + req.AirportId + " does not exist",
		}, nil
	}

	updatedAirport := &airportpb.Airport{
		AirportCode: req.AirportCode,
		AirportName: req.AirportName,
		Country:     req.Country,
		City:        req.City,
	}

	if err := tx.Model(&existingAirport).Updates(updatedAirport).Error; err != nil {
		return &airportpb.AirportResponse{
			Error: "Failed to update airport",
		}, nil
	}

	tx.Commit()

	// Invalidate the cache for the updated airport
	cacheKey := "airport:" + req.AirportId
	svc.Redis.Del(ctx, cacheKey)

	// Invalidate the airports list cache
	svc.Redis.Del(ctx, "airports_list")

	return &airportpb.AirportResponse{}, nil
}

func (svc *AirportService) DeleteAirport(ctx context.Context, req *airportpb.DeleteAirportRequest) (*airportpb.AirportResponse, error) {
	tx := svc.DB.Begin()
	defer tx.Rollback() // Ensure rollback on error

	// Check if the airport exists
	var existingAirport airportpb.Airport
	if err := tx.Where("id = ?", req.AirportId).First(&existingAirport).Error; err != nil {
		return &airportpb.AirportResponse{
			Error: "Airport with id " + req.AirportId + " does not exist",
		}, nil
	}

	if err := tx.Exec("UPDATE airports SET deleted_at = NOW() WHERE id = ?", req.AirportId).Error; err != nil {
		return &airportpb.AirportResponse{
			Error: "Delete failed",
		}, nil
	}

	tx.Commit()

	// Invalidate the cache for the deleted airport
	cacheKey := "airport:" + req.AirportId
	svc.Redis.Del(ctx, cacheKey)

	// Invalidate the airports list cache
	svc.Redis.Del(ctx, "airports_list")

	return &airportpb.AirportResponse{}, nil
}

func (svc *AirportService) AirportGetDepartureFlights(ctx context.Context, req *airportpb.AirportGetDepartureFlightRequest) (*airportpb.AirportGetDepartureFlightResponse, error) {
	cacheKey := "airport_departure_flights:" + strconv.Itoa(int(req.AirportId))

	// Check the cache first
	cachedFlights, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var flightResponses []*airportpb.Flight
		if err := json.Unmarshal([]byte(cachedFlights), &flightResponses); err == nil {
			return &airportpb.AirportGetDepartureFlightResponse{Flights: flightResponses}, nil
		}
	}

	// Cache miss, fetch the flights from the database
	var flights []models.Flight
	if err := svc.DB.Where("departure_airport_id = ?", req.AirportId).Find(&flights).Error; err != nil {
		return &airportpb.AirportGetDepartureFlightResponse{
			Error: fmt.Sprintf("Failed to retrieve airport flights: %v", err),
		}, nil
	}

	var flightResponses []*airportpb.Flight
	for _, flight := range flights {
		flightResponses = append(flightResponses, &airportpb.Flight{
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

	return &airportpb.AirportGetDepartureFlightResponse{
		Flights: flightResponses,
	}, nil
}

func (svc *AirportService) AirportGetArrivalFlights(ctx context.Context, req *airportpb.AirportGetArrivalFlightRequest) (*airportpb.AirportGetArrivalFlightResponse, error) {
	cacheKey := "airport_arrival_flights:" + strconv.Itoa(int(req.AirportId))

	// Check the cache first
	cachedFlights, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var flightResponses []*airportpb.Flight
		if err := json.Unmarshal([]byte(cachedFlights), &flightResponses); err == nil {
			return &airportpb.AirportGetArrivalFlightResponse{Flights: flightResponses}, nil
		}
	}

	// Cache miss, fetch the flights from the database
	var flights []models.Flight
	if err := svc.DB.Where("arrival_airport_id = ?", req.AirportId).Find(&flights).Error; err != nil {
		return &airportpb.AirportGetArrivalFlightResponse{
			Error: fmt.Sprintf("Failed to retrieve airport flights: %v", err),
		}, nil
	}

	var flightResponses []*airportpb.Flight
	for _, flight := range flights {
		flightResponses = append(flightResponses, &airportpb.Flight{
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

	return &airportpb.AirportGetArrivalFlightResponse{
		Flights: flightResponses,
	}, nil
}

func (svc *AirportService) AirportGetDepartureFlightsAndArrivalFlights(ctx context.Context, req *airportpb.AirportGetDepartureFlightsAndArrivalFlightRequest) (*airportpb.AirportGetDepartureFlightsAndArrivalFlightResponse, error) {
	cacheKey := "airport_departure_arrival_flights:" + strconv.Itoa(int(req.DepartureAirportId)) + ":" + strconv.Itoa(int(req.ArrivalAirportId))

	// Check the cache first
	cachedFlights, err := svc.Redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var flightResponses []*airportpb.Flight
		if err := json.Unmarshal([]byte(cachedFlights), &flightResponses); err == nil {
			return &airportpb.AirportGetDepartureFlightsAndArrivalFlightResponse{Flights: flightResponses}, nil
		}
	}

	// Cache miss, fetch the flights from the database
	var flights []models.Flight
	if err := svc.DB.Where("departure_airport_id = ? AND arrival_airport_id = ?", req.DepartureAirportId, req.ArrivalAirportId).Order("departure_time DESC").Find(&flights).Error; err != nil {
		return &airportpb.AirportGetDepartureFlightsAndArrivalFlightResponse{
			Error: fmt.Sprintf("Failed to retrieve airport flights: %v", err),
		}, nil
	}

	var flightResponses []*airportpb.Flight
	for _, flight := range flights {
		var airline models.Airline
		if err := svc.DB.Where("id = ?", flight.AirlineID).First(&airline).Error; err != nil {
			return &airportpb.AirportGetDepartureFlightsAndArrivalFlightResponse{
				Error: fmt.Sprintf("Failed to retrieve airline: %v", err),
			}, nil
		}

		flightResponses = append(flightResponses, &airportpb.Flight{
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
			Airline:              airline.AirlineName,
			Id:                   strconv.Itoa(int(flight.ID)),
		})
	}

	// Cache the result in Redis
	cachedData, _ := json.Marshal(flightResponses)
	svc.Redis.Set(ctx, cacheKey, cachedData, 0) // Set with no expiration

	return &airportpb.AirportGetDepartureFlightsAndArrivalFlightResponse{
		Flights: flightResponses,
	}, nil
}
