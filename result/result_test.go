package result

import (
	"errors"
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
