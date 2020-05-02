package search

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModelToEncoding(t *testing.T) {
	t.Run("JSON Typed Model", func(t *testing.T) {
		model := Model{Kind: OpenAPI2}
		var buffer bytes.Buffer
		assert.NoError(t, model.AsJson(&buffer))
		assert.Equal(t,
			`{"kind":"OpenAPI 2.0"}
`,
			buffer.String())
	})
	t.Run("JSON Full Model", func(t *testing.T) {
		model := Model{
			Kind:        OpenAPI2,
			Title:       "An API",
			Description: "An API Description",
			Version:     "2020.1.2",
			Schemes:     nil,
		}
		var buffer bytes.Buffer
		assert.NoError(t, model.AsJson(&buffer))
		assert.Equal(t,
			`{"kind":"OpenAPI 2.0","title":"An API","description":"An API Description","version":"2020.1.2"}
`,
			buffer.String())
	})
}
