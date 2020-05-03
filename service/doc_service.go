package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type ApiDocumentEntry struct {
	Content string
}

func WriteApiEntry(id string, content string) {

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

	_, _ = req.Do(context.Background(), es)
}
