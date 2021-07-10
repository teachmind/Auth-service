package server

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"user-service/internal/app/model"
	"user-service/internal/app/service/mocks"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		desc          string
		payload       string
		mockSvc       func() *mocks.MockUserService
		expStatusCode int
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
		},
		{
			desc:    "should return decode error",
			payload: `------------`,
			mockSvc: func() *mocks.MockUserService {
				return mocks.NewMockUserService(ctrl)
			},
			expStatusCode: http.StatusUnprocessableEntity,
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
		})
	}
}