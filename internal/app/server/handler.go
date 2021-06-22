package server

import (
	"net/http"
	"user-service/internal/app/model"
)

func (s *server) signUp(w http.ResponseWriter, r *http.Request) {
	var data model.User

	SuccessResponse(w, http.StatusCreated, data)
}
