package source

import "testing"

func TestParser(t *testing.T) {
	src, err := New(`
package source

//abc:test 123
type XYZ struct {
	sm []*string
}`)
	if err != nil {
		t.Fatal(err)
	}
	if src == nil {
		t.Fatal("source is nil")
	}

	if len(src.Structures()) != 1 {
		t.Fatalf("expected 1 struct, got %d", len(src.Structures()))
	}

}
