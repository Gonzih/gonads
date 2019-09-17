package option

import "log"

type F interface{}
type T interface{}

type Option interface {
	Some() bool
	None() bool
	Unwrap() T
	UnwrapOr(T) T
}

type some struct {
	val T
}

func Some(v T) Option {
	return &some{val: v}
}

func (o *some) Some() bool {
	return true
}

func (o *some) None() bool {
	return false
}

func (o *some) Unwrap() T {
	return o.val
}

func (o *some) UnwrapOr(_ T) T {
	return o.val
}

type none struct {
}

func None() Option {
	return &none{}
}

func (o *none) Some() bool {
	return false
}

func (o *none) None() bool {
	return true
}

func (o *none) Unwrap() T {
	log.Panic("Unwrap call on None!")
	return nil
}

func (o *none) UnwrapOr(alt T) T {
	return alt
}

func FMap(f func(F) Option, o Option) Option {
	if o.Some() {
		return f(o.Unwrap())
	}

	return o
}

func Map(f func(F) T, o Option) Option {
	if o.Some() {
		return Some(f(o.Unwrap()))
	}

	return o
}
