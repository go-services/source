package source

import "github.com/go-services/annotation"

// Annotated is the interface all the annotated nodes need to satisfy
type Annotated interface {
	Annotate(force bool) error
	Annotations() []annotation.Annotation
}

// FindAnnotations finds an annotation by name given an annotated node.
func FindAnnotations(name string, annotated Annotated) (annotations []annotation.Annotation) {
	for _, ann := range annotated.Annotations() {
		if ann.Name == name {
			annotations = append(annotations, ann)
		}
	}
	return
}
