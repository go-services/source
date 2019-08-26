package source

import "github.com/go-services/annotation"

type Annotated interface {
	Annotate(force bool) error
	Annotations() []annotation.Annotation
}

func FindAnnotations(name string, annotated Annotated) (annotations []annotation.Annotation) {
	for _, ann := range annotated.Annotations() {
		if ann.Name == name {
			annotations = append(annotations, ann)
		}
	}
	return
}
