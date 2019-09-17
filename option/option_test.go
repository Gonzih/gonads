package option

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	o := Some("test")

	assert.True(t, o.Some())
	assert.False(t, o.None())
}

func TestNone(t *testing.T) {
	o := None()

	assert.False(t, o.Some())
	assert.True(t, o.None())
}

func TestMapNone(t *testing.T) {
	o := None()
	o2 := Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	}, o)

	assert.True(t, o2.None())
}

func TestMapSome(t *testing.T) {
	o := Some("world")
	o2 := Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	}, o)

	assert.True(t, o2.Some())
	assert.Equal(t, o2.Unwrap(), "hello world")
}

func getName(name F) Option {
	if name == "potato" {
		return Some("Potato")
	}

	return None()
}

func TestFMapSome(t *testing.T) {
	o := Some("potato")
	o2 := FMap(getName, o)

	assert.True(t, o2.Some())
	assert.Equal(t, o2.Unwrap(), "Potato")
}

func TestFMapSomeWithMissingVal(t *testing.T) {
	o := Some("not potato")
	o2 := FMap(getName, o)

	assert.True(t, o2.None())
}

func TestFMapNone(t *testing.T) {
	o := None()
	o2 := FMap(getName, o)

	assert.True(t, o2.None())
}
