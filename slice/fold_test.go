package slice

import (
	"testing"
)

func Test_FoldLeft(t *testing.T) {
	t.Run("sum int numbers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := 15

		got := FoldLeft(in, 0, func(acc int, in int) int {
			return acc + in
		})

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("concatenate strings reverse order", func(t *testing.T) {
		in := []string{"I", "have", "a", "pen"}
		want := "penahaveI"

		got := FoldLeft(in, "", func(acc string, in string) string {
			return in + acc
		})

		if got != want {
			t.Errorf("expect %s, but got %s", want, got)
		}
	})
}

func Test_FoldRight(t *testing.T) {
	t.Run("sum int numbers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := 15

		got := FoldRight(in, 0, func(acc int, in int) int {
			return acc + in
		})

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("concatenate strings reverse order", func(t *testing.T) {
		in := []string{"I", "have", "a", "pen"}
		want := "penahaveI"

		got := FoldRight(in, "", func(in string, acc string) string {
			return acc + in
		})

		if got != want {
			t.Errorf("expect %s, but got %s", want, got)
		}
	})
}

func Test_ReduceLeft(t *testing.T) {
	t.Run("subtract integers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := -13

		got, err := ReduceLeft(in, func(acc int, in int) int {
			return acc - in
		})
		if err != nil {
			t.Errorf("expect no error, but got %v", err)
		}

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("subtract integers(1 value)", func(t *testing.T) {
		in := []int{1}
		want := 1

		got, err := ReduceLeft(in, func(acc int, in int) int {
			return acc - in
		})
		if err != nil {
			t.Errorf("expect no error, but got %v", err)
		}

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("if input empty slice, got error.", func(t *testing.T) {
		in := []int{}
		wantErrMsg := "s length must be more than 1. but got 0"

		got, err := ReduceLeft(in, func(acc int, in int) int {
			return acc - in
		})
		if got != 0 {
			t.Errorf("expect zero value(0), but got %d", got)
		}
		if err.Error() != wantErrMsg {
			t.Errorf("expect error message %s, but got %s", wantErrMsg, err.Error())
		}
	})
}

func Test_ReduceRight(t *testing.T) {
	t.Run("subtract integers", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		want := -5

		got, err := ReduceRight(in, func(in int, acc int) int {
			return acc - in
		})
		if err != nil {
			t.Errorf("expect no error, but got %v", err)
		}

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("subtract integers(1 value)", func(t *testing.T) {
		in := []int{1}
		want := 1

		got, err := ReduceRight(in, func(in int, acc int) int {
			return acc - in
		})
		if err != nil {
			t.Errorf("expect no error, but got %v", err)
		}

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("if input empty slice, got error.", func(t *testing.T) {
		in := []int{}
		wantErrMsg := "s length must be more than 1. but got 0"

		got, err := ReduceRight(in, func(in int, acc int) int {
			return acc - in
		})
		if got != 0 {
			t.Errorf("expect zero value(0), but got %d", got)
		}
		if err.Error() != wantErrMsg {
			t.Errorf("expect error message %s, but got %s", wantErrMsg, err.Error())
		}
	})
}
