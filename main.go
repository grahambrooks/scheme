package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", HomeHandler)
	api.HandleFunc("/stats", StatsHandler)
	api.HandleFunc("/interfaces", GetInterfaces).Methods(http.MethodGet)
	api.HandleFunc("/interfaces", SaveInterface).Methods(http.MethodPost)
	//http.Handle("/", r)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("site"))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
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

func StatsHandler(writer http.ResponseWriter, request *http.Request) {
	es, err := elasticsearch.NewDefaultClient()

	req := esapi.CountRequest{Index: []string{"interfaces"}}

	res, err := req.Do(context.Background(), es)

	mirrorResponse(res, err, writer)
}

func SaveInterface(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	if contentType == "application/json" {
		if request.ContentLength == 0 {
			log.Fatalf("No content provided in request")
		}
		es, err := elasticsearch.NewDefaultClient()
		req := esapi.IndexRequest{
			Index: "interfaces",
			//DocumentID: strconv.Itoa(i + 1),
			Body:    request.Body,
			Refresh: "true",
		}

		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		mirrorResponse(res, err, writer)
	}
	encoder := json.NewEncoder(writer)

	_ = encoder.Encode("Save Interface")

}

func GetInterfaces(writer http.ResponseWriter, request *http.Request) {
	filter := request.FormValue("query")
	query := fmt.Sprintf(`{
  "query": {
    "query_string": {
      "query": "%s"
    }
  }
}`, filter)
	fmt.Printf("ES Query %s\n", query)
	es, err := elasticsearch.NewDefaultClient()
	req := esapi.SearchRequest{
		Index: []string{"interfaces"},
		Body:  strings.NewReader(query),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	mirrorResponse(res, err, writer)
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	decoder := json.NewDecoder(res.Body)
	var body interface{}
	_ = decoder.Decode(&body)

	encoder := json.NewEncoder(writer)

	_ = encoder.Encode(struct {
		StatusCode int
		Content    interface{}
	}{StatusCode: res.StatusCode, Content: body})
}
