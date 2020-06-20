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
	ApiStore ApiStore
}

func (v ApiView) ViewHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	res, err := v.ApiStore.Get(id)

	if err != nil {
		v.ErrorView(writer, fmt.Sprintf("Requested ID '%s' not found: %v", id, err))
	} else {
		api, err := v.decodeApiSpec(res.Source.Content)

		if err != nil {
			v.ErrorView(writer, fmt.Sprintf("Unable to decode '%s' specification %v", id, err))
		} else {
			templateFilePath := filepath.Join(v.Path, "templates", "view.html")
			t, err := template.New("view").ParseFiles(templateFilePath)

			if err != nil {
				v.ErrorView(writer, fmt.Sprintf("error reading template file %s %v", templateFilePath, err))
			}
			v.Title = "some sort of title"
			v.Id = id
			v.Api = api
			err = t.Execute(writer, v)
			if err != nil {
				fmt.Printf("Error rendering tempalte %v", err)
			}
		}
	}
}

func (v ApiView) ErrorView(writer http.ResponseWriter, message string) {
	templateFilePath := filepath.Join(v.Path, "templates", "error.html")
	t, err := template.New("error").ParseFiles(templateFilePath)

	if err != nil {
		log.Printf("error reading template file %s %v", templateFilePath, err)
	}

	err = t.Execute(writer, ErrorMessage{Message: message})
	if err != nil {
		fmt.Printf("Error rendering error tempalte %v", err)
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
