package source

import (
	"github.com/go-services/annotation"
	"github.com/go-services/code"
	"go/ast"
)

type Annotated interface {
	Annotate(force bool) error
	Annotations(force bool) error
}

type Node interface {
	Exported() bool
	Code() code.Code
	String() string
	Name() string
	Begin() int
	End() int
}

type NodeWithInner interface {
	Node

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

type StructureField struct {
	exported bool
	// code representation of the struct field
	code code.StructField

	// annotations of the struct field
	annotations []annotation.Annotation

	// the beginning and end positions of the struct field definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int
}

// Structure represents a parsed structure
type Structure struct {
	exported bool
	ast      ast.Decl
	// code representation of the struct
	code code.Struct

	// the structure fields
	fields []StructureField

	// annotations of the struct
	annotations []annotation.Annotation

	// the beginning and end positions of the struct definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and end positions of the struct brackets
	// this is used to determine where to put structure fields
	innerBegin, innerEnd int
}

type InterfaceMethod struct {
	exported bool
	// code representation of the interface method
	code code.InterfaceMethod

	// annotations of the interface method
	annotations []annotation.Annotation

	// the beginning and end positions of the interface method definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int
}

// Interface represents a parsed interface
type Interface struct {
	exported bool
	ast      ast.Decl

	// code representation of the interface
	code code.Interface

	// the interface methods
	methods []InterfaceMethod

	// annotations of the interface
	annotations []annotation.Annotation

	// the beginning and end positions of the interface definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and end positions of the interface brackets
	// this is used to determine where to put interface methods
	innerBegin, innerEnd int
}

type Function struct {
	exported bool
	ast      ast.Decl

	// code representation of the function
	code code.Function

	// annotations of the interface
	annotations []annotation.Annotation

	// the beginning and end positions of the function definition
	// corresponds to the Pos() and End() of the ast declaration
	begin, end int

	// the beginning and end positions of the function brackets
	// this is used to determine where to put function code
	innerBegin, innerEnd int

	// the  beginning and end positions of the function parameters
	paramBegin, paramEnd int

	// the  beginning and end positions of the function parameters
	resultBegin, resultEnd int
}

// File represents a parsed file.
type File struct {
	pkg        string
	src        string
	imports    []Import
	structures map[string]Structure
	interfaces map[string]Interface
	functions  map[string]Function
}

func NewFile(pkg, src string) *File {
	return &File{
		pkg:        pkg,
		src:        src,
		structures: map[string]Structure{},
		interfaces: map[string]Interface{},
		functions:  map[string]Function{},
	}
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

func (s *Structure) String() string {
	return s.code.String()
}
func (s *Structure) Fields() []StructureField {
	return s.fields
}

func (s *Structure) Annotations() []annotation.Annotation {
	return s.annotations
}

func (s *Structure) Annotate(force bool) error {
	a, err := annotate(&s.code, force)
	if a != nil {
		s.annotations = append(s.annotations, *a)
	}
	return err
}
func (s *Structure) Exported() bool {
	return s.exported
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

func (i *Interface) String() string {
	return i.code.String()
}

func (i *Interface) Methods() []InterfaceMethod {
	return i.methods
}

func (i *Interface) Annotations() []annotation.Annotation {
	return i.annotations
}

func (i *Interface) Annotate(force bool) error {
	a, err := annotate(&i.code, force)
	if a != nil {
		i.annotations = append(i.annotations, *a)
	}
	return err
}
func (i *Interface) Exported() bool {
	return i.exported
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

func (f *Function) ParamBegin() int {
	return f.paramBegin
}

func (f *Function) ParamEnd() int {
	return f.paramEnd
}

func (f *Function) Code() code.Code {
	return &f.code
}

func (f *Function) String() string {
	return f.code.String()
}

func (f *Function) Params() []code.Parameter {
	return f.code.Params
}

func (f *Function) Results() []code.Parameter {
	return f.code.Results
}

func (f *Function) Receiver() *code.Parameter {
	return f.code.Recv
}

func (f *Function) Annotations() []annotation.Annotation {
	return f.annotations
}

func (f *Function) Annotate(force bool) error {
	a, err := annotate(&f.code, force)
	if a != nil {
		f.annotations = append(f.annotations, *a)
	}
	return err
}
func (f *Function) Exported() bool {
	return f.exported
}

func (f *StructureField) Name() string {
	return f.code.Name
}

func (f *StructureField) String() string {
	return f.code.String()
}

func (f *StructureField) Code() code.Code {
	return &f.code
}
func (f *StructureField) Tags() map[string]string {
	return *f.code.Tags
}

func (f *StructureField) Begin() int {
	return f.begin
}

func (f *StructureField) End() int {
	return f.end
}
func (f *StructureField) Annotations() []annotation.Annotation {
	return f.annotations
}

func (f *StructureField) Annotate(force bool) error {
	a, err := annotate(&f.code, force)
	if a != nil {
		f.annotations = append(f.annotations, *a)
	}
	return err
}
func (f *StructureField) Exported() bool {
	return f.exported
}

func (f *InterfaceMethod) Name() string {
	return f.code.Name
}
func (f *InterfaceMethod) String() string {
	return f.code.String()
}

func (f *InterfaceMethod) Code() code.Code {
	return &f.code
}

func (f *InterfaceMethod) Params() []code.Parameter {
	return f.code.Params
}

func (f *InterfaceMethod) Results() []code.Parameter {
	return f.code.Results
}

func (f *InterfaceMethod) Annotations() []annotation.Annotation {
	return f.annotations
}

func (f *InterfaceMethod) Annotate(force bool) error {
	a, err := annotate(&f.code, force)
	if a != nil {
		f.annotations = append(f.annotations, *a)
	}
	return err
}

func (f *InterfaceMethod) Begin() int {
	return f.begin
}

func (f *InterfaceMethod) End() int {
	return f.end
}
func (f *InterfaceMethod) Exported() bool {
	return f.exported
}
