package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-service/internal/app/model"
	"user-service/internal/app/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := model.User{
		ID:          1,
		Password:    "123456",
		PhoneNumber: "123456",
		CategoryId:  1,
	}

	testCases := []struct {
		desc          string
		payload       string
		mockUserSvc   func() *mocks.MockUserService
		mockAuthSvc   func() *mocks.MockAuthService
		expStatusCode int
	}{
		{
			desc:    "should success",
			payload: `{ "phone_number": "123456", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneNumberAndPassword(gomock.Any(), "123456", "123456").Return(user, nil)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				s := mocks.NewMockAuthService(ctrl)
				s.EXPECT().Encode(user).Return("auto-token", nil)
				return s
			},
			expStatusCode: http.StatusOK,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockUserSvc: func() *mocks.MockUserService {
				return mocks.NewMockUserService(ctrl)
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				return mocks.NewMockAuthService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
		},
		{
			desc:    "should return invalid credentials error",
			payload: `{ "phone_number": "123456", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneNumberAndPassword(gomock.Any(), "123456", "123456").Return(model.User{}, model.ErrInvalid)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				return mocks.NewMockAuthService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
		},
		{
			desc:    "should return internal server error",
			payload: `{ "phone_number": "123456", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneNumberAndPassword(gomock.Any(), "123456", "123456").Return(model.User{}, errors.New("server-error"))
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				return mocks.NewMockAuthService(ctrl)
			},
			expStatusCode: http.StatusInternalServerError,
		},
		{
			desc:    "should return internal server error for jwt",
			payload: `{ "phone_number": "123456", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneNumberAndPassword(gomock.Any(), "123456", "123456").Return(user, nil)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				s := mocks.NewMockAuthService(ctrl)
				s.EXPECT().Encode(user).Return("", errors.New("jwt-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockUserSvc(), tc.mockAuthSvc())

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/login", body)

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/login").HandlerFunc(s.login)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
		})
	}
}
