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

func TestService_GetUserByPhoneNumberAndPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	password, _ := util.HashPassword("123456")
	user := model.User{
		ID:          1,
		PhoneNumber: "+880123456",
		Password:    password,
	}

	testCases := []struct {
		desc        string
		phoneNumber string
		password    string
		mockRepo    func() *mocks.MockUserRepository
		expErr      error
		expUser     model.User
	}{
		{
			desc:        "should return success",
			phoneNumber: "+880123456",
			password:    "123456",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhoneNumber(gomock.Any(), "+880123456").Return(user, nil)
				return r
			},
			expErr:  nil,
			expUser: user,
		},
		{
			desc:        "should return DB error",
			phoneNumber: "+880123456",
			password:    "123456",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhoneNumber(gomock.Any(), "+880123456").Return(model.User{}, errors.New("db-error"))
				return r
			},
			expErr:  errors.New("db-error"),
			expUser: model.User{},
		},
		{
			desc:        "should return wrong password error",
			phoneNumber: "+880123456",
			password:    "wrong-password",
			mockRepo: func() *mocks.MockUserRepository {
				r := mocks.NewMockUserRepository(ctrl)
				r.EXPECT().GetUserByPhoneNumber(gomock.Any(), "+880123456").Return(user, nil)
				return r
			},
			expErr:  fmt.Errorf("wrong password :%w", model.ErrInvalid),
			expUser: model.User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewService(tc.mockRepo())
			user, err := s.GetUserByPhoneNumberAndPassword(context.Background(), tc.phoneNumber, tc.password)
			assert.Equal(t, tc.expErr, err)
			assert.EqualValues(t, tc.expUser, user)
		})
	}
}
