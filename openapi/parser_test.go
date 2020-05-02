package openapi

import (
	"apellicon/search"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadingJson(t *testing.T) {
	t.Run("JSON Uber Spec", func(t *testing.T) {
		content, err := ioutil.ReadFile("test/json/uber.json")
		assert.NoError(t, err)

		assert.True(t, len(string(content)) > 0)

		f, err := os.Open("test/json/uber.json")
		assert.NoError(t, err)

		parser := Parser{}
		model, err := parser.ParseJson(f)
		assert.NoError(t, err)
		assert.Equal(t, search.OpenAPI2, model.Kind)
		assert.Equal(t, "Uber API", model.Title)
		assert.True(t, len(model.Description) > 0)

		assert.Equal(t, 5, len(model.Resources))
	})

	t.Run("YAML Uber Spec", func(t *testing.T) {
		content, err := ioutil.ReadFile("test/yaml/uber.yaml")
		assert.NoError(t, err)

		assert.True(t, len(string(content)) > 0)

		f, err := os.Open("test/json/uber.json")
		assert.NoError(t, err)

		parser := Parser{}
		model, err := parser.ParseYaml(f)
		assert.NoError(t, err)
		assert.Equal(t, search.OpenAPI2, model.Kind)
		assert.Equal(t, "Uber API", model.Title)
		assert.True(t, len(model.Description) > 0)

		assert.Equal(t, 5, len(model.Resources))
	})

	t.Run("YAML 3 UPS To You Spec", func(t *testing.T) {
		testfilename := "test/v3.0/uspto.yaml"
		content, err := ioutil.ReadFile(testfilename)
		assert.NoError(t, err)

		assert.True(t, len(string(content)) > 0)

		f, err := os.Open(testfilename)
		assert.NoError(t, err)

		parser := Parser{}
		model, err := parser.ParseYaml(f)
		assert.NoError(t, err)
		assert.Equal(t, search.OpenAPI3, model.Kind)
		assert.Equal(t, "USPTO Data Set API", model.Title)
		assert.True(t, len(model.Description) > 0)

		assert.Equal(t, 3, len(model.Resources))
	})
}
