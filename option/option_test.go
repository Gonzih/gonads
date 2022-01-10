package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	o := Some("test")

	assert.True(t, o.Some())
	assert.False(t, o.None())
}

func TestNone(t *testing.T) {
	o := None[string]()

	assert.False(t, o.Some())
	assert.True(t, o.None())
}
