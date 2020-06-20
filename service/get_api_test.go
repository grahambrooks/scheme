package main

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIHandingAPIs(t *testing.T) {
	stubbedStore := StubApiStore{}
	server := SchemeServer{Port: 8000, ApiStore: &stubbedStore}
	t.Run("Missing API id in request", func(t *testing.T) {
		assert.HTTPError(t, server.GetApiHandler, http.MethodGet, "/apis", nil, `{"Message":"Invalid document id (missing)"}`)
		assert.HTTPBodyContains(t, server.GetApiHandler, http.MethodGet, "/apis", nil, `{"Message":"Invalid document id (missing)"}`)
	})

	t.Run("Store Error", func(t *testing.T) {
		stubbedStore.ErrorResponse = fmt.Errorf("testing error")

		req, _ := http.NewRequest(http.MethodGet, "/apis/123", nil)

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/apis/{id}", server.GetApiHandler)
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Contains(t, rr.Body.String(), `"Error reading API specification: testing error`)
	})

	t.Run("Valid get API request", func(t *testing.T) {
		stubbedStore.ErrorResponse = nil
		stubbedStore.GetResponse.Source.Content = "Yep that's me"

		req, _ := http.NewRequest(http.MethodGet, "/apis/123", nil)

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/apis/{id}", server.GetApiHandler)
		router.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK)
		assert.Contains(t, rr.Body.String(), `Yep that's me`)
	})

	t.Run("Add API Specification", func(t *testing.T) {
		stubbedStore.ErrorResponse = nil
		stubbedStore.PutRequestId = "api:id"
		stubbedStore.PutRequestBody = `{"test":"response"}`
		stubbedStore.PutResponse = &esapi.Response{
			StatusCode: 200,
			Header:     nil,
			Body:       responseReader(`{ "response": "test"}`),
		}

		request := TestRequest{
			Method:  http.MethodPost,
			Url:     "/apis/123",
			Headers: map[string]string{"Content-Type": "application/openapi+json"},
			Body:    `{ "swagger": "2.0" }`,
		}
		RouteHTTPBodyContains(t, "/apis/{id}", server.NewApiHandler, request, http.StatusOK, `{"response":"test"}`)
	})

}

func TestAPIRegistration(t *testing.T) {
	stubbedStore := StubApiStore{}
	server := SchemeServer{Port: 8000, ApiStore: &stubbedStore}

	t.Run("Register API with no request body", func(t *testing.T) {
		request := TestRequest{
			Method:  http.MethodPost,
			Url:     "/registrations",
			Headers: map[string]string{"Content-Type": "application/openapi+json"},
			Body:    `{}`,
		}
		RouteHTTPBodyContains(t, "/registrations", server.NewRegistration, request, http.StatusBadRequest, `"ID missing in registration request"`)
	})

	t.Run("Register API with ID but no URL", func(t *testing.T) {
		request := TestRequest{
			Method:  http.MethodPost,
			Url:     "/registrations",
			Headers: map[string]string{"Content-Type": "application/openapi+json"},
			Body:    `{ "id": "api:id" }`,
		}
		RouteHTTPBodyContains(t, "/registrations", server.NewRegistration, request, http.StatusBadRequest, `"URL missing in registration request"`)
	})

	t.Run("Register API with ID but no URL", func(t *testing.T) {
		request := TestRequest{
			Method:  http.MethodPost,
			Url:     "/registrations",
			Headers: map[string]string{"Content-Type": "application/openapi+json"},
			Body: `{
  "id": "api:id",
  "url": "http://some/url"
}`,
		}
		RouteHTTPBodyContains(t, "/registrations", server.NewRegistration, request, http.StatusBadRequest, `"Error requested API specification Get \"http://some/url\"`)
	})
}
