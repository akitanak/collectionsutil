package collectionsutil

type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Numeric](s []T) T {
	return FoldLeft(s, 0, func(acc T, in T) T {
		return acc + in
	})
}

func Product[T Numeric](s []T) T {
	return FoldLeft(s, 1, func(acc T, in T) T {
		return acc * in
	})
}

func Max[T Numeric](s []T) T {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return s[0]
	}

	return FoldLeft(s[1:], s[0], func(max T, in T) T {
		if max < in {
			return in
		}
		return max
	})
}

func Min[T Numeric](s []T) T {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return s[0]
	}

	return FoldLeft(s[1:], s[0], func(min T, in T) T {
		if min > in {
			return in
		}
		return min
	})
}
