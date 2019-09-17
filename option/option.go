package option

import "log"

type F interface{}
type T interface{}

type Option interface {
	IsSome() bool
	IsNone() bool
	Unwrap() T
	UnwrapOr(T) T
	Map(func(F) T) Option
	FMap(func(F) Option) Option
}

type SomeImpl struct {
	val T
}

func Some(v T) Option {
	return &SomeImpl{val: v}
}

func (o *SomeImpl) IsSome() bool {
	return true
}

func (o *SomeImpl) IsNone() bool {
	return !o.IsSome()
}

func (o *SomeImpl) Unwrap() T {
	return o.val
}

func (o *SomeImpl) UnwrapOr(_ T) T {
	return o.val
}

func (o *SomeImpl) Map(f func(F) T) Option {
	return Some(f(o.Unwrap()))
}

func (o *SomeImpl) FMap(f func(F) Option) Option {
	return f(o.Unwrap())
}

type NoneImpl struct {
}

func None() Option {
	return &NoneImpl{}
}

func (o *NoneImpl) IsSome() bool {
	return false
}

func (o *NoneImpl) IsNone() bool {
	return !o.IsSome()
}

func (o *NoneImpl) Unwrap() T {
	log.Panic("Unwrap call on None!")
	return nil
}

func (o *NoneImpl) UnwrapOr(alt T) T {
	return alt
}

func (o *NoneImpl) Map(f func(F) T) Option {
	return o
}

func (o *NoneImpl) FMap(f func(F) Option) Option {
	return o
}
