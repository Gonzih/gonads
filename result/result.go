package result

type Result[T any] struct {
	v   *T
	err error
}

func From[T any](v T, err error) *Result[T] {
	if err != nil {
		return Err[T](err)
	} else {
		return Ok(v)
	}
}

func Ok[T any](v T) *Result[T] {
	return &Result[T]{&v, nil}
}

func Err[T any](err error) *Result[T] {
	return &Result[T]{nil, err}
}

func (r *Result[T]) Ok() bool {
	return r.err == nil
}

func (r *Result[T]) Err() bool {
	return !r.Ok()
}

func (r *Result[T]) Unwrap() (*T, error) {
	return r.v, r.err
}

func (r *Result[T]) UnwrapOr(other *T) *T {
	if r.Ok() {
		return r.v
	} else {
		return other
	}
}

func Map[F any, T any](o *Result[F], f func(v F) T) *Result[T] {
	if o.Ok() {
		return Ok(f(*o.v))
	} else {
		return Err[T](o.err)
	}
}
