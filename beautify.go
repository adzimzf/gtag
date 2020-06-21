package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"gtats/model"
	"sort"
)

func BeautifyTag(str *model.Struct) {
	var tagResults []model.TagResult
	maxLen := make(map[string]int)
	for _, field := range str.Fields {
		tagResult := model.TagResult{
			FiledName: field.Name,
			Values:    make(map[string]model.TagValue),
		}

		for _, key := range field.Tags.GetKeys() {
			val := field.Tags.Get(key)
			tagResult.Values[key] = model.TagValue{
				Val: val,
				Len: len(val),
			}

			//
			currentLen := len(key) + len(val) + 3
			m, ok := maxLen[key]
			if !ok {
				maxLen[key] = currentLen
			} else {
				if m < currentLen {
					maxLen[key] = currentLen
				}
			}

		}
		if field.Tags.Len() != 0 {
			tagResults = append(tagResults, tagResult)
		}
	}
	writeTag(str, tagResults, maxLen)
}

func writeTag(str *model.Struct, tagResults []model.TagResult, maxLen map[string]int) {
	for _, n2 := range tagResults {
		if len(n2.Values) == 0 {
			continue
		}
		st, ok := str.Node.(*ast.StructType)
		if ok {
			for _, i2 := range st.Fields.List {
				if i2.Names[0].Name == n2.FiledName {
					bTag := generateTagResult(n2, maxLen)
					if i2.Tag != nil {
						i2.Tag.Value = bTag
					}
				}

			}
		}
	}
}

func generateTagResult(tr model.TagResult, mLen map[string]int) string {
	buf := bytes.NewBufferString("`")
	for i, trKey := range sortedKey(tr.Values) {
		var spaceLen int
		if len(tr.Values) > i+1 {
			spaceLen = mLen[trKey] - (len(trKey) + tr.Values[trKey].Len + 3) + 1
		}
		buf.WriteString(fmt.Sprintf("%s:\"%s\"%s", trKey, tr.Values[trKey].Val, addSpace(spaceLen)))
	}
	buf.WriteString("`")
	return buf.String()
}

func sortedKey(tv map[string]model.TagValue) []string {
	keys := make([]string, 0, len(tv))
	for k := range tv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func addSpace(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}
