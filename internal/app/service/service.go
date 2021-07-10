package service

import (
	"context"
	"user-service/internal/app/model"
)

//go:generate mockgen -source service.go -destination ./mocks/mock_service.go -package mocks

type AuthService interface {
	Decode(tokenStr string) (*model.JwtCustomClaims, error)
	Encode(user model.User) (string, error)
}

// UserRepository to fetch user by PhoneNumber
type UserRepository interface {
	InsertUser(ctx context.Context, user model.User) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error)
}

// UserService to fetch user by PhoneNumber and Password
type UserService interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByPhoneAndPassword(ctx context.Context, phoneNumber, password string) (model.User, error)
}
