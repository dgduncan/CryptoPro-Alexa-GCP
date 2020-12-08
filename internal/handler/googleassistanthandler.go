package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GoogleAssistantHandler test
type GoogleAssistantHandler struct {
}

// Handle test
func (a *GoogleAssistantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bodyBytes))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
