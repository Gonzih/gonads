package result

type [T any]Result struct {
	v T
	err error
}

func Ok[T any](v T) Result {
	return &Result{v}
}

func Err[T any](err error) [T]Result {
	return &Result{err}
}

func (r *Result) Ok() bool {
	return r.err == nil
}

func (r *Result) Err() bool {
	return !r.Ok()
}

func (r *Result) Unwrap[T any]() (T, error) {
	return (r.v, r.err)
}

func (r *OkImpl) UnwrapOr[T any](other T) T {
	if r.Ok() {
		return r.v
	} else {
		return other
	}
}
