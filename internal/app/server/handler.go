package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"user-service/internal/app/model"

	"github.com/rs/zerolog/log"
)

func (s *server) login(w http.ResponseWriter, r *http.Request) {
	var data model.User

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}

	if err := data.ValidateLogin(); err != nil {
		ErrInvalidEntityResponse(w, "Invalid Input", err)
		return
	}

	user, err := s.userService.GetUserByPhoneNumberAndPassword(r.Context(), data.PhoneNumber, data.Password)

	if err != nil {
		if errors.Is(err, model.ErrInvalid) || errors.Is(err, model.ErrNotFound) {
			ErrInvalidEntityResponse(w, "invalid credentials", err)
			return
		}
		log.Error().Err(err).Msgf("[login] failed to fetch login data Error: %v", err)
		ErrInternalServerResponse(w, "failed to fetch login data", err)
		return
	}

	token, err := s.authService.Encode(user)

	if err != nil {
		log.Error().Err(err).Msgf("[login] failed to generate jwt token Error: %v", err)
		ErrInternalServerResponse(w, "failed to generate jwt token", err)
		return
	}

	loginResponse := model.LoginResponse{
		ID:          user.ID,
		CategoryId:  user.CategoryId,
		PhoneNumber: user.PhoneNumber,
		Token:       token,
	}

	SuccessResponse(w, http.StatusOK, loginResponse)
}
