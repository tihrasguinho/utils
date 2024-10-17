package slices

func Map[T, S any](input []T, fn func(T) S) []S {
	var new []S
	for _, v := range input {
		new = append(new, fn(v))
	}
	return new
}

func Filter[T any](input []T, fn func(T) bool) []T {
	var new []T
	for _, v := range input {
		if fn(v) {
			new = append(new, v)
		}
	}
	return new
}

func Fold[T, S any](input []T, initialValue S, fn func(prev S, next T) S) S {
	var result S = initialValue
	for _, v := range input {
		result = fn(result, v)
	}
	return result
}

func Reverse[T any](input []T) []T {
	var new []T
	for i := len(input) - 1; i >= 0; i-- {
		new = append(new, input[i])
	}
	return new
}

func Contains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}
	return false
}

func Distinct[T comparable](input []T) []T {
	var new []T
	for _, v := range input {
		if !Contains(new, v) {
			new = append(new, v)
		}
	}
	return new
}

func Sort[T any](input *[]T, compare func(T, T) bool) {
	for i := 0; i < len(*input); i++ {
		for j := 0; j < len(*input)-i-1; j++ {
			if compare((*input)[j], (*input)[j+1]) {
				(*input)[j], (*input)[j+1] = (*input)[j+1], (*input)[j]
			}
		}
	}
}
