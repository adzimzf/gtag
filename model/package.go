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

func (t Tags) Get(key string) string {
	return reflect.StructTag(strings.Trim(t.Str, "`")).Get(key)
}

// IsValid check whether the validity of struct tag
func (t Tags) IsValid() bool {
	tag := t.Str
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && (tag[i] == ' ' || tag[i] == '`') {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

	}
	return false
}

// GetKeys get the list keys of a struct tag
// the algorithm is not originally mine
// I copied from go/src/reflect/type.go:1136
// then modify some lines
func (t Tags) GetKeys() (keys []string) {
	// don't change the initial value
	tag := t.Str
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && (tag[i] == ' ' || tag[i] == '`') {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		keys = append(keys, tag[:i])
		tag = tag[i+1:]

		// Find the value then remove it.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		tag = tag[i+1:]
	}
	return
}
