package server

import (
	"net/http"
)

func (s *server) test(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	SuccessResponse(w, http.StatusOK, token)
}
