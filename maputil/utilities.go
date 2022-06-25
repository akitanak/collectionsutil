package maputil

func GetOrElse[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	v, ok := m[key]
	if !ok {
		return defaultValue
	}
	return v
}

func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], len(m))
	i := 0
	for k, v := range m {
		entries[i] = Entry[K, V]{k, v}
		i++
	}

	return entries
}

func FilterKeys[K comparable, V any](m map[K]V, filterFn func(key K) bool) map[K]V {
	filtered := make(map[K]V)
	for k, v := range m {
		if filterFn(k) {
			filtered[k] = v
		}
	}
	return filtered
}

func Transform[K comparable, V any](m map[K]V, transformFn func(k K, v V) V) map[K]V {
	transformed := make(map[K]V)
	for k, v := range m {
		transformed[k] = transformFn(k, v)
	}
	return transformed
}
