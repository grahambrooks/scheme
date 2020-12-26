package server

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"net/http"
)

func (s *SchemeServer) HomeHandler(writer http.ResponseWriter, _ *http.Request) {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		s.Log("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		s.Log("Error getting response: %s", err)
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
