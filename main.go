package main

import (
	"Go_gRPC/config"
	"Go_gRPC/internal/db"
	"Go_gRPC/internal/services"
	"Go_gRPC/pb/airlinepb"
	"Go_gRPC/pb/airportpb"
	"Go_gRPC/pb/authpb"
	"Go_gRPC/pb/flightpb"
	"Go_gRPC/pb/reservationpb"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"

	"Go_gRPC/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type server struct {
	authpb.UnimplementedAuthenticationServiceServer
	reservationpb.UnimplementedReservationServiceServer
	airlinepb.UnimplementedAirlineServiceServer
	airportpb.UnimplementedAirportServiceServer
	flightpb.UnimplementedFlightServiceServer
	DB *gorm.DB
}

func main() {
	cfg := config.LoadConfig()
	ctx := context.Background()

	// Database setup
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Redis setup
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	rdb := redis.NewClient(&redis.Options{
		Username:  "red-csur2pt6l47c73813ojg",
		Addr:      "singapore-redis.render.com:6379",
		Password:  "WEYmIIgtoa8o8Fw0BadV70y7sUFpU3yP",
		DB:        0,
		TLSConfig: tlsConfig,
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// Start gRPC server
	go startGRPCServer(database, rdb)

	// Start Gin HTTP server
	startGinServer()
}

func startGRPCServer(database *gorm.DB, rdb *redis.Client) {
	lis, err := net.Listen("tcp", "0.0.0.0:443")
	if err != nil {
		log.Fatalf("Failed to listen on port 443: %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxRecvMsgSize(100*1024*1024),
		grpc.UnaryInterceptor(AuthInterceptor()),
	)
	authService := &services.AuthenticationServiceServer{DB: database}
	authpb.RegisterAuthenticationServiceServer(s, authService)
	reservationService := &services.ReservationService{DB: database}
	reservationpb.RegisterReservationServiceServer(s, reservationService)
	airlineService := &services.AirlineService{DB: database, Redis: rdb}
	airlinepb.RegisterAirlineServiceServer(s, airlineService)
	airportService := &services.AirportService{DB: database, Redis: rdb}
	airportpb.RegisterAirportServiceServer(s, airportService)
	flightService := &services.FlightService{DB: database, Redis: rdb}
	flightpb.RegisterFlightServiceServer(s, flightService)

	// Register reflection for gRPC server
	reflection.Register(s)

	log.Println("gRPC server is running on 0.0.0.0:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func startGinServer() {
	router := gin.Default()

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin HTTP server!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		// Example route for login
		var loginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Implement your login logic here
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	// Start Gin server
	go func() {
		if err := router.Run("0.0.0.0:8080"); err != nil {
			log.Fatalf("Failed to start Gin server: %v", err)
		}
	}()
}

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		log.Printf("Method Called: %s", info.FullMethod)
		if strings.Contains(info.FullMethod, "AuthenticationService") {
			return handler(ctx, req)
		}

		// Extract token from metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		token := md["authorization"]
		if len(token) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization token")
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(token[0], bearerPrefix) {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token format")
		}
		extractedToken := token[0][len(bearerPrefix):]
		log.Printf("Extracted Token: %s", extractedToken)

		claims := &utils.JWTClaims{}
		tkn, err := jwt.ParseWithClaims(extractedToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("jwtsecretkeyahihi"), nil
		})
		if err != nil || !tkn.Valid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		return handler(ctx, req)
	}
}
