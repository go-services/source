package source

import (
	"go/build"
	"os"
)

type ImportInfo struct {
	Dir  string
	Name string
}
type BuildContext interface {
	Import(path string) (*ImportInfo, error)
	Cwd() (string, error)
}

type DefaultBuildContext struct{}

func (d DefaultBuildContext) Import(path string) (*ImportInfo, error) {
	cwd, err := d.Cwd()
	if err != nil {
		return nil, err
	}
	pkg, err := build.Import(path, cwd, 0)
	if err != nil {
		return nil, err
	}
	return &ImportInfo{
		Dir:  pkg.Dir,
		Name: pkg.Name,
	}, nil
}

func (d DefaultBuildContext) Cwd() (string, error) {
	cwd, err := os.Getwd()
	return cwd, err
}
