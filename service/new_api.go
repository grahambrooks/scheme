package main

import (
	"apellicon/openapi"
	"apellicon/search"
	"apellicon/wadl"
	"bytes"
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func NewApiHandler(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) == 0 {
		errorResponse(writer, "Invalid document id (empty)")
	} else {
		document, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		WriteApiEntry(id, string(document))
		model, err := interfaceModel(contentType, ioutil.NopCloser(bytes.NewReader(document)))
		if err != nil {
			errorResponse(writer, fmt.Sprintf("error parsing request %v", err))
		} else {
			var buffer bytes.Buffer
			err = model.AsJson(&buffer)
			if err != nil {
				errorResponse(writer, fmt.Sprintf("Failed to encode model %v", err))
			} else {
				log.Printf("Search Model %s", buffer.String())

				es, err := elasticsearch.NewDefaultClient()
				req := esapi.IndexRequest{
					Index:      SearchIndexName,
					DocumentID: id,
					Body:       bytes.NewReader(buffer.Bytes()),
					Refresh:    "true",
					ErrorTrace: true,
				}

				res, err := req.Do(context.Background(), es)
				mirrorResponse(res, err, writer)
			}
		}
	}
}

func interfaceModel(contentType string, spec io.ReadCloser) (search.Model, error) {
	switch contentType {
	case "application/openapi+json":
		parser := openapi.Parser{}
		return parser.ParseJson(spec)
	case "application/openapi+yaml":
		parser := openapi.Parser{}
		return parser.ParseYaml(spec)
	case "application/wadl+xml":
		parser := wadl.Parser{}
		return parser.Parse(spec)
	default:
		return search.Model{}, fmt.Errorf("Content type '%s' not suported. Supported content types application/openapi2+json, applicatiion/openapi2+json", contentType)
	}

}
