package rules

import (
	"github.com/yalp/jsonpath"
	"reflect"
	"strings"
)

type RuleSpec struct {
	Description string
	Path        string
	Test        string
}

type Rule struct {
	Description string
	Test        func(interface{}) bool
}

func CompileRule(spec RuleSpec) Rule {
	testFn := CompileTest(spec.Test)
	return Rule{
		Description: spec.Description,
		Test: func(o interface{}) bool {
			paths, err := jsonpath.Read(o, spec.Path)

			if err != nil {
				return false
			}

			return testFn(paths)
		},
	}
}

func CompileTest(test string) func(interface{}) bool {
	return func(value interface{}) bool {
		switch {
		case strings.EqualFold(test, "exists"):
			return value != nil
		case strings.EqualFold(test, "!exists"):
			return value == nil
		default:
			return reflect.DeepEqual(value, test)
		}
	}
}
