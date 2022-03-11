package collectionsutil

import "fmt"

func FoldLeft[IN, OUT any](s []IN, z OUT, fn func(acc OUT, in IN) OUT) OUT {
	acc := z
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func FoldRight[IN, OUT any](s []IN, z OUT, fn func(in IN, acc OUT) OUT) OUT {
	acc := z
	for i := len(s); i > 0; i-- {
		acc = fn(s[i-1], acc)
	}
	return acc
}

func ReduceLeft[T any](s []T, fn func(acc T, in T) T) (T, error) {
	var acc T
	if len(s) < 1 {
		return acc, fmt.Errorf("s length must be more than 1. but got %d", len(s))
	}
	acc = s[0]
	for i := 1; i < len(s); i++ {
		acc = fn(acc, s[i])
	}
	return acc, nil
}

func ReduceRight[T any](s []T, fn func(in T, acc T) T) (T, error) {
	var acc T
	if len(s) < 1 {
		return acc, fmt.Errorf("s length must be more than 1. but got %d", len(s))
	}
	acc = s[len(s)-1]
	for i := len(s) - 1; i > 0; i-- {
		acc = fn(s[i-1], acc)
	}
	return acc, nil
}
