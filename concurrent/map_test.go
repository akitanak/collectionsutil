package concurrent

import (
	"strconv"
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
