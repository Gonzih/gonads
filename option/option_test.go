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

func TestMapSome(t *testing.T) {
	o := Some(10)
	o2 := Map(o, func(i int) int {
		return i + 2
	})

	assert.True(t, o.Some())
	assert.True(t, o2.Some())
	v, err := o2.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, 12, *v)
}

func TestMapNone(t *testing.T) {
	o := None[int]()
	o2 := Map(o, func(i int) int {
		return i + 2
	})

	assert.True(t, o.None())
	assert.True(t, o2.None())
	v, err := o2.Unwrap()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}
