package slice

import "testing"

func Test_SplitAt(t *testing.T) {
	t.Run("empty slice splitted two empty slice", func(t *testing.T) {
		t.Parallel()

		in := []int{}
		n := 0
		want1 := []int{}
		want2 := []int{}

		got1, got2 := SplitAt(in, n)

		if len(got1) != len(want1) {
			t.Errorf("result length is expected %d, but got %v", len(want1), len(got1))
		}
		if len(got2) != len(want2) {
			t.Errorf("result length is expected %d, but got %v", len(want2), len(got2))
		}

		for i, e := range got1 {
			if e != want1[i] {
				t.Errorf("splitted value(index: %d) is expected %d, but got %d", i, want1[i], e)
			}
		}
		for i, e := range got2 {
			if e != want2[i] {
				t.Errorf("splitted value(index: %d) is expected %d, but got %d", i, want2[i], e)
			}
		}
	})

	t.Run("split half position", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6}
		n := 3
		want1 := []int{1, 2, 3}
		want2 := []int{4, 5, 6}

		got1, got2 := SplitAt(in, n)

		if len(got1) != len(want1) {
			t.Errorf("result length is expected %d, but got %v", len(want1), len(got1))
		}
		if len(got2) != len(want2) {
			t.Errorf("result length is expected %d, but got %v", len(want2), len(got2))
		}

		for i, e := range got1 {
			if e != want1[i] {
				t.Errorf("splitted value(index: %d) is expected %d, but got %d", i, want1[i], e)
			}
		}
		for i, e := range got2 {
			if e != want2[i] {
				t.Errorf("splitted value(index: %d) is expected %d, but got %d", i, want2[i], e)
			}
		}
	})
}
