package option

import "errors"

type Option[T any] struct {
	v       *T
	is_some bool
}

func Some[T any](v T) *Option[T] {
	return &Option[T]{v: &v, is_some: true}
}

func None[T any]() *Option[T] {
	return &Option[T]{is_some: false}
}

func (o *Option[T]) Some() bool {
	return o.is_some
}

func (o *Option[T]) None() bool {
	return !o.Some()
}

func (o *Option[T]) Unwrap() (*T, error) {
	if o.Some() {
		return o.v, nil
	} else {
		return nil, errors.New("Trying to Unwrap None")
	}
}

func (o *Option[T]) UnwrapOr(other *T) *T {
	if o.Some() {
		return o.v
	} else {
		return other
	}
}
