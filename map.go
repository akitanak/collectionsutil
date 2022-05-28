package collectionsutil

// Map maps input slice values to output slice values `l` with specified function `fn`.
func Map[IN, OUT any](s []IN, fn func(v IN) OUT) []OUT {
	mapped := make([]OUT, len(s))
	for i, v := range s {
		mapped[i] = fn(v)
	}
	return mapped
}

// FlatMap maps input slice values `l` with specified function `fn` that returns slice and flatten slices.
func FlatMap[IN, OUT any](s []IN, fn func(v IN) []OUT) []OUT {
	mapped := make([]OUT, 0)
	for _, v := range s {
		vs := fn(v)
		mapped = append(mapped, vs...)
	}
	return mapped
}

// Flatten flattens nested slice `s`.
func Flatten[T any](s [][]T) []T {
	flattened := make([]T, 0)
	for _, inner := range s {
		flattened = append(flattened, inner...)
	}
	return flattened
}
