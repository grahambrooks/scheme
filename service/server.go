package main

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type SchemeServer struct {
	Port     int
	ApiStore ApiStore
}

func (s *SchemeServer) ListenAndServe() {
	log.SetFormatter(&log.JSONFormatter{})
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", s.HomeHandler)
	api.HandleFunc("/stats", s.ServiceStats)
	api.HandleFunc("/search", s.SearchApiHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis", s.ListApisHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis/{id}", s.NewApiHandler).Methods(http.MethodPost)
	api.HandleFunc("/apis/{id}", s.GetApiHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis/{id}/updates", s.GetApiHandler).Methods(http.MethodPost)
	api.HandleFunc("/registrations", s.NewRegistration).Methods(http.MethodPost)

	view := ApiView{Path: s.ContentPath(), ApiStore: s.ApiStore}

	r.Path("/view/{id}").HandlerFunc(view.ViewHandler).Methods(http.MethodGet)
	log.Printf("Serving static content from %s", s.ContentPath())
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(s.ContentPath()))))

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Service running on port 8000")
	log.Fatal(srv.ListenAndServe())

}

func (s *SchemeServer) ContentPath() string {
	contentPath := "site"
	_, err := os.Stat(contentPath)
	if err != nil {
		contentPath = "service/site"
	}
	return contentPath
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func mirrorResponse(res *esapi.Response, err error, writer http.ResponseWriter) {
	if err != nil {
		errorResponse(writer, "No connection to Elastic service")
	} else {
		decoder := json.NewDecoder(res.Body)
		var body interface{}
		_ = decoder.Decode(&body)

		encoder := json.NewEncoder(writer)
		_ = encoder.Encode(body)

		_ = res.Body.Close()
	}
}
