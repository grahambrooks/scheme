package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"strings"
)

const (
	SearchIndexName   = "apellicon-search"
	DocumentIndexName = "apellicon-docs"
)

type ElasticApiStore struct {
}

func NewApiStore() *ElasticApiStore {
	return &ElasticApiStore{}
}

type ElasticGetResponse struct {
	Id     string `json:"_id"`
	Source struct {
		Content string `json:"Content"`
	} `json:"_source"`
}

func (store *ElasticApiStore) Get(id string) (ElasticGetResponse, error) {
	es, _ := elasticsearch.NewDefaultClient()
	req := esapi.GetRequest{
		Index:      DocumentIndexName,
		DocumentID: id,
	}
	res, _ := req.Do(context.Background(), es)

	decoder := json.NewDecoder(res.Body)
	var elasticResponse ElasticGetResponse
	err := decoder.Decode(&elasticResponse)

	return elasticResponse, err
}

func (store *ElasticApiStore) TextSearch(filter string) (*esapi.Response, error) {
	query := fmt.Sprintf(`{
  "query": {
    "query_string": {
      "query": "%s"
    }
  }
}`, filter)
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	req := esapi.SearchRequest{
		Index: []string{SearchIndexName},
		Body:  strings.NewReader(query),
	}
	return req.Do(context.Background(), es)
}

func (store *ElasticApiStore) List() (*esapi.Response, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	req := esapi.SearchRequest{
		Index: []string{SearchIndexName},
	}

	return req.Do(context.Background(), es)
}

func (store *ElasticApiStore) Save(id string, content string) (*esapi.Response, error) {

	entry := ApiDocumentEntry{
		Content: content,
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)

	_ = encoder.Encode(entry)
	es, _ := elasticsearch.NewDefaultClient()

	req := esapi.IndexRequest{
		Index:      DocumentIndexName,
		DocumentID: id,
		Body:       bytes.NewReader(buffer.Bytes()),
		Refresh:    "true",
		ErrorTrace: true,
	}

	return req.Do(context.Background(), es)
}

func (store*ElasticApiStore) IndexDocument(id string, content []byte) (*esapi.Response, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	req := esapi.IndexRequest{
		Index:      SearchIndexName,
		DocumentID: id,
		Body:       bytes.NewReader(content),
		Refresh:    "true",
		ErrorTrace: true,
	}

	return req.Do(context.Background(), es)
}