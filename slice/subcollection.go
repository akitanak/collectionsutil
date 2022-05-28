package slice

// Filter filters specified slice values s with specified function fn.
func Filter[T any](s []T, fn func(e T) bool) []T {
	filtered := make([]T, 0)
	for _, e := range s {
		if fn(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

// Take returns a slice consists leading n elements.
func Take[T any](s []T, n int) []T {
	return s[:n]
}

// Drop returns a slice except leading n elements.
func Drop[T any](s []T, n int) []T {
	return s[n:]
}

// TakeWhile returns some leading elements that satisfies the condition that represented with function fn.
func TakeWhile[T any](s []T, fn func(e T) bool) []T {
	lastMatch := 0
	for _, e := range s {
		if !fn(e) {
			break
		}
		lastMatch++
	}
	return s[:lastMatch]
}

// DropWhile returns a slice that droped elements until satisfy the condition that represented with function fn.
func DropWhile[T any](s []T, fn func(e T) bool) []T {
	lastMatch := 0
	for _, e := range s {
		if !fn(e) {
			break
		}
		lastMatch++
	}
	return s[lastMatch:]
}
