package source

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/go-services/annotation"
	"github.com/go-services/code"
)

func parseType(expr ast.Expr, imports []Import) code.Type {
	// try to find simple types that we can represent with code
	// if not use RawType to still be able to print the type
	tp := code.Type{}
	switch t := expr.(type) {
	case *ast.Ident:
		tp.Qualifier = t.Name
		return tp
	case *ast.SelectorExpr:
		qual, ok := t.X.(*ast.Ident)
		if !ok {
			tp.RawType = &jen.Statement{}
			parseComplexType(expr, tp.RawType)
			return tp
		}
		for _, i := range imports {
			if i.code.Alias == qual.Name {
				tp.Import = &i.code
				tp.Qualifier = t.Sel.Name
				return tp
			} else if i.pkg == qual.Name {
				tp.Import = &i.code
				tp.Import.Alias = i.pkg
				tp.Qualifier = t.Sel.Name
				return tp
			}
		}
		tp.RawType = &jen.Statement{}
		parseComplexType(expr, tp.RawType)
		return tp
	case *ast.StarExpr:
		tp = parseType(t.X, imports)
		if tp.RawType == nil {
			tp.Pointer = true
			return tp
		}
		tp.RawType = &jen.Statement{}
		parseComplexType(expr, tp.RawType)
		return tp
	case *ast.ArrayType:
		tp = parseType(t.Elt, imports)
		if tp.RawType == nil {
			tp.ArrayType = true
			return tp
		}
		tp.RawType = &jen.Statement{}
		parseComplexType(expr, tp.RawType)
		return tp
	case *ast.MapType:
		keyType := parseType(t.Key, imports)
		valueType := parseType(t.Value, imports)
		tp.MapType = &struct {
			Key   code.Type
			Value code.Type
		}{
			Key:   keyType,
			Value: valueType,
		}
		tp.RawType = &jen.Statement{}
		parseComplexType(expr, tp.RawType)
		return tp
	default:
		tp.RawType = &jen.Statement{}
		parseComplexType(expr, tp.RawType)
		return tp
	}
}
func parseComplexType(expr ast.Expr, statement *jen.Statement) {
	switch t := expr.(type) {
	case *ast.Ident:
		statement.Id(t.Name)
	case *ast.StarExpr:
		statement.Id("*")
		parseComplexType(t.X, statement)
	case *ast.SelectorExpr:
		qual := &jen.Statement{}
		parseComplexType(t.X, qual)
		statement.Add(qual).Dot(t.Sel.Name)
	case *ast.ArrayType:
		qual := &jen.Statement{}
		parseComplexType(t.Elt, qual)
		statement.Index().Add(qual)
	case *ast.MapType:
		key := &jen.Statement{}
		parseComplexType(t.Key, key)
		value := &jen.Statement{}
		parseComplexType(t.Value, value)
		statement.Map(key).Add(value)
	}
}

func cleanComment(comment string) string {
	comment = strings.TrimPrefix(comment, "//")
	comment = strings.TrimPrefix(comment, "/*")
	comment = strings.TrimSuffix(comment, "*/")
	comment = strings.TrimSpace(comment)
	return comment
}
func parseComments(docs *ast.CommentGroup) (dc []code.Comment) {
	if docs == nil {
		return dc
	}
	for _, c := range docs.List {
		if c == nil {
			continue
		}
		dc = append(dc, code.NewComment(cleanComment(c.Text)))
	}
	return
}

func annotate(c code.Code, force bool) ([]annotation.Annotation, error) {
	var annotations []annotation.Annotation
	for _, c := range c.Docs() {
		a, err := annotation.Parse(cleanComment(c.String()))
		if err != nil {
			if force {
				return nil, err
			}
			continue
		}
		annotations = append(annotations, *a)
	}
	return annotations, nil
}

// this is used to add code to the body of a code node.
// e.x to the body of a function, to the fields of a structure to the methods of an interface.
func appendCodeToInner(src string, node NodeWithInner, c code.Code) string {
	pre := strings.TrimRight(src[:node.InnerEnd()], "\n") + "\n"
	mid := ""
	lines := strings.Split(c.String(), "\n")
	for _, l := range lines {
		mid += "\t" + l + "\n"
	}
	end := src[node.InnerEnd():]
	return fmt.Sprintf(
		"%s%s%s",
		pre,
		mid,
		end,
	)
}
