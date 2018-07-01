package source

import (
	"github.com/go-services/annotation"
	"github.com/go-services/code"
	"go/ast"
)

type Annotated interface {
	Annotate() error
}

type Node interface {
	Code() code.Code
	Name() string
	Begin() int
	End() int
	InnerBegin() int
	InnerEnd() int
}

// Import represents an import
type Import struct {
	ast ast.Decl
	// code representation of import
	code code.Import
	// the package name of the import
	// e.x some/import/path does not guaranty the package name to be `path` we need
	// a way to get that package name when using import in types
	pkg string
}

// Structure represents a parsed structure
type Structure struct {
	ast ast.Decl
	// code representation of the struct
	code code.Struct

	// annotations of the struct
	annotations []annotation.Annotation

	// the beginning and positions of the struct definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and positions of the struct brackets
	// this is used to determine where to put structure fields
	innerBegin, innerEnd int
}

// Interface represents a parsed interface
type Interface struct {
	ast ast.Decl

	// code representation of the interface
	code code.Interface

	// annotations of the interface
	annotations []annotation.Annotation

	// the beginning and positions of the interface definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and positions of the interface brackets
	// this is used to determine where to put interface methods
	innerBegin, innerEnd int
}

type Function struct {
	ast ast.Decl

	// code representation of the function
	code code.Function

	// annotations of the interface
	annotations []annotation.Annotation

	// the beginning and positions of the function definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and positions of the function brackets
	// this is used to determine where to put function code
	innerBegin, innerEnd int
}

// File represents a parsed file.
type File struct {
	pkg        string
	src        string
	imports    []Import
	structures []Structure
	interfaces []Interface
	functions  []Function
}

func (s *Structure) Name() string {
	return s.code.Name
}

func (s *Structure) Begin() int {
	return s.begin
}

func (s *Structure) End() int {
	return s.end
}

func (s *Structure) InnerBegin() int {
	return s.innerBegin
}

func (s *Structure) InnerEnd() int {
	return s.innerEnd
}

func (s *Structure) Code() code.Code {
	return &s.code
}

func (s *Structure) Struct() code.Struct {
	return s.code
}

func (s *Structure) Annotate() error {
	for _, c := range s.code.Docs() {
		a, err := annotation.Parse(cleanComment(c.String()))
		if err != nil {
			return err
		}
		s.annotations = append(s.annotations, *a)
	}
	return nil
}

func (i *Interface) Name() string {
	return i.code.Name
}

func (i *Interface) Begin() int {
	return i.begin
}

func (i *Interface) End() int {
	return i.end
}

func (i *Interface) InnerBegin() int {
	return i.innerBegin
}

func (i *Interface) InnerEnd() int {
	return i.innerEnd
}

func (i *Interface) Code() code.Code {
	return &i.code
}

func (i *Interface) Interface() code.Interface {
	return i.code
}

func (i *Interface) Annotate() error {
	for _, c := range i.code.Docs() {
		a, err := annotation.Parse(cleanComment(c.String()))
		if err != nil {
			return err
		}
		i.annotations = append(i.annotations, *a)
	}
	return nil
}

func (f *Function) Name() string {
	return f.code.Name
}

func (f *Function) Begin() int {
	return f.begin
}

func (f *Function) End() int {
	return f.end
}

func (f *Function) InnerBegin() int {
	return f.innerBegin
}

func (f *Function) InnerEnd() int {
	return f.innerEnd
}

func (f *Function) Code() code.Code {
	return &f.code
}

func (f *Function) Function() code.Function {
	return f.code
}

func (f *Function) Annotate() error {
	for _, c := range f.code.Docs() {
		a, err := annotation.Parse(cleanComment(c.String()))
		if err != nil {
			return err
		}
		f.annotations = append(f.annotations, *a)
	}
	return nil
}
