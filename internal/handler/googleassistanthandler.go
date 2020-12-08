package handler

import (
	"net/http"
)

// GoogleAssistantHandler test
type GoogleAssistantHandler struct {
}

// Handle test
func (a *GoogleAssistantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	// w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
