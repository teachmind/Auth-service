package server

import (
	"net/http"
)

func (s *server) tokenValidation(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	claim, err := s.authService.Decode(token)
	if err != nil {
		ErrUnauthorizedResponse(w, "invalid token", err)
		return
	}
	SuccessResponse(w, http.StatusOK, claim.User)
}
