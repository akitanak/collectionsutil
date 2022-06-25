package sliceutil

// ForEach apply specified function to specified slice values.
func ForEach[T any](s []T, fn func(e T)) {
	for _, e := range s {
		fn(e)
	}
}

// Reverse reverses specified slice values order.
func Reverse[T any](s []T) []T {
	reversed := make([]T, len(s))
	lastIndex := len(s) - 1
	for i, v := range s {
		reversed[lastIndex-i] = v
	}
	return reversed
}
