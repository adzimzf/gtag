package model

import (
	"go/ast"
	"reflect"
	"strings"
)

// Struct is data type for struct type
type Struct struct {
	Fields    []Field
	MaxLenTag map[string]int
	Node      ast.Node
}

type Field struct {
	Name string
	Tags Tags
}

type Tags struct {
	// Str is the actual tag without parsing
	Str       string
	Node      ast.Node
	StrResult string
}

func (t Tags) Parse() []string {
	return strings.Split(strings.Trim(t.Str, "`"), " ")

}

func (t Tags) Len() int {
	return len(strings.Split(strings.Trim(t.Str, "`"), " "))
}

func (t Tags) Get(key string) string {
	return reflect.StructTag(strings.Trim(t.Str, "`")).Get(key)
}

func (t Tags) GetKeys() (keys []string) {
	for _, tag := range strings.Split(strings.Trim(t.Str, "`"), " ") {
		if k := strings.Split(tag, ":")[0]; k != "" {
			keys = append(keys, k)
		}
	}
	return
}

type TagResult struct {
	FiledName string
	Values    map[string]TagValue
}
type TagValue struct {
	Val         string
	Len, MaxLen int
}
