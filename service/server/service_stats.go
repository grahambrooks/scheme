package server

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"net/http"
)

func (s *SchemeServer) ServiceStats(writer http.ResponseWriter, _ *http.Request) {
	es, err := elasticsearch.NewDefaultClient()

	req := esapi.CountRequest{Index: []string{"interfaces"}}

	res, err := req.Do(context.Background(), es)

	mirrorResponse(res, err, writer)
}
