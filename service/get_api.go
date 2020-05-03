package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"net/http"
)

type ElasticGetResponse struct {
	Id     string `json:"_id"`
	Source struct {
		Content string `json:"Content"`
	} `json:"_source"`
}

func GetApiHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) == 0 {
		errorResponse(writer, "Invalid document id (empty)")
	} else {
		es, _ := elasticsearch.NewDefaultClient()
		req := esapi.GetRequest{
			Index:      DocumentIndexName,
			DocumentID: id,
		}
		res, _ := req.Do(context.Background(), es)

		decoder := json.NewDecoder(res.Body)
		var elasticResponse ElasticGetResponse
		_ = decoder.Decode(&elasticResponse)

		_, _ = writer.Write([]byte(elasticResponse.Source.Content))
	}
}
