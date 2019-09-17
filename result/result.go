package result

import "log"

type F interface{}
type T interface{}

type Result interface {
	IsOk() bool
	IsErr() bool
	Ok() T
	Err() error
	Map(func(F) T) Result
	FMap(func(F) Result) Result
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

func (r *ok) Map(f func(F) T) Result {
	return Ok(f(r.Ok()))
}

func (r *ok) FMap(f func(F) Result) Result {
	return f(r.Ok())
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

func (r *err) Map(f func(F) T) Result {
	return r
}

func (r *err) FMap(f func(F) Result) Result {
	return r
}

func From(v T, e error) Result {
	if e != nil {
		return Err(e)
	}

	return Ok(v)
}
