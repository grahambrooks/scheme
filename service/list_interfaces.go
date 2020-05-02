package main

import (
	"context"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"net/http"
)

func ListInterfaces(writer http.ResponseWriter, request *http.Request) {
	log.Print("List Interfaces")
	es, err := elasticsearch.NewDefaultClient()
	req := esapi.SearchRequest{
		Index: []string{"interfaces"},
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	mirrorResponse(res, err, writer)
}

