package collectionsutil

func Filter[V any](l []V, fn func(v V) bool) []V {
	filtered := make([]V, 0)
	for _, v := range l {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
