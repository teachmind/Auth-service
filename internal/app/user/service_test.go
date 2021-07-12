package user

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"user-service/internal/app/model"
	"user-service/internal/app/service/mocks"
	"user-service/internal/app/util"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc     string
		payload  model.User
		mockRepo func() *mocks.MockUserRepository
		expErr   error
	}{
		{
			desc: "should return success",
			payload: model.User{
				PhoneNumber: "01738799349",
				Password:    "123456",
				CategoryId:  1,
			},
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(nil)
				return r
			},
			expErr: nil,
		},

		{
			desc: "should return db error",
			payload: model.User{
				PhoneNumber: "01738799349",
				Password:    "12345",
				CategoryId:  1,
			},
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(errors.New("db-error"))
				return r
			},
			expErr: errors.New("db-error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			err := s.CreateUser(context.Background(), tc.payload)
			assert.Equal(t, tc.expErr, err)
		})
	}
}

func TestService_GetUserByPhoneAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password, _ := util.HashPassword("123456")
	user := model.User{
		ID:          1,
		PhoneNumber: "01738799349",
		Password:    password,
		CategoryId:  1,
	}

	testCases := []struct {
		desc     string
		phone    string
		password string
		mockRepo func() *mocks.MockUserRepository
		expErr   error
		expUser  model.User
	}{
		{
			desc:     "should return success",
			phone:    "01738799349",
			password: "123456",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhone(gomock.Any(), "01738799349").Return(user, nil)
				return r
			},
			expErr:  nil,
			expUser: user,
		},
		{
			desc:     "should return invalid request error",
			phone:    "",
			password: "",
			mockRepo: func() *mocks.MockUserRepository {
				return mocks.NewMockUserRepository(ctrl)
			},
			expErr:  fmt.Errorf("invalid login request :%w", model.ErrInvalid),
			expUser: model.User{},
		},
		{
			desc:     "should return DB error",
			phone:    "01738799349",
			password: "123456",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhone(gomock.Any(), "01738799349").Return(model.User{}, errors.New("db-error"))
				return r
			},
			expErr:  errors.New("db-error"),
			expUser: model.User{},
		},
		{
			desc:     "should return wrong password error",
			phone:    "01738799349",
			password: "wrong-password",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhone(gomock.Any(), "01738799349").Return(user, nil)
				return r
			},
			expErr:  fmt.Errorf("wrong password :%w", model.ErrInvalid),
			expUser: model.User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			user, err := s.GetUserByPhoneAndPassword(context.Background(), tc.phone, tc.password)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expUser, user)
		})
	}
}
