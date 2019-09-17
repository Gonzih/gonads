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

func (s *some) Some() bool {
	return true
}

func (s *some) None() bool {
	return false
}

func (s *some) Unwrap() T {
	return s.val
}

func (s *some) UnwrapOr(_ T) T {
	return s.val
}

type none struct {
}

func None() Option {
	return &none{}
}

func (n *none) Some() bool {
	return false
}

func (n *none) None() bool {
	return true
}

func (n *none) Unwrap() T {
	log.Panic("Unwrap call on None!")
	return nil
}

func (n *none) UnwrapOr(alt T) T {
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