package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgduncan/CryptoPro-Alexa-GCP/internal/alexa"
	"github.com/dgduncan/CryptoPro-Alexa-GCP/internal/coinbase"
	"github.com/google/uuid"
)

// AlexaHandler test
type AlexaHandler struct {
	Client *coinbase.CoinbaseClient
}

// Handle test
func (a *AlexaHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	rsp, err := a.Client.GetSpotPrice(ctx)
	if err != nil {
		fmt.Println("Crap")
		return
	}

	resp := alexa.Response{
		UUID:       uuid.New().String(),
		UpdateDate: time.Now().Format(time.RFC3339),
		TitleText:  "Test",
		MainText:   "The price of bitcoin is " + rsp.Data.Amount,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
