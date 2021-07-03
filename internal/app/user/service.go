package user

import (
	"context"
	"fmt"
	"user-service/internal/app/model"
	svc "user-service/internal/app/service"
	"user-service/internal/app/util"
)

type service struct {
	repo svc.UserRepository
}

// NewService Initiates new user repository service
func NewService(repo svc.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

/* GetUserByPhoneNumberAndPassword to get User by PhoneNumber and Password */
func (s *service) GetUserByPhoneNumberAndPassword(ctx context.Context, phone_number, password string) (model.User, error) {
	if phone_number == "" || password == "" {
		return model.User{}, fmt.Errorf("invalid login request :%w", model.ErrInvalid)
	}
	user, err := s.repo.GetUserByPhoneNumber(ctx, phone_number)
	if err != nil {
		return model.User{}, err
	}
	if !util.CheckPasswordHash(password, user.Password) {
		return model.User{}, fmt.Errorf("wrong password :%w", model.ErrInvalid)
	}
	return user, nil
}
