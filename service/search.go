package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"net/http"
	"strings"
)

func SearchApiHandler(writer http.ResponseWriter, request *http.Request) {
	filter := request.FormValue("query")

	if len(filter) > 0 {
		query := fmt.Sprintf(`{
  "query": {
    "query_string": {
      "query": "%s"
    }
  }
}`, filter)
		es, err := elasticsearch.NewDefaultClient()
		req := esapi.SearchRequest{
			Index: []string{SearchIndexName},
			Body:  strings.NewReader(query),
		}
		res, err := req.Do(context.Background(), es)
		mirrorResponse(res, err, writer)
	} else {
		ListApisHandler(writer, request)
	}
}
