package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *ApelliconServer) GetApiHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if len(id) == 0 {
		errorResponse(writer, "Invalid document id (missing)")
	} else {
		elasticResponse, err := s.ApiStore.Get(id)

		if err != nil {
			errorResponse(writer, fmt.Sprintf("Error reading API specification %v", err))
		}

		_, _ = writer.Write([]byte(elasticResponse.Source.Content))
	}
}
