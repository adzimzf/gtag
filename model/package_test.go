package model

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestTags_GetKeys(t1 *testing.T) {
	type fields struct {
		Str       string
		Node      ast.Node
		StrResult string
	}
	tests := []struct {
		name     string
		fields   fields
		wantKeys []string
	}{
		{
			fields: fields{
				Str:       `gorm:"date"json:"date"`,
				StrResult: "",
			},
			wantKeys: []string{"gorm", "json"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tags{
				Str:       tt.fields.Str,
				Node:      tt.fields.Node,
				StrResult: tt.fields.StrResult,
			}
			if gotKeys := t.GetKeys(); !reflect.DeepEqual(gotKeys, tt.wantKeys) {
				t1.Errorf("GetKeys() = %v, want %v", gotKeys, tt.wantKeys)
			}
		})
	}
}

func TestTags_IsValid(t1 *testing.T) {
	type fields struct {
		Str       string
		Node      ast.Node
		StrResult string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				Str: `db:"name"`,
			},
			want: true,
		},
		{
			fields: fields{
				Str: `db:"name"json:"name"`,
			},
			want: true,
		},
		{
			fields: fields{
				Str: `db"name"json:"name"`,
			},
			want: false,
		},
		{
			fields: fields{
				Str: `db:name"json:"name"`,
			},
			want: false,
		},
		{
			fields: fields{
				Str: `db:"name"json`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Tags{
				Str:       tt.fields.Str,
				Node:      tt.fields.Node,
				StrResult: tt.fields.StrResult,
			}
			if got := t.IsValid(); got != tt.want {
				t1.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
