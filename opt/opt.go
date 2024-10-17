package opt

type Opt[T any] interface {
	HasValue() bool
	Get() T
}

type some[T any] struct {
	Value T
}

func Some[T any](value T) some[T] {
	return some[T]{Value: value}
}

func (s some[T]) HasValue() bool {
	return true
}

func (s some[T]) Get() T {
	return s.Value
}

type none[T any] struct{}

func None[T any]() none[T] {
	return none[T]{}
}

func (n none[T]) HasValue() bool {
	return false
}

func (n none[T]) Get() T {
	return *new(T)
}
