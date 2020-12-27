package server

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gorilla/mux"
	"github.com/grahambrooks/scheme/service/store"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

type SchemeServer struct {
	Server   *http.Server
	Router   *mux.Router
	Port     int
	ApiStore store.ApiStore
}

func NewSchemeServer(port int, apiStore store.ApiStore) *SchemeServer {
	log.SetFormatter(&log.JSONFormatter{})
	s := SchemeServer{Port: port, ApiStore: apiStore}
	r := mux.NewRouter()
	s.Router = r
	r.Use(loggingMiddleware)
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", s.HomeHandler)
	api.HandleFunc("/info", s.InfoHandler).Methods(http.MethodGet)
	api.HandleFunc("/stats", s.ServiceStats)
	api.HandleFunc("/search", s.SearchApiHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis", s.ListApisHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis/{id}", s.NewApiHandler).Methods(http.MethodPost)
	api.HandleFunc("/apis/{id}", s.GetApiHandler).Methods(http.MethodGet)
	api.HandleFunc("/apis/{id}/updates", s.GetApiHandler).Methods(http.MethodPost)
	api.HandleFunc("/registrations", s.NewRegistration).Methods(http.MethodPost)

	contentPath := ContentPath()
	view := ApiView{Path: contentPath, ApiStore: s.ApiStore}

	r.Path("/view/{id}").HandlerFunc(view.ViewHandler).Methods(http.MethodGet)
	s.Log("Serving static content from %s", contentPath)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(contentPath))))

	s.Server = &http.Server{
		Handler:      r,
		Addr:         ":" + strconv.Itoa(s.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &s
}

func (s *SchemeServer) Log(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (s *SchemeServer) ListenAndServe() {
	s.Log("Service starting on port %d", s.Port)
	log.Fatal(s.Server.ListenAndServe())
}

func ContentPath() string {
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
