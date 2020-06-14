package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiRegistrationRequest struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type ApiRegistration struct {
	Id      string                 `json:"id"`
	Url     string                 `json:"url"`
	Request ApiRegistrationRequest `json:"request"`
}

func (s *ApelliconServer) NewRegistration(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var registration ApiRegistration
	err := decoder.Decode(&registration)

	if len(registration.Id) == 0 {
		errorResponse(writer, "ID missing in registration request")
	} else {
		if len(registration.Url) == 0 {
			errorResponse(writer, "URL missing in registration request")
		}
	}

	if err != nil {
		errorResponse(writer, fmt.Sprintf("Unable to decode request %v", err))
	} else {
		client := &http.Client{}

		req, err := http.NewRequest(registration.Request.Method, registration.Url, strings.NewReader(registration.Request.Body))

		if err != nil {
			errorResponse(writer, fmt.Sprintf("Error building requested API specification request %v", err))
		} else {
			for k, v := range registration.Request.Headers {
				req.Header.Add(k, v)
			}

			resp, err := client.Do(req)
			if err != nil {
				errorResponse(writer, fmt.Sprintf("Error requested API specification %v", err))
			} else {
				content, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					errorResponse(writer, fmt.Sprintf("Error reading API response body %v", err))
				} else {
					saved, err := s.ApiStore.Save(registration.Id, string(content))
					mirrorResponse(saved, err, writer)
				}
			}
		}
	}
}
