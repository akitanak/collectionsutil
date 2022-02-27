package collectionsutil

func Map[IN, OUT any](l []IN, fn func(v IN) OUT) []OUT {
	mapped := make([]OUT, len(l))
	for i, v := range l {
		mapped[i] = fn(v)
	}
	return mapped
}

func FlatMap[IN, OUT any](l []IN, fn func(v IN) []OUT) []OUT {
	mapped := make([]OUT, 0)
	for _, v := range l {
		vs := fn(v)
		mapped = append(mapped, vs...)
	}
	return mapped
}
