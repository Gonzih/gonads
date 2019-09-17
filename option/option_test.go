package option

import (
	"fmt"
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

func TestMapNone(t *testing.T) {
	o := None()
	o2 := o.Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	})

	assert.True(t, o2.IsNone())
}

func TestMapSome(t *testing.T) {
	o := Some("world")
	o2 := o.Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	})

	assert.True(t, o2.IsSome())
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
	o2 := o.FMap(getName)

	assert.True(t, o2.IsSome())
	assert.Equal(t, o2.Unwrap(), "Potato")
}

func TestFMapSomeWithMissingVal(t *testing.T) {
	o := Some("not potato")
	o2 := o.FMap(getName)

	assert.True(t, o2.IsNone())
}

func TestFMapNone(t *testing.T) {
	o := None()
	o2 := o.FMap(getName)

	assert.True(t, o2.IsNone())
}
