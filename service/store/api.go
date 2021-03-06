package store

import "github.com/elastic/go-elasticsearch/v7/esapi"

type ApiStore interface {
	Get(id string) (ElasticGetResponse, error)
	IndexDocument(id string, content []byte) (*esapi.Response, error)
	TextSearch(filter string) (*esapi.Response, error)
	List() (*esapi.Response, error)
	Save(id string, content string) (*esapi.Response, error)
}
