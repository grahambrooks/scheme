package main

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	SearchIndexName = "apellicon-search"
	DocumentIndexName = "apellicon-docs"
)

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", HomeHandler)
	api.HandleFunc("/stats", ServiceStats)
	api.HandleFunc("/search", SearchApiHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis", ListApisHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis/{id}", NewApiHandler).Methods(http.MethodPost)
	api.HandleFunc("/apis/{id}", GetApiHandler).Methods(http.MethodGet)
	//http.Handle("/", r)

	contentPath := "site"
	_, err := os.Stat(contentPath)
	if err != nil {
		contentPath = "service/site"
	}

	log.Printf("Serving static content from %s", contentPath)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(contentPath))))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Service running on port 8000")

	log.Fatal(srv.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func mirrorResponse(res *esapi.Response, err error, writer http.ResponseWriter) {
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	decoder := json.NewDecoder(res.Body)
	var body interface{}
	_ = decoder.Decode(&body)

	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(body)

	_ = res.Body.Close()
}

type ErrorMessage struct {
	Message string
}

func errorResponse(writer http.ResponseWriter, s string) {
	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(ErrorMessage{Message: s})
	writer.WriteHeader(http.StatusBadRequest)

}

