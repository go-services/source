package source

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-errors/errors"
	"github.com/go-services/code"
)

type fileParser struct {
	ast  *ast.File
	file *file
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

func newParser() *fileParser {
	return &fileParser{}
}

func (p *fileParser) parse(src string) (*file, error) {
	// parse the source
	fSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fSet, "file.go", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// store ast representation in file
	p.ast = astFile

	// parse package
	if p.ast.Name == nil {
		return nil, errors.New("no package found")
	}

	// store source in file
	p.file = newFile(p.ast.Name.Name, src, p.ast)
	p.file.imports = p.parseImports()

	// parse code nodes
	for _, d := range p.ast.Decls {
		switch p.getType(d) {
		case token.STRUCT:
			structures, err := p.parseStructure(d.(*ast.GenDecl))
			if err != nil {
				return nil, err
			}
			for _, s := range structures {
				p.file.structures[s.Name()] = s
			}
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
			p.file.functions[function.Name()] = function
		case token.INTERFACE:
			ifc, err := p.parseInterface(d.(*ast.GenDecl))
			if err != nil {
				return nil, err
			}
			p.file.interfaces[ifc.Name()] = ifc
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
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		pkg, err := build.Import(imp.code.Path, dir, 0)
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
func (p *fileParser) parseStructure(d *ast.GenDecl) ([]Structure, error) {
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
	ft.paramBegin = int(d.Type.Params.Opening)
	ft.paramEnd = int(d.Type.Params.Closing) - 1
	ft.code = *code.NewFunction(
		d.Name.Name,
		code.ParamsFunctionOption(
			f.parseParams(d.Type.Params)...,
		),
		code.ResultsFunctionOption(
			f.parseParams(d.Type.Results)...,
		),
	)
	ft.exported = ast.IsExported(d.Name.Name)
	if d.Recv != nil && len(d.Recv.List) > 0 {
		ft.code.Recv = &f.parseParams(d.Recv)[0]
	}
	return ft, ft.Annotate(false)
}
func (s *structParser) Parse(d *ast.GenDecl) ([]Structure, error) {
	var stcs []Structure
	// we do not need to test this because it should never come to this point if
	// the declaration is not a structure
	// here we get the type so we have the name
	for _, v := range d.Specs {

		tp := v.(*ast.TypeSpec)
		st := Structure{
			ast:    d,
			fields: []StructureField{},
			begin:  int(tp.Pos()) - 1,
			end:    int(tp.End()) - 1,
		}
		// get fields if any
		fields := tp.Type.(*ast.StructType).Fields
		if fields != nil {
			st.innerBegin = int(fields.Opening)
			st.innerEnd = int(fields.Closing) - 1
		}
		// create the code representation
		sf, sfl := s.parseStructureFields(fields)
		st.code = *code.NewStructWithFields(
			tp.Name.Name,
			sf,
			parseComments(d.Doc)...,
		)
		st.exported = ast.IsExported(tp.Name.Name)
		st.fields = sfl
		err := st.Annotate(false)
		if err != nil {
			return nil, err
		}
		stcs = append(stcs, st)
	}

	return stcs, nil
}
func (i *interfaceParser) Parse(d *ast.GenDecl) (Interface, error) {
	inf := Interface{
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
		inf.innerBegin = int(methods.Opening)
		inf.innerEnd = int(methods.Closing) - 1
	}
	// create the code representation
	im, ims := i.parseInterfaceMethods(methods)
	inf.code = *code.NewInterface(
		tp.Name.Name,
		im,
		parseComments(d.Doc)...,
	)
	inf.exported = ast.IsExported(tp.Name.Name)
	inf.methods = ims
	return inf, inf.Annotate(false)
}
func (i *interfaceParser) parseInterfaceMethods(methods *ast.FieldList) ([]code.InterfaceMethod, []InterfaceMethod) {
	var list []code.InterfaceMethod
	var sList []InterfaceMethod
	if methods == nil {
		return list, sList
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
			ims := InterfaceMethod{
				code:  im,
				begin: int(f.Pos()) - 1,
				end:   int(f.End()) - 1,
			}
			ims.exported = ast.IsExported(n.Name)

			_ = ims.Annotate(false)
			list = append(list, im)
			sList = append(sList, ims)
		}
	}
	return list, sList
}
func (s *structParser) parseStructureFields(fields *ast.FieldList) ([]code.StructField, []StructureField) {
	var list []code.StructField
	var sList []StructureField
	if fields == nil {
		return list, sList
	}
	for _, f := range fields.List {
		if f == nil {
			continue
		}

		if len(f.Names) == 0 {
			sf := code.NewStructField("", parseType(f.Type, s.imports), parseComments(f.Doc)...)
			if f.Tag != nil && f.Tag.Kind == token.STRING {
				// remove ` before parsing
				sf.Tags = parseTags(f.Tag.Value[1 : len(f.Tag.Value)-1])
			}
			list = append(list, *sf)
			stf := StructureField{
				code:  *sf,
				begin: int(f.Pos()) - 1,
				end:   int(f.End()) - 1,
			}
			_ = stf.Annotate(false)
			sList = append(sList, stf)
			continue
		}
		for _, n := range f.Names {
			sf := code.NewStructField(n.Name, parseType(f.Type, s.imports), parseComments(f.Doc)...)
			if f.Tag != nil && f.Tag.Kind == token.STRING {
				// remove ` before parsing
				sf.Tags = parseTags(f.Tag.Value[1 : len(f.Tag.Value)-1])
			}
			list = append(list, *sf)
			stf := StructureField{
				code:  *sf,
				begin: int(f.Pos()) - 1,
				end:   int(f.End()) - 1,
			}
			stf.exported = ast.IsExported(n.Name)
			_ = stf.Annotate(false)
			sList = append(sList, stf)
		}
	}
	return list, sList

}

// this is copied and modified from https://golang.org/src/reflect/type.go?s=31821:31842#L1174
func parseTags(tag string) *code.FieldTags {
	tags := code.FieldTags{}
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// Scan to colon. A space, a quote or a control character is a syntax error.
		// Strictly speaking, control chars include the range [0x7f, 0x9f], not just
		// [0x00, 0x1f], but in practice, we ignore the multi-byte control characters
		// as it is simpler to inspect the tag's bytes than the tag's runes.
		i = 0
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := tag[:i]
		tag = tag[i+1:]

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		quotedValue := tag[:i+1]
		value, err := strconv.Unquote(quotedValue)
		if err != nil {
			break
		}
		tags[name] = value
		tag = tag[i+1:]
	}
	return &tags
}
