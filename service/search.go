package main

import (
	"fmt"
	"net/http"
)

func (s *SchemeServer) SearchApiHandler(writer http.ResponseWriter, request *http.Request) {
	filter := request.FormValue("query")

	if len(filter) > 0 {
		res, err := s.ApiStore.TextSearch(filter)
		if err != nil {
			errorResponse(writer, fmt.Sprintf("Error connecting to API store (elastic) %v", err))
		} else {
			mirrorResponse(res, err, writer)

		}
	} else {
		s.ListApisHandler(writer, request)
	}
}
