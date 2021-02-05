package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	o := Some("test")

	assert.True(t, o.IsSome())
	assert.False(t, o.IsNone())
}

func TestNone(t *testing.T) {
	o := None()

	assert.False(t, o.IsSome())
	assert.True(t, o.IsNone())
}
