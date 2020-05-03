package main

import (
	"context"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"net/http"
)

func ListApisHandler(writer http.ResponseWriter, request *http.Request) {
	es, err := elasticsearch.NewDefaultClient()
	req := esapi.SearchRequest{
		Index: []string{SearchIndexName},
	}

	res, err := req.Do(context.Background(), es)
	mirrorResponse(res, err, writer)
}

