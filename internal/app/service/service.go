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

type UserRepository interface {
	GetUserByPhoneNumber(ctx context.Context, phone_number string) (model.User, error)
}

type UserService interface {
	GetUserByPhoneNumberAndPassword(ctx context.Context, phone_number, password string) (model.User, error)
}
