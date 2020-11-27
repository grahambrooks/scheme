package rules

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
	"strings"
	"testing"
)

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

		for key := range j.(map[string]interface{}) {
			fmt.Println("YAML Key type", reflect.TypeOf(key))
			switch reflect.TypeOf(key).Kind() {
			case reflect.String:
				fmt.Println("Key value", j.(map[string]interface{})[key])
			}
		}
	})
}
