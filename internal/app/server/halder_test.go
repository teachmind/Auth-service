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

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		mockSvc       func() *mocks.MockUserService
		expStatusCode int
		expResponse string
	}{
		{
			desc:    "should success",
			payload: `{ "phone_number": "01738799349", "password": "123456" }`,
			mockSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil)
				return s
			},
			expStatusCode: http.StatusCreated,
			expResponse:   `{"success":true,"errors":null,"data":"successful"}`,
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockSvc: func() *mocks.MockUserService {
				return mocks.NewMockUserService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid user error",
			payload: `{ "phone_number": "01738799349", "password": "123456" }`,
			mockSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(model.ErrInvalid)
				return s
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid user","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: `{ "phone_number": "01738799349", "password": "123456" }`,
			mockSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(errors.New("server-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to create user","severity":"error"}],"data":null}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			s := NewServer(":8080", tc.mockSvc(), nil)

			w := httptest.NewRecorder()
			body := strings.NewReader(tc.payload)
			r := httptest.NewRequest(http.MethodPost, "/api/v1/signup", body)

			router := mux.NewRouter()
			router.Methods(http.MethodPost).Path("/api/v1/signup").HandlerFunc(s.signUp)
			router.ServeHTTP(w, r)
			assert.Equal(t, tc.expStatusCode, w.Code)
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := model.User{
		ID:          1,
		PhoneNumber: "+8801712345678",
		Password:    "123456",
		CategoryID:  1,
	}

	testCases := []struct {
		desc          string
		payload       string
		mockUserSvc   func() *mocks.MockUserService
		mockAuthSvc   func() *mocks.MockAuthService
		expStatusCode int
		expResponse   string
	}{
		{
			desc:    "should success",
			payload: `{ "phone_number": "+8801712345678", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneAndPassword(gomock.Any(), "+8801712345678", "123456").Return(user, nil)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				s := mocks.NewMockAuthService(ctrl)
				s.EXPECT().Encode(user).Return("auto-token", nil)
				return s
			},
			expStatusCode: http.StatusOK,
			expResponse:   `{"success":true,"errors":null,"data":{"id":1,"phone_number":"+8801712345678","category_id":1,"token":"auto-token"}}`,
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
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid character '-' in numeric literal","message_title":"Decode Error","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return invalid credentials error",
			payload: `{ "phone_number": "+8801712345678", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneAndPassword(gomock.Any(), "+8801712345678", "123456").Return(model.User{}, model.ErrInvalid)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				return mocks.NewMockAuthService(ctrl)
			},
			expStatusCode: http.StatusBadRequest,
			expResponse:   `{"success":false,"errors":[{"code":"INVALID","message":"invalid","message_title":"invalid credentials","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error",
			payload: `{ "phone_number": "+8801712345678", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneAndPassword(gomock.Any(), "+8801712345678", "123456").Return(model.User{}, errors.New("server-error"))
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				return mocks.NewMockAuthService(ctrl)
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"server-error","message_title":"failed to fetch login data","severity":"error"}],"data":null}`,
		},
		{
			desc:    "should return internal server error for jwt",
			payload: `{ "phone_number": "+8801712345678", "password": "123456" }`,
			mockUserSvc: func() *mocks.MockUserService {
				s := mocks.NewMockUserService(ctrl)
				s.EXPECT().GetUserByPhoneAndPassword(gomock.Any(), "+8801712345678", "123456").
					Return(user, nil)
				return s
			},
			mockAuthSvc: func() *mocks.MockAuthService {
				s := mocks.NewMockAuthService(ctrl)
				s.EXPECT().Encode(user).Return("", errors.New("jwt-error"))
				return s
			},
			expStatusCode: http.StatusInternalServerError,
			expResponse:   `{"success":false,"errors":[{"code":"SERVER_ERROR","message":"jwt-error","message_title":"failed to generate jwt token","severity":"error"}],"data":null}`,
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
			assert.Equal(t, tc.expResponse, w.Body.String())
		})
	}
}
