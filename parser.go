package source

import (
	"bytes"
	"github.com/go-errors/errors"
	"github.com/go-services/code"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io"
	"strconv"
	"strings"
)

type Parser interface {
	Parse(src io.Reader) (*File, error)
}

type fileParser struct {
	ast  *ast.File
	file *File
}
type structParser struct {
	imports []Import
}
type functionParser struct {
	imports []Import
}
type interfaceParser struct {
	imports []Import
}

func NewParser() Parser {
	return &fileParser{}
}

func (p *fileParser) Parse(src io.Reader) (*File, error) {
	// get the data
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, src); err != nil {
		return nil, err
	}

	// store source in file
	p.file = &File{
		src: string(buf.Bytes()),
	}

	// parse the source
	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, "file.go", p.file.src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// store ast representation in file
	p.ast = astFile

	// parse package
	if p.ast.Name == nil {
		return nil, errors.New("no package found")
	}

	p.file.pkg = p.ast.Name.Name
	p.file.imports = p.parseImports()

	// parse code nodes
	for _, d := range p.ast.Decls {
		switch p.getType(d) {
		case token.STRUCT:
			structure, err := p.parseStructure(d.(*ast.GenDecl))
			if err != nil {
				return nil, err
			}
			p.file.structures = append(p.file.structures, structure)
		case token.FUNC:
			function, err := p.parseFunction(d.(*ast.FuncDecl))
			if err != nil {
				return nil, err
			}
			// copy body from the source and add it to the function.
			// this is just a simple way to get the body when using code.Function
			innerBody := p.file.src[function.InnerBegin():function.InnerEnd()]
			function.code.AddStringBody(strings.TrimSpace(innerBody))

			// add the function
			p.file.functions = append(p.file.functions, function)
		case token.INTERFACE:
			ifc, err := p.parseInterface(d.(*ast.GenDecl))
			if err != nil {
				return nil, err
			}
			p.file.interfaces = append(p.file.interfaces, ifc)
		}
	}
	return p.file, nil
}
func (p *fileParser) parseImports() (imports []Import) {
	// find imports
	for _, i := range p.ast.Imports {
		imp := Import{
			code: code.Import{},
		}
		if i.Name != nil {
			imp.code.Alias = i.Name.Name
		}

		pth, err := strconv.Unquote(i.Path.Value)
		if err == nil {
			imp.code.Path = pth
		}

		pkg, err := build.Import(imp.code.Path, ".", 0)
		if err == nil {
			imp.code.FilePath = pkg.Dir
			imp.pkg = pkg.Name
		}
		imports = append(imports, imp)
	}
	return imports

}

func (p *fileParser) isStructure(d ast.Decl) bool {
	gDecl, ok := d.(*ast.GenDecl)
	if !ok || gDecl.Tok != token.TYPE {
		return false
	}
	if len(gDecl.Specs) != 1 {
		return false
	}
	tp, ok := gDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return false
	}
	if tp.Name == nil || tp.Name.Name == "" {
		return false
	}
	if _, ok := tp.Type.(*ast.StructType); !ok {
		return false
	}
	return true
}

func (p *fileParser) isInterface(d ast.Decl) bool {
	gDecl, ok := d.(*ast.GenDecl)
	if !ok || gDecl.Tok != token.TYPE {
		return false
	}
	if len(gDecl.Specs) != 1 {
		return false
	}
	tp, ok := gDecl.Specs[0].(*ast.TypeSpec)
	if !ok {
		return false
	}
	if tp.Name == nil || tp.Name.Name == "" {
		return false
	}
	if _, ok := tp.Type.(*ast.InterfaceType); !ok {
		return false
	}
	return true
}

func (p *fileParser) isFunction(d ast.Decl) bool {
	gDecl, ok := d.(*ast.FuncDecl)
	if !ok {
		return false
	}
	if gDecl.Name == nil || gDecl.Name.Name == "" {
		return false
	}
	return true
}

func (p *fileParser) getType(d ast.Decl) token.Token {
	if p.isStructure(d) {
		return token.STRUCT
	} else if p.isFunction(d) {
		return token.FUNC
	} else if p.isInterface(d) {
		return token.INTERFACE
	}
	return token.ILLEGAL
}

func (p *fileParser) parseFunction(d *ast.FuncDecl) (Function, error) {
	fp := &functionParser{
		imports: p.file.imports,
	}
	return fp.Parse(d)

}
func (p *fileParser) parseStructure(d *ast.GenDecl) (Structure, error) {
	sp := &structParser{
		imports: p.file.imports,
	}
	return sp.Parse(d)
}

func (p *fileParser) parseInterface(d *ast.GenDecl) (Interface, error) {
	ip := &interfaceParser{
		imports: p.file.imports,
	}
	return ip.Parse(d)
}

func (f *functionParser) parseParams(params *ast.FieldList) []code.Parameter {
	var list []code.Parameter
	if params == nil {
		return list
	}
	for _, p := range params.List {
		if p == nil {
			continue
		}
		if len(p.Names) == 0 {
			prm := code.NewParameter("", parseType(p.Type, f.imports))
			list = append(list, *prm)
			continue
		}
		for _, n := range p.Names {
			prm := code.NewParameter(n.Name, parseType(p.Type, f.imports))
			list = append(list, *prm)
		}
	}
	return list
}
func (f *functionParser) Parse(d *ast.FuncDecl) (Function, error) {
	ft := Function{
		ast:   d,
		begin: int(d.Pos()) - 1,
		end:   int(d.End()) - 1,
	}
	if d.Body != nil {
		ft.innerBegin = int(d.Body.Lbrace)
		ft.innerEnd = int(d.Body.Rbrace) - 1
	}
	ft.code = *code.NewFunction(
		d.Name.Name,
		code.ParamsFunctionOption(
			f.parseParams(d.Type.Params)...,
		),
		code.ResultsFunctionOption(
			f.parseParams(d.Type.Results)...,
		),
	)
	if d.Recv != nil && len(d.Recv.List) > 0 {
		ft.code.Recv = &f.parseParams(d.Recv)[0]
	}
	return ft, nil
}
func (s *structParser) Parse(d *ast.GenDecl) (Structure, error) {
	st := Structure{
		ast:   d,
		begin: int(d.Pos()) - 1,
		end:   int(d.End()) - 1,
	}

	// we do not need to test this because it should never come to this point if
	// the declaration is not a structure
	// here we get the type so we have the name
	tp := d.Specs[0].(*ast.TypeSpec)

	// get fields if any
	fields := tp.Type.(*ast.StructType).Fields
	if fields != nil {
		st.innerBegin = int(fields.Opening)
		st.innerEnd = int(fields.Closing) - 1
	}
	// create the code representation
	st.code = *code.NewStructWithFields(
		tp.Name.Name,
		s.parseStructureFields(fields),
		parseComments(d.Doc)...,
	)
	return st, nil
}
func (i *interfaceParser) Parse(d *ast.GenDecl) (Interface, error) {
	st := Interface{
		ast:   d,
		begin: int(d.Pos()) - 1,
		end:   int(d.End()) - 1,
	}

	// we do not need to test this because it should never come to this point if
	// the declaration is not a structure
	// here we get the type so we have the name
	tp := d.Specs[0].(*ast.TypeSpec)

	// get fields if any
	methods := tp.Type.(*ast.InterfaceType).Methods
	if methods != nil {
		st.innerBegin = int(methods.Opening)
		st.innerEnd = int(methods.Closing) - 1
	}
	// create the code representation
	st.code = *code.NewInterface(
		tp.Name.Name,
		i.parseInterfaceMethods(methods),
		parseComments(d.Doc)...,
	)
	return st, nil
}
func (i *interfaceParser) parseInterfaceMethods(methods *ast.FieldList) []code.InterfaceMethod {
	var list []code.InterfaceMethod
	if methods == nil {
		return list
	}
	for _, f := range methods.List {
		tp, ok := f.Type.(*ast.FuncType)
		if !ok {
			continue
		}
		if f == nil {
			continue
		}
		for _, n := range f.Names {
			mp := functionParser{
				imports: i.imports,
			}
			mth, err := mp.Parse(&ast.FuncDecl{
				Name: n,
				Type: tp,
			})
			if err != nil {
				continue
			}
			im := code.NewInterfaceMethod(
				mth.code.Name,
				code.ParamsFunctionOption(mth.code.Params...),
				code.ResultsFunctionOption(mth.code.Results...),
				code.DocsFunctionOption(parseComments(f.Doc)...),
			)
			list = append(list, im)
		}
	}
	return list
}
func (s *structParser) parseStructureFields(fields *ast.FieldList) []code.StructField {
	var list []code.StructField
	if fields == nil {
		return list
	}
	for _, f := range fields.List {
		if f == nil {
			continue
		}

		if len(f.Names) == 0 {
			sf := code.NewStructField("", parseType(f.Type, s.imports), parseComments(f.Doc)...)
			list = append(list, *sf)
			continue
		}
		for _, n := range f.Names {
			sf := code.NewStructField(n.Name, parseType(f.Type, s.imports), parseComments(f.Doc)...)
			list = append(list, *sf)
		}
	}
	return list

}
