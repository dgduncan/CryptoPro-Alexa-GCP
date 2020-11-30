package main

import (
	"net/http"
	"os"

	"github.com/dgduncan/CryptoPro-Alexa-GCP/internal/coinbase"
	"github.com/dgduncan/CryptoPro-Alexa-GCP/internal/handler"
	"github.com/go-chi/chi"
)

func init() {
}

func main() {

	// ********* Initializers ***********

	coinbaseClient := &coinbase.CoinbaseClient{
		HTTP: http.DefaultClient,
	}
	alexaHandler := &handler.AlexaHandler{
		Client: coinbaseClient,
	}

	// ***********************************

	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Get("/alexa", alexaHandler.Handle)
	http.ListenAndServe(":"+port, r)
}
