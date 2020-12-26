package server

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string
}

func errorResponse(writer http.ResponseWriter, s string) {
	writer.WriteHeader(http.StatusBadRequest)
	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(ErrorMessage{Message: s})
}
