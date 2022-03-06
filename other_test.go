package collectionsutil

import (
	"testing"
)

func Test_Reverse(t *testing.T) {
	testCase := map[string]struct {
		in   []int
		want []int
	}{
		"even-numbered items": {
			in:   []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		"odd-numbered items": {
			in:   []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{7, 6, 5, 4, 3, 2, 1},
		},
	}

	for name, tc := range testCase {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := Reverse(tc.in)

			if len(got) != len(tc.want) {
				t.Errorf("expected length %d, but got %d", len(tc.want), len(got))
			}
			for i, v := range got {
				if v != tc.want[i] {
					t.Errorf("index %d: expected %d, but got %d", i, tc.want[i], v)
				}
			}

			r := Reverse(got)

			if len(r) != len(tc.in) {
				t.Errorf("expected length %d, but got %d", len(tc.want), len(got))
			}
			for i, v := range r {
				if v != tc.in[i] {
					t.Errorf("index %d: expected %d, but got %d", i, tc.in[i], v)
				}
			}
		})
	}
}
