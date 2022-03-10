package collectionsutil

func ForEach[T any](s []T, fn func(e T)) {
	for _, e := range s {
		fn(e)
	}
}

func Reverse[T any](l []T) []T {
	reversed := make([]T, len(l))
	lastIndex := len(l) - 1
	for i, v := range l {
		reversed[lastIndex-i] = v
	}
	return reversed
}
