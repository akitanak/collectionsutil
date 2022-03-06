package collectionsutil

import "fmt"

func FoldLeft[IN, OUT any](l []IN, z OUT, fn func(acc OUT, in IN) OUT) OUT {
	acc := z
	for _, v := range l {
		acc = fn(acc, v)
	}
	return acc
}

func FoldRight[IN, OUT any](l []IN, z OUT, fn func(in IN, acc OUT) OUT) OUT {
	acc := z
	for i := len(l); i > 0; i-- {
		acc = fn(l[i-1], acc)
	}
	return acc
}

func ReduceLeft[T any](l []T, fn func(acc T, in T) T) (T, error) {
	var acc T
	if len(l) < 1 {
		return acc, fmt.Errorf("l length must be more than 1. but got %d", len(l))
	}
	acc = l[0]
	for i := 1; i < len(l); i++ {
		acc = fn(acc, l[i])
	}
	return acc, nil
}
