package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

var coinbaseClient CoinbaseClient

type alexaResponse struct {
	UUID       string `json:"uid"`
	UpdateDate string `json:"updateDate"`
	TitleText  string `json:"titleText"`
	MainText   string `json:"mainText"`
}

func init() {
	client := http.DefaultClient
	coinbaseClient = CoinbaseClient{
		HTTP: client,
	}

}

func main() {
	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/alexa", func(w http.ResponseWriter, r *http.Request) {
		rsp, err := coinbaseClient.GetSpotPrice(r.Context())
		if err != nil {
			fmt.Println("Crap")
		}

		resp := alexaResponse{
			UUID:       uuid.New().String(),
			UpdateDate: time.Now().Format(time.RFC3339),
			TitleText:  "Test",
			MainText:   "The price of bitcoin is " + rsp.Data.Amount,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	})
	http.ListenAndServe(":"+port, r)
}
