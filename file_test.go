package source

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/go-services/annotation"
	"github.com/go-services/code"
)

func TestStructure_Name(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.Name(); got != tt.want {
				t.Errorf("Structure.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_Begin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.Begin(); got != tt.want {
				t.Errorf("Structure.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_End(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.End(); got != tt.want {
				t.Errorf("Structure.End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_InnerBegin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.InnerBegin(); got != tt.want {
				t.Errorf("Structure.InnerBegin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_InnerEnd(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.InnerEnd(); got != tt.want {
				t.Errorf("Structure.InnerEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_Code(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Code
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.Code(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Structure.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_Struct(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Struct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := s.Struct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Structure.Struct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructure_Annotate(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Struct
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Structure{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if err := s.Annotate(); (err != nil) != tt.wantErr {
				t.Errorf("Structure.Annotate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInterface_Name(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.Name(); got != tt.want {
				t.Errorf("Interface.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_Begin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.Begin(); got != tt.want {
				t.Errorf("Interface.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_End(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.End(); got != tt.want {
				t.Errorf("Interface.End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_InnerBegin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.InnerBegin(); got != tt.want {
				t.Errorf("Interface.InnerBegin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_InnerEnd(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.InnerEnd(); got != tt.want {
				t.Errorf("Interface.InnerEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_Code(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Code
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.Code(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_Interface(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Interface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := i.Interface(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interface.Interface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterface_Annotate(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Interface
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interface{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if err := i.Annotate(); (err != nil) != tt.wantErr {
				t.Errorf("Interface.Annotate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFunction_Name(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.Name(); got != tt.want {
				t.Errorf("Function.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Begin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.Begin(); got != tt.want {
				t.Errorf("Function.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_End(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.End(); got != tt.want {
				t.Errorf("Function.End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_InnerBegin(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.InnerBegin(); got != tt.want {
				t.Errorf("Function.InnerBegin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_InnerEnd(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.InnerEnd(); got != tt.want {
				t.Errorf("Function.InnerEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Code(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Code
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.Code(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Function(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Function
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if got := f.Function(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Function.Function() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFunction_Annotate(t *testing.T) {
	type fields struct {
		ast         ast.Decl
		code        code.Function
		annotations []annotation.Annotation
		begin       int
		end         int
		innerBegin  int
		innerEnd    int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				ast:         tt.fields.ast,
				code:        tt.fields.code,
				annotations: tt.fields.annotations,
				begin:       tt.fields.begin,
				end:         tt.fields.end,
				innerBegin:  tt.fields.innerBegin,
				innerEnd:    tt.fields.innerEnd,
			}
			if err := f.Annotate(); (err != nil) != tt.wantErr {
				t.Errorf("Function.Annotate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
