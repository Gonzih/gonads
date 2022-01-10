package result

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	r := Ok("test")

	assert.True(t, r.Ok())
	assert.False(t, r.Err())
	v, err := r.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, "test", *v)
}

func TestErr(t *testing.T) {
	e := errors.New("error message")
	r := Err[string](e)

	assert.False(t, r.Ok())
	assert.True(t, r.Err())
	v, err := r.Unwrap()
	assert.Nil(t, v)
	assert.Equal(t, e, err)
}

func TestMapOk(t *testing.T) {
	r := Ok(10)
	r2 := Map(r, func(i int) int {
		return i + 2
	})

	assert.True(t, r.Ok())
	assert.True(t, r2.Ok())
	v, err := r2.Unwrap()
	assert.Nil(t, err)
	assert.Equal(t, 12, *v)
}

func TestMapErr(t *testing.T) {
	e := errors.New("error message")
	r := Err[int](e)
	r2 := Map(r, func(i int) int {
		return i + 2
	})

	assert.True(t, r.Err())
	assert.True(t, r2.Err())
	v, err := r2.Unwrap()
	assert.Nil(t, v)
	assert.NotNil(t, err)
}
