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

type AlexaResponse struct {
	UUID       string `json:"uid"`
	UpdateDate string `json:"updateDate"`
	TitleText  string `json:"titleText"`
	MainText   string `json:"mainText"`
}

func main() {
	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
		// w.Write([]byte("welcome"))
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/alexa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(uuid.New())
		resp := AlexaResponse{
			UUID:       uuid.New().String(),
			UpdateDate: time.Now().Format(time),
			TitleText:  "Test",
			MainText:   "Test",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	})
	http.ListenAndServe(":"+port, r)
	// http.HandleFunc("/", indexHandler)
	// http.Handle

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("Defaulting to port %s", port)
	// }

	// log.Printf("Listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	fmt.Fprint(w, "Hello, World!")
// }
