package collectionsutil

func Filter[T any](s []T, fn func(e T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range s {
		if fn(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func Take[T any](s []T, n int) []T {
	return s[:n]
}

func Drop[T any](s []T, n int) []T {
	return s[n:]
}
