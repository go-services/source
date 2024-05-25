package source

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"

	"github.com/go-errors/errors"
	"github.com/go-services/code"
)

type Options struct {
	buildContext BuildContext
}

type Option func(*Options)

func WithBuildContext(buildContext BuildContext) Option {
	return func(o *Options) {
		o.buildContext = buildContext
	}
}

type fileParser struct {
	ast          *ast.File
	file         *file
	buildContext BuildContext
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

func newParser(opts ...Option) *fileParser {
	options := Options{
		buildContext: DefaultBuildContext{},
	}
	for _, o := range opts {
		o(&options)
	}
	return &fileParser{
		buildContext: options.buildContext,
	}
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
		case token.TYPE:
			err := p.parseTypeSpec(d.(*ast.GenDecl))
			if err != nil {
				return nil, err
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
		pkg, err := p.buildContext.Import(imp.code.Path)
		if err == nil {
			imp.code.FilePath = pkg.Dir
			imp.pkg = pkg.Name
		}
		imports = append(imports, imp)
	}
	return imports
}

func (p *fileParser) isType(d ast.Decl) bool {
	gDecl, ok := d.(*ast.GenDecl)
	if !ok || gDecl.Tok != token.TYPE {
		return false
	}
	return true
}

func (p *fileParser) isStructure(spec ast.Spec) bool {
	tp, ok := spec.(*ast.TypeSpec)
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

func (p *fileParser) isInterface(spec ast.Spec) bool {
	tp, ok := spec.(*ast.TypeSpec)
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

func (p *fileParser) getType(spec ast.Decl) token.Token {
	if p.isType(spec) {
		return token.TYPE
	} else if p.isFunction(spec) {
		return token.FUNC
	}
	return token.ILLEGAL
}

func (p *fileParser) parseTypeSpec(d *ast.GenDecl) error {
	for _, spec := range d.Specs {
		tp := spec.(*ast.TypeSpec)
		if len(d.Specs) == 1 {
			tp.Doc = d.Doc
		}
		if p.isInterface(spec) {
			ifc, err := p.parseInterface(tp)
			if err != nil {
				return err
			}
			p.file.interfaces[ifc.Name()] = ifc
		} else if p.isStructure(spec) {
			structures, err := p.parseStructure(tp)
			if err != nil {
				return err
			}
			p.file.structures[structures.Name()] = structures
		}
	}
	return nil
}

func (p *fileParser) parseFunction(d *ast.FuncDecl) (Function, error) {
	fp := &functionParser{
		imports: p.file.imports,
	}
	return fp.Parse(d)
}

func (p *fileParser) parseStructure(spec *ast.TypeSpec) (Structure, error) {
	sp := &structParser{
		imports: p.file.imports,
	}
	return sp.Parse(spec)
}

func (p *fileParser) parseInterface(spec *ast.TypeSpec) (Interface, error) {
	ip := &interfaceParser{
		imports: p.file.imports,
	}
	return ip.Parse(spec)
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

		tp := parseType(p.Type, f.imports)
		if tp == nil {
			// type not supported
			continue
		}
		if len(p.Names) == 0 {
			prm := code.NewParameter("", *tp)
			list = append(list, *prm)
			continue
		}
		for _, n := range p.Names {
			prm := code.NewParameter(n.Name, *tp)
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
		code.DocsFunctionOption(parseComments(d.Doc)...),
	)
	ft.exported = ast.IsExported(d.Name.Name)
	if d.Recv != nil && len(d.Recv.List) > 0 {
		ft.code.Recv = &f.parseParams(d.Recv)[0]
	}
	return ft, nil
}

func (s *structParser) Parse(tp *ast.TypeSpec) (Structure, error) {
	st := Structure{
		ast:    tp,
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
		parseComments(tp.Doc)...,
	)
	st.exported = ast.IsExported(tp.Name.Name)
	st.fields = sfl
	return st, nil
}

func (i *interfaceParser) Parse(tp *ast.TypeSpec) (Interface, error) {
	inf := Interface{
		ast:   tp,
		begin: int(tp.Pos()) - 1,
		end:   int(tp.End()) - 1,
	}

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
		parseComments(tp.Doc)...,
	)
	inf.exported = ast.IsExported(tp.Name.Name)
	inf.methods = ims
	return inf, nil
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
		tp := parseType(f.Type, s.imports)
		if tp == nil {
			// type not supported
			continue
		}
		if len(f.Names) == 0 {
			sf := code.NewStructField("", *tp, parseComments(f.Doc)...)
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
			sList = append(sList, stf)
			continue
		}
		for _, n := range f.Names {
			sf := code.NewStructField(n.Name, *tp, parseComments(f.Doc)...)
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
