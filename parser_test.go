package source

import (
	"go/ast"
	"go/token"
	"io"
	"reflect"
	"testing"

	"github.com/go-services/code"
)

func TestNewParser(t *testing.T) {
	tests := []struct {
		name string
		want Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_Parse(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		src io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			got, err := p.Parse(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_parseImports(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	tests := []struct {
		name        string
		fields      fields
		wantImports []Import
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			if gotImports := p.parseImports(); !reflect.DeepEqual(gotImports, tt.wantImports) {
				t.Errorf("fileParser.parseImports() = %v, want %v", gotImports, tt.wantImports)
			}
		})
	}
}

func Test_fileParser_isStructure(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d ast.Decl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			if got := p.isStructure(tt.args.d); got != tt.want {
				t.Errorf("fileParser.isStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_isInterface(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d ast.Decl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			if got := p.isInterface(tt.args.d); got != tt.want {
				t.Errorf("fileParser.isInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_isFunction(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d ast.Decl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			if got := p.isFunction(tt.args.d); got != tt.want {
				t.Errorf("fileParser.isFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_getType(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d ast.Decl
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   token.Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			if got := p.getType(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileParser.getType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_parseFunction(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d *ast.FuncDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Function
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			got, err := p.parseFunction(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.parseFunction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileParser.parseFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_parseStructure(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d *ast.GenDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Structure
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			got, err := p.parseStructure(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.parseStructure() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileParser.parseStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileParser_parseInterface(t *testing.T) {
	type fields struct {
		ast  *ast.File
		file *File
	}
	type args struct {
		d *ast.GenDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Interface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &fileParser{
				ast:  tt.fields.ast,
				file: tt.fields.file,
			}
			got, err := p.parseInterface(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("fileParser.parseInterface() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileParser.parseInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_functionParser_parseParams(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		params *ast.FieldList
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []code.Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &functionParser{
				imports: tt.fields.imports,
			}
			if got := f.parseParams(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("functionParser.parseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_functionParser_Parse(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		d *ast.FuncDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Function
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &functionParser{
				imports: tt.fields.imports,
			}
			got, err := f.Parse(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("functionParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("functionParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structParser_Parse(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		d *ast.GenDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Structure
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &structParser{
				imports: tt.fields.imports,
			}
			got, err := s.Parse(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("structParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interfaceParser_Parse(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		d *ast.GenDecl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Interface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &interfaceParser{
				imports: tt.fields.imports,
			}
			got, err := i.Parse(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("interfaceParser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interfaceParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interfaceParser_parseInterfaceMethods(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		methods *ast.FieldList
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []code.InterfaceMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &interfaceParser{
				imports: tt.fields.imports,
			}
			if got := i.parseInterfaceMethods(tt.args.methods); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("interfaceParser.parseInterfaceMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structParser_parseStructureFields(t *testing.T) {
	type fields struct {
		imports []Import
	}
	type args struct {
		fields *ast.FieldList
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []code.StructField
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &structParser{
				imports: tt.fields.imports,
			}
			if got := s.parseStructureFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structParser.parseStructureFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
