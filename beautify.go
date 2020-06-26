package gtag

import (
	"fmt"
	"go/ast"
	"gtag/model"
	"sort"
	"strings"
)

type tagMap struct {
	MaxLength, Count int
}

func addSpace(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

// Beautify beautify the tag of given struct
func Beautify(str *model.Struct) {
	tagMapCount := make(map[string]tagMap)
	for _, field := range str.Fields {
		for _, tagKey := range field.Tags.GetKeys() {
			tagLength := len(field.Tags.Get(tagKey)) + len(tagKey)
			tm, ok := tagMapCount[tagKey]
			if !ok {
				tagMapCount[tagKey] = tagMap{
					MaxLength: tagLength,
					Count:     1,
				}
			} else {
				tmpTagCount := tagMap{
					MaxLength: tm.MaxLength,
					Count:     tm.Count + 1,
				}

				if tmpTagCount.MaxLength < tagLength {
					tmpTagCount.MaxLength = tagLength
				}
				tagMapCount[tagKey] = tmpTagCount
			}
		}
	}

	sortedKey := sortByCount(tagMapCount)
	for _, field := range str.Fields {
		var newTag string
		for _, key := range sortedKey {
			if tagValue := field.Tags.Get(key); tagValue != "" {
				spaces := tagMapCount[key].MaxLength - (len(key + tagValue)) + 1
				newTag += fmt.Sprintf("%s:\"%s\"%s", key, tagValue, addSpace(spaces))
			}
		}
		writeToNode(field.Tags.Node, newTag)
	}
}

func writeToNode(node ast.Node, newTag string) {
	if newTag == "" {
		return
	}

	tag, ok := node.(*ast.BasicLit)
	if !ok {
		return
	}
	tag.Value = fmt.Sprintf("`%s`", strings.TrimRight(newTag, " "))
}

func sortByCount(tagMapCount map[string]tagMap) (keys []string) {
	type kv struct {
		Key   string
		Value tagMap
	}

	var ss []kv
	for k, v := range tagMapCount {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value.Count > ss[j].Value.Count
	})

	for _, kv := range ss {
		keys = append(keys, kv.Key)
	}
	return
}
