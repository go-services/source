package source

import (
	"github.com/go-services/code"
	"fmt"
	"go/format"
)

type Source struct {
	file   *File
	parser *fileParser
}

func New(src string) (*Source, error) {
	p := newParser()
	f, err := p.parse(src)
	if err != nil {
		return nil, err
	}
	return &Source{
		file:   f,
		parser: p,
	}, nil
}
func (s *Source) AddFieldToStruct(name string, field *code.StructField) error {
	structure, ok := s.file.structures[name]
	if !ok {
		return fmt.Errorf("structure with name %s not found", name)
	}
	s.file.src = appendCodeToInner(s.file.src, &structure, field)
	f, err := s.parser.parse(s.file.src)
	if err != nil {
		return err
	}
	s.file = f
	return nil
}

func (s *Source) Code() (string, error) {
	src, err := format.Source([]byte(s.file.src))
	return string(src), err
}
