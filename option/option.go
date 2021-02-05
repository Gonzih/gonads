package option

import "errors"

type [T any]Option struct {
	v T
	is_some bool
}

func Some[T any](v T) Option {
	return &Option{v: v, is_some: true}
}

func None[T any]() Option {
	return &Option{}
}

func (o *Option) Some() bool {
	return op.is_some
}

func (o *Option) None() bool {
	return !op.Some()
}

func (o *Option) Unwrap[T any]() T {
	if o.Some() {
		return (o.v, nil)
	} else {
		return (nil, errors.New("Trying to Unwrap None"))
	}
}

func (o *Option) UnwrapOr[T any](other T) T {
	if o.Some() {
		return o.v
	} else {
		return other
	}
}
