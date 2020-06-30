package fileparser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"

	"github.com/adzimzf/gtag"
	"github.com/adzimzf/gtag/model"
)

type FileParser struct {
	fileSet *token.FileSet
	file    *ast.File
	absSrc  string
	cfg     *gtag.Config
}

func NewFileParser(cfg *gtag.Config) (*FileParser, error) {
	absSrc, err := filepath.Abs(cfg.FileSource)
	if err != nil {
		return nil, err
	}
	fp := &FileParser{
		fileSet: token.NewFileSet(),
		cfg:     cfg,
		absSrc:  absSrc,
	}
	fp.file, err = parser.ParseFile(fp.fileSet, fp.absSrc, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return fp, nil
}

func (f *FileParser) Write() error {
	var buf bytes.Buffer
	err := format.Node(&buf, f.fileSet, f.file)
	if err != nil {
		return err
	}
	if !f.cfg.IsOverWrite {
		fmt.Println(buf.String())
		return nil
	}
	return ioutil.WriteFile(f.absSrc, buf.Bytes(), 0)
}

func (f *FileParser) FindStructs() (res []*model.Struct) {
	for _, decl := range f.file.Decls {
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
				&model.Struct{
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

		var fieldName string
		if len(field.Names) > 0 {
			fieldName = field.Names[0].Name
		} else {
			idn, ok := field.Type.(*ast.Ident)
			if ok {
				fieldName = idn.Name
			}
		}

		fields = append(fields, model.Field{
			Name: fieldName,
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
	return model.Tags{Str: tagString, Node: bl}
}
