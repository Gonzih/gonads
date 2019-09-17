package result

import "log"

type F interface{}
type T interface{}

type Result interface {
	IsOk() bool
	IsErr() bool
	Ok() T
	Err() error
}

type ok struct {
	val T
}

func Ok(v T) Result {
	return &ok{val: v}
}

func (r *ok) IsOk() bool {
	return true
}

func (r *ok) IsErr() bool {
	return false
}

func (r *ok) Ok() T {
	return r.val
}

func (r *ok) Err() error {
	log.Panic("Not an Err type!")
	return nil
}

type err struct {
	err error
}

func Err(e error) Result {
	return &err{err: e}
}

func (r *err) IsOk() bool {
	return false
}

func (r *err) IsErr() bool {
	return true
}

func (r *err) Ok() T {
	log.Panic("Not an Ok type!")
	return nil
}

func (r *err) Err() error {
	return r.err
}

func FMap(f func(F) Result, r Result) Result {
	if r.IsOk() {
		return f(r.Ok())
	}

	return r
}

func Map(f func(F) T, r Result) Result {
	if r.IsOk() {
		return Ok(f(r.Ok()))
	}

	return r
}
