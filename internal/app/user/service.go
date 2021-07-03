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

func NewService(repo svc.UserRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(ctx context.Context, user model.User) error {

	if user.PhoneNumber == "" || user.Password == "" {
		return fmt.Errorf("invalid user request :%w", model.ErrInvalid)
	}

	if err := model.SignUpPhoneValidation(user.PhoneNumber); !err {
		return fmt.Errorf("invalid phone number :%w", model.ErrInvalid)
	}

	if p, err := util.HashPassword(user.Password); err == nil {
		user.Password = p
	}
	return s.repo.InsertUser(ctx, user)
}

func (s *service) GetUserByPhoneAndPassword(ctx context.Context, phone_number, password string) (model.User, error) {
	if phone_number == "" || password == "" {
		return model.User{}, fmt.Errorf("invalid login request :%w", model.ErrInvalid)
	}
	user, err := s.repo.GetUserByPhone(ctx, phone_number)
	if err != nil {
		return model.User{}, err
	}
	if !util.CheckPasswordHash(password, user.Password) {
		return model.User{}, fmt.Errorf("wrong password :%w", model.ErrInvalid)
	}
	return user, nil
}
