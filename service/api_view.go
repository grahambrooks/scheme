package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type ApiView struct {
	Path  string
	Id    string
	Title string
	Api   interface{}
}

func (v ApiView) readApiSpec(id string) string {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.GetRequest{
		Index:      DocumentIndexName,
		DocumentID: id,
	}
	res, _ := req.Do(context.Background(), es)

	decoder := json.NewDecoder(res.Body)
	var elasticResponse ElasticGetResponse
	_ = decoder.Decode(&elasticResponse)

	return elasticResponse.Source.Content
}

func (v ApiView) ViewHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	spec := v.readApiSpec(id)

	api, err := v.decodeApiSpec(spec)

	templateFilePath := filepath.Join(v.Path, "templates", "view.html")
	t, err := template.New("view").ParseFiles(templateFilePath)

	if err != nil {
		log.Printf("error reading template file %s %v", templateFilePath, err)
	}
	v.Title = "some sort of title"
	v.Id = id
	v.Api = api
	err = t.Execute(writer, v)
	if err != nil {
		fmt.Printf("Error rendering tempalte %v", err)
	}
}

func (v ApiView) decodeApiSpec(spec string) (interface{}, error) {
	decoder := json.NewDecoder(strings.NewReader(spec))

	var api interface{}
	err := decoder.Decode(&api)
	if err != nil {
		decoder := yaml.NewDecoder(strings.NewReader(spec))
		err = decoder.Decode(&api)
		if err != nil {
			log.Printf("error reading API %v", err)
		}
	}
	return api, err
}
