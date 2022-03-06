package collectionsutil

func Reverse[T any](l []T) []T {
	reversed := make([]T, len(l))
	lastIndex := len(l) - 1
	for i, v := range l {
		reversed[lastIndex-i] = v
	}
	return reversed
}
