package fileparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"gtag/model"
)

func ParseFile(source string) ([]model.Struct, *token.FileSet, *ast.File, error) {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, source, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed parsing source file %v: %v", source, err)
	}

	return getStructs(file), fs, file, nil
}

func getStructs(file *ast.File) (res []model.Struct) {
	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			res = append(res,
				model.Struct{
					Fields: getFields(st),
					Node:   st,
				})
		}
	}
	return res
}

func getFields(st *ast.StructType) []model.Field {
	var fields []model.Field
	for _, field := range st.Fields.List {
		fields = append(fields, model.Field{
			Name: field.Names[0].Name,
			Tags: parseTags(field.Tag),
		})
	}
	return fields
}

func parseTags(bl *ast.BasicLit) model.Tags {

	var tagString string
	if bl != nil {
		tagString = bl.Value
	}
	return model.Tags{Str: tagString}
}
