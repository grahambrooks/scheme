package main

import (
	"net/http"
)

func (s *SchemeServer) ListApisHandler(writer http.ResponseWriter, _ *http.Request) {
	res, err := s.ApiStore.List()
	mirrorResponse(res, err, writer)
}
