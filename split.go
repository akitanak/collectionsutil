package collectionsutil

// SplitAt splits slice at specified index.
func SplitAt[T any](s []T, n int) ([]T, []T) {
	return Take(s, n), Drop(s, n)
}
