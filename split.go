package collectionsutil

func SplitAt[T any](s []T, n int) ([]T, []T) {
	return Take(s, n), Drop(s, n)
}
