package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"net/http"
	"strings"
)

func SearchHandler(writer http.ResponseWriter, request *http.Request) {
	filter := request.FormValue("query")
	query := fmt.Sprintf(`{
  "query": {
    "query_string": {
      "query": "%s"
    }
  }
}`, filter)
	log.Printf("Search %s\n", filter)
	es, err := elasticsearch.NewDefaultClient()
	req := esapi.SearchRequest{
		Index: []string{"interfaces"},
		Body:  strings.NewReader(query),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	mirrorResponse(res, err, writer)
}
