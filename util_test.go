package source

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/go-services/code"
)

func Test_parseType(t *testing.T) {
	type args struct {
		expr    ast.Expr
		imports []Import
	}
	tests := []struct {
		name string
		args args
		want code.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseType(tt.args.expr, tt.args.imports); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseComplexType(t *testing.T) {
	type args struct {
		expr      ast.Expr
		statement *jen.Statement
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseComplexType(tt.args.expr, tt.args.statement)
		})
	}
}

func Test_cleanComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanComment(tt.args.comment); got != tt.want {
				t.Errorf("cleanComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseComments(t *testing.T) {
	type args struct {
		docs *ast.CommentGroup
	}
	tests := []struct {
		name   string
		args   args
		wantDc []code.Comment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDc := parseComments(tt.args.docs); !reflect.DeepEqual(gotDc, tt.wantDc) {
				t.Errorf("parseComments() = %v, want %v", gotDc, tt.wantDc)
			}
		})
	}
}
