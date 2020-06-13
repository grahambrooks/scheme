package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type ApiView struct {
	Path     string
	Id       string
	Title    string
	Api      interface{}
	ApiStore *ApiStore
}

func (v ApiView) ViewHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	res, err := v.ApiStore.Get(id)

	api, err := v.decodeApiSpec(res.Source.Content)

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
