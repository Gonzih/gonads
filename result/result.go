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

type OkImpl struct {
	val T
}

func Ok(v T) Result {
	return &OkImpl{val: v}
}

func (r *OkImpl) IsOk() bool {
	return true
}

func (r *OkImpl) IsErr() bool {
	return false
}

func (r *OkImpl) Ok() T {
	return r.val
}

func (r *OkImpl) Err() error {
	log.Panic("Not an Err type!")
	return nil
}

func (r *OkImpl) Map(f func(F) T) Result {
	return Ok(f(r.Ok()))
}

func (r *OkImpl) FMap(f func(F) Result) Result {
	return f(r.Ok())
}

type ErrImpl struct {
	err error
}

func Err(e error) Result {
	return &err{err: e}
}

func (r *ErrImpl) IsOk() bool {
	return false
}

func (r *ErrImpl) IsErr() bool {
	return true
}

func (r *ErrImpl) Ok() T {
	log.Panic("Not an Ok type!")
	return nil
}

func (r *ErrImpl) Err() error {
	return r.err
}

func (r *ErrImpl) Map(f func(F) T) Result {
	return r
}

func (r *ErrImpl) FMap(f func(F) Result) Result {
	return r
}

func From(v T, e error) Result {
	if e != nil {
		return Err(e)
	}

	return Ok(v)
}
