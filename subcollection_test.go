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

		got := Filter(in, fn)

		if len(got) != len(want) {
			t.Errorf("result length is expected %d, but got %v", len(want), len(got))
		}

		for i, e := range got {
			if e != want[i] {
				t.Errorf("filtered value(index: %d) is expected %d, but got %d", i, want[i], e)
			}
		}
	})
}

func Test_Take(t *testing.T) {
	t.Run("take 5", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num := 5
		want := []int{1, 2, 3, 4, 5}

		got := Take(in, num)

		if len(got) != num {
			t.Errorf("expect slice have %d elements, but got %d", num, len(got))
		}
		for i, e := range got {
			if e != want[i] {
				t.Errorf("expect %d at index %d, but got %d", want[i], i, e)
			}
		}
	})
}

func Test_Drop(t *testing.T) {
	t.Run("drop first 5 elements", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num := 5
		want := []int{6, 7, 8, 9, 10}

		got := Drop(in, num)

		if len(got) != num {
			t.Errorf("expect slice have %d elements, but got %d", num, len(got))
		}
		for i, e := range got {
			if e != want[i] {
				t.Errorf("expect %d at index %d, but got %d", want[i], i, e)
			}
		}
	})
}
