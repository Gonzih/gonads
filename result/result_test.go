package result

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	r := Ok("test")

	assert.True(t, r.IsOk())
	assert.False(t, r.IsErr())
	assert.Equal(t, "test", r.Ok())
}

func TestErr(t *testing.T) {
	e := errors.New("error message")
	r := Err(e)

	assert.False(t, r.IsOk())
	assert.True(t, r.IsErr())
	assert.Equal(t, e, r.Err())
}

func TestMapOk(t *testing.T) {
	r := Ok("test")

	r2 := Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	}, r)

	assert.True(t, r2.IsOk())
	assert.Equal(t, "hello test", r2.Ok())
}

func TestMapErr(t *testing.T) {
	e := errors.New("error message")
	r := Err(e)

	r2 := Map(func(s F) T {
		return fmt.Sprintf("hello %v", s)
	}, r)

	assert.False(t, r2.IsOk())
	assert.Equal(t, e, r2.Err())
}

func lookupName(name F) Result {
	if name == "potato" {
		return Ok("Potato")
	}

	return Err(errors.New("missing name"))
}

func TestFMapOk(t *testing.T) {
	r := Ok("potato")

	r2 := FMap(lookupName, r)

	assert.True(t, r2.IsOk())
	assert.Equal(t, "Potato", r2.Ok())
}

func TestFMapOkWrongName(t *testing.T) {
	r := Ok("noone")

	r2 := FMap(lookupName, r)

	assert.False(t, r2.IsOk())
}

func TestFMapErr(t *testing.T) {
	e := errors.New("error message")
	r := Err(e)

	r2 := FMap(lookupName, r)

	assert.False(t, r2.IsOk())
	assert.Equal(t, e, r2.Err())
}