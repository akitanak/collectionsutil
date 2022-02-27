package collectionsutil

import "testing"

func Test_Filter(t *testing.T) {
	t.Run("filter odd number", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6, 7, 8}
		fn := func(v int) bool {
			return v%2 == 1
		}
		want := []int{1, 3, 5, 7}

		result := Filter(in, fn)

		if len(result) != len(want) {
			t.Errorf("result length is expected %d, but got %v", len(want), len(result))
		}

		for i, v := range result {
			if v != want[i] {
				t.Errorf("filtered value(index: %d) is expected %d, but got %d", i, want[i], v)
			}
		}
	})
}
