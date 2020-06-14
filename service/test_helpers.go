package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func responseReader(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

type TestRequest struct {
	Method  string
	Url     string
	Headers map[string]string
	Body    string
}

func RouteHTTPBodyContains(t *testing.T, path string, handler http.HandlerFunc, request TestRequest, expectedStatusCode int, expectedResponse interface{}) {
	req, _ := http.NewRequest(request.Method, request.Url, strings.NewReader(request.Body))

	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc(path, handler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, expectedStatusCode)
	assert.Contains(t, rr.Body.String(), expectedResponse)
}
