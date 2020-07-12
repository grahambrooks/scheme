package rules

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
	"testing"
)


func NormalizeStructure(from interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	//t := reflect.TypeOf(from).
	//
	//switch {
	//case t.Implements(map[interface{}]interface{}):
	//}
	//
	//
	//kt := t.Key()
	//t.
	//
	//kind := t.Kind()

	return result
}

func TestNormalization(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		normalized := NormalizeStructure(nil)

		assert.Len(t, normalized, 0)
	})

	t.Run("map[interface{}]interface{} normalization", func(t *testing.T) {
		from := make(map[interface{}]interface{})
		from["key"] = 1

		//normalized := NormalizeStructure(from)
		//
		//assert.Len(t, normalized, 1)
	})

}

func TestStructuredTraversal(t *testing.T) {
	t.Run("Empty JSON", func(t *testing.T) {
		emptyJson := `{}`

		decoder := json.NewDecoder(strings.NewReader(emptyJson))
		var j interface{}
		_ = decoder.Decode(&j)

		tj := reflect.TypeOf(j)

		fmt.Println(tj)
	})
	t.Run("JSON KeyValue", func(t *testing.T) {
		emptyJson := `{
  "key": 1
}`

		decoder := json.NewDecoder(strings.NewReader(emptyJson))
		var j interface{}
		_ = decoder.Decode(&j)

		tj := reflect.TypeOf(j)

		fmt.Println(tj)
	})

	t.Run("YAML Key Value", func(t *testing.T) {
		emptyYaml := `key: 1`

		decoder := yaml.NewDecoder(strings.NewReader(emptyYaml))
		var j interface{}
		_ = decoder.Decode(&j)

		tj := reflect.TypeOf(j)

		fmt.Println(tj)

		for key := range j.(map[interface{}]interface{}) {
			fmt.Println("YAML Key type", reflect.TypeOf(key))
			switch reflect.TypeOf(key).Kind() {
			case reflect.String:
				fmt.Println("Key value", j.(map[interface{}]interface{})[key.(string)])
			}
		}
	})
}
