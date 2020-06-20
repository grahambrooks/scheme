package main

import (
	"context"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"net/http"
)

func (s *SchemeServer) ServiceStats(writer http.ResponseWriter, request *http.Request) {
	es, err := elasticsearch.NewDefaultClient()

	req := esapi.CountRequest{Index: []string{"interfaces"}}

	res, err := req.Do(context.Background(), es)

	mirrorResponse(res, err, writer)
}
