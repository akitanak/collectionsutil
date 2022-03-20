package concurrent

import (
	"testing"
)

func Test_ForEach(t *testing.T) {
	t.Run("doubled", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5}
		want := []int{2, 4, 6, 8, 10}
		recorder := make([]int, len(in))
		counter := 0
		fn := func(e int) {
			recorder[counter] = e * 2
			counter++
		}

		ForEach(in, fn)

		for i, e := range recorder {
			if e != want[i] {
				t.Errorf("expected %d at index %d, but got %d", want[i], i, e)
			}
		}
	})
}
