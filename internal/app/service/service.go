package service

import (
	"context"
	"user-service/internal/app/model"
)

//go:generate mockgen -source service.go -destination ./mocks/mock_service.go -package mocks

type UserService interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByPhoneAndPassword(ctx context.Context, phone, password string) (model.User, error)
}

type UserRepository interface {
	InsertUser(ctx context.Context, user model.User) error
	GetUserByPhone(ctx context.Context, phone string) (model.User, error)
}

type AuthService interface {
	Decode(tokenStr string) (*model.JwtCustomClaims, error)
	Encode(user model.User) (string, error)
}
