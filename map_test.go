package collectionsutil

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_Map(t *testing.T) {
	t.Run("int squaring", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5}
		fn := func(v int) int {
			return v * v
		}
		want := []int{1, 4, 9, 16, 25}

		result := Map(in, fn)

		if len(result) != len(want) {
			t.Errorf("result length is wrong. want: %d, got: %d", len(want), len(result))
		}

		for i, v := range result {
			if v != want[i] {
				t.Errorf("mapped value(index: %d) is wrong. want: %v, got: %v", i, want[i], v)
			}
		}
	})

	t.Run("int squaring", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 22, 333, 4444, 55555}
		fn := func(v int) string {
			return strconv.Itoa(v)
		}
		want := []string{"1", "22", "333", "4444", "55555"}

		result := Map(in, fn)

		if len(result) != len(want) {
			t.Errorf("result length is wrong. want: %d, got: %d", len(want), len(result))
		}

		for i, v := range result {
			if v != want[i] {
				t.Errorf("mapped value(index: %d) is wrong. want: %v, got: %v", i, want[i], v)
			}
		}
	})
}

type Either[T any] struct {
	right T
	left  error
}

// func (e Either[T]) String() string {
// 	return fmt.Sprintf("{right: %v, left: %s}", e.right, e.left)
// }

func Test_FlatMap(t *testing.T) {
	t.Run("extract comma values list", func(t *testing.T) {
		t.Parallel()

		in := []string{
			"1,2,3,4,5",
			"6,7,eight,9,10",
		}
		fn := func(v string) []Either[int] {
			mapped := make([]Either[int], 0)
			for _, v := range strings.Split(v, ",") {
				i, err := strconv.Atoi(v)
				if err != nil {
					mapped = append(mapped, Either[int]{left: err})
					continue
				}
				mapped = append(mapped, Either[int]{right: i})
			}
			return mapped
		}
		want := []Either[int]{
			{right: 1}, {right: 2}, {right: 3}, {right: 4}, {right: 5},
			{right: 6}, {right: 7}, {left: fmt.Errorf("strconv.Atoi: parsing \"eight\": invalid syntax")}, {right: 9}, {right: 10},
		}

		result := FlatMap(in, fn)

		if len(result) != len(want) {
			t.Errorf("result length is expected %d, but got %d", len(want), len(result))
		}

		for i, v := range result {
			if want[i].right != 0 && v.right != want[i].right {
				t.Errorf("result right(index: %d) is expected %v, but got %v", i, want[i].right, v.right)
			}

			if want[i].left != nil && errors.Is(v.left, want[i].left) {
				t.Errorf("result left(index: %d) is expected %v, but got %v", i, want[i].left, v.left)
			}
		}
	})
}
