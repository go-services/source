package source

import (
	"fmt"
	"strings"
)

type Annotation struct {
	Name string
	Args []string
}

func (a Annotation) String() string {
	return fmt.Sprintf("gs:%s %s", a.Name, strings.Join(a.Args, " "))
}

// Annotated is the interface all the annotated nodes need to satisfy
type Annotated interface {
	Annotate() error
	Annotations() []Annotation
}

// FindAnnotations finds an annotation by name given an annotated node.
func FindAnnotations(name string, annotated Annotated) (annotations []Annotation) {
	for _, ann := range annotated.Annotations() {
		if ann.Name == name {
			annotations = append(annotations, ann)
		}
	}
	return
}
