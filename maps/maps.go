package maps

func Map[T, S, K, V comparable](m map[T]S, fn func(T, S) (K, V)) map[K]V {
	new := map[K]V{}
	for k, v := range m {
		key, value := fn(k, v)
		new[key] = value
	}
	return new
}

func Filter[T, S comparable](m map[T]S, fn func(T, S) bool) map[T]S {
	new := map[T]S{}
	for k, v := range m {
		if fn(k, v) {
			new[k] = v
		}
	}
	return new
}

func Keys[T, S comparable](input map[T]S) []T {
	var keys []T
	for k := range input {
		keys = append(keys, k)
	}
	return keys
}

func Values[T, S comparable](input map[T]S) []S {
	var values []S
	for _, v := range input {
		values = append(values, v)
	}
	return values
}
