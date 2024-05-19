package source

import (
	"reflect"
	"testing"
)

type MockAnnotated struct {
	annotations []Annotation
}

func (m MockAnnotated) Annotate() error {
	panic("implement me")
}

func (m MockAnnotated) Annotations() []Annotation {
	return m.annotations
}

func TestFindAnnotations(t *testing.T) {
	type args struct {
		name      string
		annotated Annotated
	}
	tests := []struct {
		name            string
		args            args
		wantAnnotations []Annotation
	}{
		{
			name: "Should return the test annotation",
			args: args{
				name: "test",
				annotated: MockAnnotated{
					annotations: []Annotation{
						{
							Name: "abc",
						},
						{
							Name: "test",
						},
					},
				},
			},
			wantAnnotations: []Annotation{
				{
					Name: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnnotations := FindAnnotations(tt.args.name, tt.args.annotated); !reflect.DeepEqual(gotAnnotations, tt.wantAnnotations) {
				t.Errorf("FindAnnotations() = %v, want %v", gotAnnotations, tt.wantAnnotations)
			}
		})
	}
}
