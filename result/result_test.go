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
