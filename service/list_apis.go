package main

import (
	"net/http"
)

func (s *SchemeServer) ListApisHandler(writer http.ResponseWriter, request *http.Request) {
	res, err := s.ApiStore.List()
	mirrorResponse(res, err, writer)
}
