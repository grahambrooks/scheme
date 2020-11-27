package main

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"net/http"
)

func (s *SchemeServer) HomeHandler(writer http.ResponseWriter, _ *http.Request) {

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
