package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	SearchIndexName   = "apellicon-search"
	DocumentIndexName = "apellicon-docs"
)

type ApiView struct {
	Path  string
	Id    string
	Title string
	Api   interface{}
}

func (v ApiView) readApiSpec(id string) string {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.GetRequest{
		Index:      DocumentIndexName,
		DocumentID: id,
	}
	res, _ := req.Do(context.Background(), es)

	decoder := json.NewDecoder(res.Body)
	var elasticResponse ElasticGetResponse
	_ = decoder.Decode(&elasticResponse)

	return elasticResponse.Source.Content
}

func (v ApiView) ViewHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	spec := v.readApiSpec(id)

	api, err := v.decodeApiSpec(spec)

	templateFilePath := filepath.Join(v.Path, "templates", "view.html")
	t, err := template.New("view").ParseFiles(templateFilePath)

	if err != nil {
		log.Printf("error reading template file %s %v", templateFilePath, err)
	}
	v.Title = "some sort of title"
	v.Id = id
	v.Api = api
	err = t.Execute(writer, v)
	if err != nil {
		fmt.Printf("Error rendering tempalte %v", err)
	}
}

func (v ApiView) decodeApiSpec(spec string) (interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(spec))

	var api interface{}
	err := decoder.Decode(&api)
	if err != nil {
		decoder := yaml.NewDecoder(strings.NewReader(spec))
		err = decoder.Decode(&api)
		if err != nil {
			log.Printf("error reading API %v", err)
		}
	}
	return api, err
}

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
	view := ApiView{Path: contentPath}

	r.Path("/view/{id}").HandlerFunc(view.ViewHandler).Methods(http.MethodGet)
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
