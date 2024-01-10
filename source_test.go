package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSource(t *testing.T) {
	src, err := New(`
package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type XYZ struct {
	sm []*string
}

type ABC struct {
	sm []XYZ
}

func TestNewSource(t *testing.T) {
}
`)
	assert.NoError(t, err)
	assert.NotNil(t, src)
}
