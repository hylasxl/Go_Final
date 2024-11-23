package services

import (
	"Go_gRPC/internal/models"
	"Go_gRPC/pb/reservationpb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type ReservationService struct {
	reservationpb.UnimplementedReservationServiceServer
	DB *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{DB: db}
}

func (svc *ReservationService) BookFlight(ctx context.Context, req *reservationpb.BookFlightRequest) (*reservationpb.BookFlightResponse, error) {
	var seat models.Seat
	tx := svc.DB.Begin()
	if err := tx.Model(&models.Seat{}).Where("id = ?", req.SeatId).First(&seat).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if seat.IsAvailable == false {
		tx.Rollback()
		return &reservationpb.BookFlightResponse{
			Error: "Seat is not available",
		}, nil
	}

	if err := tx.Model(&models.Seat{}).Where("id = ?", req.SeatId).Update("is_available", 0).Error; err != nil {
		tx.Rollback()
		println(err)
		return &reservationpb.BookFlightResponse{
			Error: "Error while updating seat",
		}, nil
	}

	reservation := &models.Reservation{
		UserID:            uint(req.UserId),
		FlightID:          uint(req.FlightId),
		ReservationDate:   time.Now(),
		ReservationStatus: "Confirmed",
		PassportID:        int(req.PassportId),
		TotalPrice:        seat.Price,
	}

	if err := tx.Create(reservation).Error; err != nil {
		tx.Rollback()
		return &reservationpb.BookFlightResponse{
			Error: "Error while creating reservation",
		}, nil
	}

	ticket := &models.Ticket{
		ReservationID: reservation.ID,
		SeatID:        seat.ID,
		TicketNumber:  GenerateRandomTicketNumber(20),
		Price:         float64(seat.Price),
		IssueDate:     time.Now(),
		PassengerID:   uint(req.UserId),
		FlightID:      uint(req.FlightId),
	}

	if err := tx.Create(ticket).Error; err != nil {
		tx.Rollback()
		return &reservationpb.BookFlightResponse{
			Error: "Error while creating ticket",
		}, nil
	}

	var user models.User
	if err := tx.Model(&models.User{}).Where("id = ?", req.UserId).First(&user).Error; err != nil {
		tx.Rollback()
		return &reservationpb.BookFlightResponse{
			Error: "Error while getting user",
		}, nil
	}

	if err := tx.Commit().Error; err != nil {
	}

	res := &reservationpb.BookFlightResponse{
		TicketNumber:         ticket.TicketNumber,
		ReservationId:        uint32(reservation.ID),
		SeatNumber:           seat.SeatNumber,
		Price:                uint32(seat.Price),
		IssueDate:            timestamppb.New(ticket.IssueDate),
		PassengerName:        user.FirstName + " " + user.LastName,
		PassengerPhonenumber: user.PhoneNumber,
		PassengerEmail:       user.Email,
	}
	tx.Commit()
	return res, nil
}
func (svc *ReservationService) CancelTicket(ctx context.Context, req *reservationpb.CancelTicketRequest) (*reservationpb.CancelTicketResponse, error) {
	tx := svc.DB.Begin()

	// Check if reservation exists and is confirmed

	// Fetch the ticket
	var ticket models.Ticket
	if err := tx.Model(&models.Ticket{}).
		Where("id = ? AND passenger_id = ?", req.TicketId, req.UserId).
		First(&ticket).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &reservationpb.CancelTicketResponse{
				Error: "Ticket not found",
			}, nil
		}
		return &reservationpb.CancelTicketResponse{
			Error: "Failed to retrieve ticket",
		}, nil
	}

	var reservation models.Reservation
	if err := tx.Model(&models.Reservation{}).
		Where("id = ? AND reservation_status = 'Confirmed' AND user_id = ?", ticket.ReservationID, req.UserId).
		First(&reservation).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &reservationpb.CancelTicketResponse{
				Error: "Reservation not found or already cancelled",
			}, nil
		}
		return &reservationpb.CancelTicketResponse{
			Error: "Failed to query reservation",
		}, nil
	}

	// Fetch the user
	var user models.User
	if err := tx.Model(&models.User{}).
		Where("id = ?", ticket.PassengerID).
		First(&user).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &reservationpb.CancelTicketResponse{
				Error: "User not found",
			}, nil
		}
		return &reservationpb.CancelTicketResponse{
			Error: "Failed to retrieve user",
		}, nil
	}

	// Update seat availability
	if err := tx.Model(&models.Seat{}).
		Where("id = ?", ticket.SeatID).
		Update("is_available", 1).Error; err != nil {
		tx.Rollback()
		return &reservationpb.CancelTicketResponse{
			Error: "Failed to update seat availability",
		}, nil
	}

	// Cancel the reservation
	if err := tx.Model(&models.Reservation{}).
		Where("id = ?", reservation.ID).
		Update("reservation_status", "Canceled").Error; err != nil {
		tx.Rollback()
		return &reservationpb.CancelTicketResponse{
			Error: "Failed to cancel reservation",
		}, nil
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return &reservationpb.CancelTicketResponse{
			Error: "Internal server error",
		}, nil
	}

	// Success
	return &reservationpb.CancelTicketResponse{}, nil
}

func GenerateRandomTicketNumber(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	ticketNumber := make([]byte, length)
	for i := range ticketNumber {
		ticketNumber[i] = charset[rand.Intn(len(charset))]
	}
	return string(ticketNumber)
}
func (svc *ReservationService) GetUserTickets(ctx context.Context, req *reservationpb.GetUserTicketRequest) (*reservationpb.GetUserTicketResponse, error) {
	var tickets []models.Ticket
	tx := svc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
		tx.Rollback()
	}()

	// Preload related models to reduce queries
	if err := tx.Preload("Reservation").
		Preload("Reservation.Flight").
		Preload("Reservation.Flight.DepartureAirport").
		Preload("Reservation.Flight.ArrivalAirport").
		Preload("Seat").
		Where("passenger_id = ?", req.UserId).
		Order("issue_date DESC").
		Find(&tickets).Error; err != nil {
		return &reservationpb.GetUserTicketResponse{
			Tickets: nil,
			Error:   fmt.Sprintf("failed to fetch tickets: %v", err),
		}, nil
	}

	var protoTickets []*reservationpb.Ticket
	for _, t := range tickets {
		protoTickets = append(protoTickets, convertToProtoTicket(t, t.Reservation, t.Reservation.Flight, t.Seat, t.Reservation.Flight.DepartureAirport, t.Reservation.Flight.ArrivalAirport))
	}

	tx.Commit()
	return &reservationpb.GetUserTicketResponse{
		Tickets: protoTickets,
	}, nil
}

func convertToProtoTicket(t models.Ticket, reservation models.Reservation, flight models.Flight, seat models.Seat, departureAirport, arrivalAirport models.Airport) *reservationpb.Ticket {
	return &reservationpb.Ticket{
		Reservation: &reservationpb.TReservation{
			ID:                uint32(reservation.ID),
			ReservationStatus: reservation.ReservationStatus,
			ReservationDate:   timestamppb.New(reservation.ReservationDate),
			PassportID:        int32(reservation.PassportID),
		},
		Seat: &reservationpb.TSeat{
			ID:         uint32(seat.ID),
			SeatNumber: seat.SeatNumber,
			SeatClass:  seat.SeatClass,
		},
		Flight: &reservationpb.TFlight{
			ID:            uint32(flight.ID),
			FlightNumber:  flight.FlightNumber,
			DepartureTime: timestamppb.New(flight.DepartureTime),
			ArrivalTime:   timestamppb.New(flight.ArrivalTime),
			DepartureAirport: &reservationpb.TAirport{
				AirportCode: departureAirport.AirportCode,
				AirportName: departureAirport.AirportName,
				Country:     departureAirport.Country,
				City:        departureAirport.City,
				Id:          strconv.Itoa(int(departureAirport.ID)),
			},
			ArrivalAirport: &reservationpb.TAirport{
				AirportCode: arrivalAirport.AirportCode,
				AirportName: arrivalAirport.AirportName,
				Country:     arrivalAirport.Country,
				City:        arrivalAirport.City,
				Id:          strconv.Itoa(int(arrivalAirport.ID)),
			},
			AvailableSeats:       uint32(flight.AvailableSeats),
			Duration:             uint32(flight.Duration),
			UpdatedDepartureTime: timestamppb.New(flight.UpdatedDepartureTime),
			UpdatedArrivalTime:   timestamppb.New(flight.UpdatedArrivalTime),
			Reason:               flight.RescheduleReason,
		},
		TicketNumber: t.TicketNumber,
		Price:        float32(t.Price),
		TicketID:     uint32(t.ID),
		IssueDate:    timestamppb.New(t.IssueDate),
		UserID:       uint32(t.PassengerID),
	}
}

func (svc *ReservationService) GetTickets(ctx context.Context, req *reservationpb.GetTicketRequest) (*reservationpb.GetTicketResponse, error) {
	var tickets []models.Ticket
	tx := svc.DB.Begin()

	if err := tx.Preload("Reservation").
		Preload("Reservation.Flight").
		Preload("Reservation.Flight.DepartureAirport").
		Preload("Reservation.Flight.ArrivalAirport").
		Preload("Seat").
		Order("issue_date DESC").
		Find(&tickets).Error; err != nil {
		return &reservationpb.GetTicketResponse{
			Tickets: nil,
			Error:   fmt.Sprintf("Failed to fetch tickets: %v", err),
		}, nil
	}

	var protoTickets []*reservationpb.Ticket
	for _, t := range tickets {
		protoTickets = append(protoTickets, convertToProtoTicket(t, t.Reservation, t.Reservation.Flight, t.Seat, t.Reservation.Flight.DepartureAirport, t.Reservation.Flight.ArrivalAirport))
	}

	tx.Commit()
	return &reservationpb.GetTicketResponse{
		Tickets: protoTickets,
	}, nil
}
