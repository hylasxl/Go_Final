package main

import (
	"Go_gRPC/internal/services"
	"Go_gRPC/pb/airlinepb"
	"Go_gRPC/pb/airportpb"
	"Go_gRPC/pb/authpb"
	"Go_gRPC/pb/flightpb"
	"Go_gRPC/pb/reservationpb"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"Go_gRPC/config"
	"Go_gRPC/internal/db"

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
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//if err := db.MigrateDB(database); err != nil {
	//	log.Fatalf("failed to migrate database: %v", err)
	//}
	// Ignore the second return value
	// Initialize the Redis client using parsed URL components
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // For skipping verification of the server's certificate
	}
	rdb := redis.NewClient(&redis.Options{
		Username: "red-csur2pt6l47c73813ojg",
		Addr:     "singapore-redis.render.com:6379", // Redis server host and port
		Password: "WEYmIIgtoa8o8Fw0BadV70y7sUFpU3yP",
		// Redis password (if any)
		DB:        0,         // Default database (0)
		TLSConfig: tlsConfig, // Enable TLS for secure connection (required for rediss://)
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong) // Should print "PONG" if connected successfully
	fixedIP := ""
	lis, err := net.Listen("tcp", fixedIP+":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpc.EnableTracing = true
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

	reflection.Register(s)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("failed to get network interfaces: %v", err)
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			fmt.Printf("Server is running on IP: %s\n", ipNet.IP.String())
		}
	}
	go func() {
		log.Println("Starting gRPC server on " + fixedIP + ":50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("Shutting down server...")

	s.GracefulStop()

	log.Println("Server stopped gracefully")
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
