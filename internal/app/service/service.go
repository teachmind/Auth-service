package service

import (
	"user-service/internal/app/model"
)

//go:generate mockgen -source service.go -destination ./mocks/mock_service.go -package mocks
type AuthService interface {
	Decode(tokenStr string) (*model.JwtCustomClaims, error)
	Encode(user model.User) (string, error)
}
