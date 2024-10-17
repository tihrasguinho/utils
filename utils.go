package utils

import (
	"github.com/tihrasguinho/utils/opt"
)

func Some[T any](value T) opt.Opt[T] {
	return opt.Some(value)
}

func None[T any]() opt.Opt[T] {
	return opt.None[T]()
}
