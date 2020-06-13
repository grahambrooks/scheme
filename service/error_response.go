package main

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string
}

func errorResponse(writer http.ResponseWriter, s string) {
	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(ErrorMessage{Message: s})
	writer.WriteHeader(http.StatusBadRequest)

}
