package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grahambrooks/scheme/openapi"
	"github.com/grahambrooks/scheme/search"
	"github.com/grahambrooks/scheme/wadl"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *SchemeServer) NewApiHandler(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) == 0 {
		errorResponse(writer, "Invalid document id (empty)")
	} else {
		document, err := ioutil.ReadAll(request.Body)
		defer func() { _ = request.Body.Close() }()
		_, err = s.ApiStore.Save(id, string(document))

		model, err := parseContent(contentType, ioutil.NopCloser(bytes.NewReader(document)))
		if err != nil {
			errorResponse(writer, fmt.Sprintf("error parsing request %v", err))
		} else {
			var buffer bytes.Buffer
			err = model.AsJson(&buffer)
			if err != nil {
				errorResponse(writer, fmt.Sprintf("Failed to encode model %v", err))
			} else {
				log.Printf("Search Model %s", buffer.String())

				res, err := s.ApiStore.IndexDocument(id, buffer.Bytes())
				mirrorResponse(res, err, writer)
			}
		}
	}
}

func parseContent(contentType string, spec io.ReadCloser) (search.Model, error) {
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
	case "":
		return search.Model{}, fmt.Errorf("missing Content-Type. Supported content types application/openapi+json, applicatiion/openapi+yaml or application/wadl+xml")
	default:
		return search.Model{}, fmt.Errorf("content type '%s' not supported. Supported content types application/openapi+json, applicatiion/openapi+yaml or application/wadl+xml", contentType)
	}
}
