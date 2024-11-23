package services

import (
	models2 "Go_gRPC/internal/models"
	"Go_gRPC/pb/authpb"
	"Go_gRPC/pkg/utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"strings"
	"time"
)

type AuthenticationServiceServer struct {
	authpb.UnimplementedAuthenticationServiceServer
	DB *gorm.DB
}

func (svc *AuthenticationServiceServer) SignIn(ctx context.Context, req *authpb.SignInRequest) (*authpb.SignInResponse, error) {
	var user models2.User

	if err := svc.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return &authpb.SignInResponse{Error: "The username or password is not correct"}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return &authpb.SignInResponse{Error: "The password is invalid"}, nil
	}

	claims := utils.JWTClaims{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		RoleID:      user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	accessToken, err := utils.GenerateAccessToken(claims)
	if err != nil {
		return &authpb.SignInResponse{Error: "Failed to generate access token: " + err.Error()}, nil
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Username)
	if err != nil {
		return &authpb.SignInResponse{Error: "Failed to generate refresh token: " + err.Error()}, nil
	}

	return &authpb.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Claims: &authpb.JWTClaims{
			Username:    claims.Username,
			FirstName:   claims.FirstName,
			LastName:    claims.LastName,
			Email:       claims.Email,
			PhoneNumber: claims.PhoneNumber,
			DateOfBirth: timestamppb.New(claims.DateOfBirth),
			Address:     claims.Address,
			RoleId:      uint32(claims.RoleID),
			IsActive:    claims.IsActive,
		},
		UserId: uint32(user.ID),
	}, nil
}

func (svc *AuthenticationServiceServer) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	var existingUser models2.User

	req.Email = strings.ToLower(req.Email)

	if err := svc.DB.Where("username = ? OR email = ? OR phone_number = ?", req.Username, req.Email, req.PhoneNumber).First(&existingUser).Error; err == nil {
		if existingUser.Username == req.Username {
			return &authpb.RegisterResponse{Error: "Username already exists"}, nil
		}
		if existingUser.Email == req.Email {
			return &authpb.RegisterResponse{Error: "Email already exists"}, nil
		}
		if existingUser.PhoneNumber == req.PhoneNumber {
			return &authpb.RegisterResponse{Error: "Phone number already exists"}, nil
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return &authpb.RegisterResponse{Error: "Failed to hash password: " + err.Error()}, nil
	}

	newUser := models2.User{
		Username:    req.Username,
		Password:    string(hashedPassword),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		DateOfBirth: req.DateOfBirth.AsTime(),
		Address:     req.Address,
		RoleID:      1,
	}

	if err := svc.DB.Create(&newUser).Error; err != nil {
		return &authpb.RegisterResponse{Error: "Failed to register user: " + err.Error()}, nil
	}

	return &authpb.RegisterResponse{Error: ""}, nil
}

func (svc *AuthenticationServiceServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		return &authpb.RefreshTokenResponse{Error: "Invalid or expired refresh token"}, nil
	}

	var user models2.User
	if err := svc.DB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
		return &authpb.RefreshTokenResponse{Error: "User not found"}, nil
	}

	newClaims := utils.JWTClaims{
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		DateOfBirth: user.DateOfBirth,
		Address:     user.Address,
		RoleID:      user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
			Subject:   user.Username,
			IssuedAt:  time.Now().Unix(),
		},
	}

	newAccessToken, err := utils.GenerateAccessToken(newClaims)
	if err != nil {
		return &authpb.RefreshTokenResponse{Error: "Failed to generate new access token: " + err.Error()}, nil
	}

	return &authpb.RefreshTokenResponse{
		AccessToken: newAccessToken,
	}, nil
}
