package main

import "github.com/elastic/go-elasticsearch/esapi"

type StubApiStore struct {
	GetResponse    ElasticGetResponse
	ErrorResponse  error
	PutRequestId   string
	PutRequestBody string
	PutResponse    *esapi.Response
}

func (s *StubApiStore) Get(id string) (ElasticGetResponse, error) {
	s.GetResponse.Id = id
	return s.GetResponse, s.ErrorResponse
}

func (s *StubApiStore) TextSearch(filter string) (*esapi.Response, error) {
	panic("implement me")
}

func (s *StubApiStore) List() (*esapi.Response, error) {
	panic("implement me")
}

func (s *StubApiStore) Save(id string, content string) (*esapi.Response, error) {
	s.PutRequestId = id
	s.PutRequestBody = content

	return nil, s.ErrorResponse
}

func (s *StubApiStore) IndexDocument(id string, content []byte) (*esapi.Response, error) {
	return s.PutResponse, s.ErrorResponse
}
