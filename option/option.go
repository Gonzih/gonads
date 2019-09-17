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

type some struct {
	val T
}

func Some(v T) Option {
	return &some{val: v}
}

func (o *some) IsSome() bool {
	return true
}

func (o *some) IsNone() bool {
	return !o.IsSome()
}

func (o *some) Unwrap() T {
	return o.val
}

func (o *some) UnwrapOr(_ T) T {
	return o.val
}

func (o *some) Map(f func(F) T) Option {
	return Some(f(o.Unwrap()))
}

func (o *some) FMap(f func(F) Option) Option {
	return f(o.Unwrap())
}

type none struct {
}

func None() Option {
	return &none{}
}

func (o *none) IsSome() bool {
	return false
}

func (o *none) IsNone() bool {
	return !o.IsSome()
}

func (o *none) Unwrap() T {
	log.Panic("Unwrap call on None!")
	return nil
}

func (o *none) UnwrapOr(alt T) T {
	return alt
}

func (o *none) Map(f func(F) T) Option {
	return o
}

func (o *none) FMap(f func(F) Option) Option {
	return o
}
