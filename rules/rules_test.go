package rules

import (
	"github.com/grahambrooks/scheme/openapi"
	"github.com/grahambrooks/scheme/search"
	"github.com/stretchr/testify/assert"
	"github.com/yalp/jsonpath"
	"os"
	"testing"
)

func LoadAPIYaml(path string) (interface{}, error) {
	parser := openapi.Parser{}
	f, err := os.Open(path)
	if err != nil {
		return search.Model{}, err
	}

	return parser.ParseRawYaml(f)
}

func LoadAPIJson(path string) (interface{}, error) {
	parser := openapi.Parser{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return parser.ParseRawJson(f)
}

type TestSpec struct {
	name     string
	spec     RuleSpec
	data     func() interface{}
	expected bool
}

func Test(t *testing.T) {
	tests := []TestSpec{
		{
			name: "Missing with exists",
			spec: RuleSpec{
				Description: "Existence",
				Path:        "$.missing",
				Test:        "exists",
			},
			data: func() interface{} {
				return make(map[string]interface{})
			},
			expected: false,
		},
		{
			name: "Set with exists",
			spec: RuleSpec{
				Description: "Existence",
				Path:        "$.missing",
				Test:        "exists",
			},
			data: func() interface{} {
				m := make(map[string]interface{})
				m["missing"] = "I am here"
				return m
			},
			expected: true,
		},
		{
			name: "Successful value compare",
			spec: RuleSpec{
				Description: "Successful value compare",
				Path:        "$.value",
				Test:        "I am here",
			},
			data: func() interface{} {
				m := make(map[string]interface{})
				m["value"] = "I am here"
				return m
			},
			expected: true,
		},
		{
			name: "Unsuccessful  value compare",
			spec: RuleSpec{
				Description: "Successful value compare",
				Path:        "$.value",
				Test:        "I over there",
			},
			data: func() interface{} {
				m := make(map[string]interface{})
				m["value"] = "I am here"
				return m
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rule := CompileRule(test.spec)

			assert.Equal(t, test.expected, rule.Test(test.data()))
		})
	}
}

func TestCompilingRules(t *testing.T) {

	t.Run("Simple", func(t *testing.T) {
		spec := RuleSpec{
			Description: "Existence",
			Path:        "$.missing",
			Test:        "exists",
		}

		rule := CompileRule(spec)

		assert.False(t, rule.Test(make(map[string]interface{})))
	})

	t.Run("Simple does exist", func(t *testing.T) {
		spec := RuleSpec{
			Description: "Existence",
			Path:        "$.missing",
			Test:        "exists",
		}

		rule := CompileRule(spec)

		v := make(map[string]interface{})
		v["missing"] = true
		assert.True(t, rule.Test(v))
	})
}

func TestRules(t *testing.T) {
	t.Run("Existence", func(t *testing.T) {
		model, err := LoadAPIYaml("testdata/complete.yaml")

		assert.NoError(t, err)
		assert.NotNil(t, model)

		paths, err := jsonpath.Read(model, "$.paths.*")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$.paths")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$[\"host\"]")
		assert.NoError(t, err)
		assert.Equal(t, "api.uber.com", paths.(string))

		paths, err = jsonpath.Read(model, "$..responses.*")
		assert.NoError(t, err)
		assert.NotNil(t, paths)
	})

	t.Run("JSON Existence", func(t *testing.T) {
		model, err := LoadAPIJson("testdata/complete.json")

		assert.NoError(t, err)
		assert.NotNil(t, model)

		paths, err := jsonpath.Read(model, "$.paths.*")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$.paths")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$.host")
		assert.NoError(t, err)
		assert.NotNil(t, paths)

		paths, err = jsonpath.Read(model, "$..responses.*")
		assert.NoError(t, err)
		assert.NotNil(t, paths)
	})
}
