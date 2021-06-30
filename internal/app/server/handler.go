package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"user-service/internal/app/model"

	"github.com/rs/zerolog/log"
)

func (s *server) signUp(w http.ResponseWriter, r *http.Request) {
	var data model.User

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		ErrUnprocessableEntityResponse(w, "Decode Error", err)
		return
	}

	if err := s.userService.CreateUser(r.Context(), data); err != nil {
		if errors.Is(err, model.ErrInvalid) {
			ErrInvalidEntityResponse(w, "invalid user", err)
			return
		}
		log.Error().Err(err).Msgf("[signUp] failed to create user Error: %v", err)
		ErrInternalServerResponse(w, "failed to create user", err)
		return
	}
	SuccessResponse(w, http.StatusCreated, "successful")
}
