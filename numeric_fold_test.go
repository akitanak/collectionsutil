package collectionsutil

import "testing"

func Test_Sum(t *testing.T) {
	t.Run("sum int", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 55

		got := Sum(in)

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("sum float64", func(t *testing.T) {
		t.Parallel()

		in := []float64{.1, .2, .3, .4, .5, .6, .7, .8, .9, 1.0}
		want := 5.5

		got := Sum(in)
		t.Log("")
		if got != want {
			t.Errorf("expect %v, but got %v", want, got)
		}
	})

	testcases := map[string]struct {
		in   []int
		want int
	}{
		"sum 1 element": {
			in:   []int{5},
			want: 5,
		},
		"sum empty slice": {
			in:   []int{},
			want: 0,
		},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Sum(tc.in)

			if got != tc.want {
				t.Errorf("expect %d, but got %d", tc.want, got)
			}
		})

	}
}

func Test_Product(t *testing.T) {
	t.Run("product int", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		want := 3628800

		got := Product(in)

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("product float64", func(t *testing.T) {
		t.Parallel()

		in := []float64{1., 2., 3., 4., 5., 6., 7., 8., 9., 10.}
		want := 3628800.

		got := Product(in)

		if got != want {
			t.Errorf("expect %v, but got %v", want, got)
		}
	})

	testcases := map[string]struct {
		in   []int
		want int
	}{
		"product 1 element": {
			in:   []int{5},
			want: 5,
		},
		"product empty slice": {
			in:   []int{},
			want: 1,
		},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Product(tc.in)

			if got != tc.want {
				t.Errorf("expect %d, but got %d", tc.want, got)
			}
		})
	}
}

func Test_Max(t *testing.T) {
	t.Run("max int", func(t *testing.T) {
		t.Parallel()
		in := []int{1, -5, -9, 10, 3, 9, 5, 8}
		want := 10

		got := Max(in)

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("max float", func(t *testing.T) {
		t.Parallel()
		in := []float64{1.0, 10.1, -5.0, -9.0, 10.0, 3.5, 9.9, 5.0, 8.99}
		want := 10.1

		got := Max(in)

		if got != want {
			t.Errorf("expect %f, but got %f", want, got)
		}
	})

	testcases := map[string]struct {
		in   []int
		want int
	}{
		"max 1 element": {
			in:   []int{5},
			want: 5,
		},
		"max empty slice": {
			in:   []int{},
			want: 0,
		},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Max(tc.in)

			if got != tc.want {
				t.Errorf("expect %d, but got %d", tc.want, got)
			}
		})
	}
}

func Test_Min(t *testing.T) {
	t.Run("min int", func(t *testing.T) {
		t.Parallel()
		in := []int{1, -5, -9, 10, 3, 9, 5, 8}
		want := -9

		got := Min(in)

		if got != want {
			t.Errorf("expect %d, but got %d", want, got)
		}
	})

	t.Run("min float", func(t *testing.T) {
		t.Parallel()
		in := []float64{1.0, -10.1, -5.0, 9.0, 10.0, 3.5, -9.9, 5.0, 8.99}
		want := -10.1

		got := Min(in)

		if got != want {
			t.Errorf("expect %f, but got %f", want, got)
		}
	})

	testcases := map[string]struct {
		in   []int
		want int
	}{
		"min 1 element": {
			in:   []int{5},
			want: 5,
		},
		"min empty slice": {
			in:   []int{},
			want: 0,
		},
	}
	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Min(tc.in)

			if got != tc.want {
				t.Errorf("expect %d, but got %d", tc.want, got)
			}
		})
	}
}
