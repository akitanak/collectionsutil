package maputil

import (
	"fmt"
	"sort"
	"testing"
)

func Test_GetOrElse(t *testing.T) {
	testcases := map[string]struct {
		inMap     map[string]int
		inKey     string
		inDefault int
		want      int
	}{
		"if key does not exist, get default value": {
			inMap: map[string]int{
				"a": 10,
			},
			inKey:     "b",
			inDefault: -10,
			want:      -10,
		},
		"if key exist, get correspond value": {
			inMap: map[string]int{
				"a": 10,
			},
			inKey:     "a",
			inDefault: -10,
			want:      10,
		},
	}

	for name, tc := range testcases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := GetOrElse(tc.inMap, tc.inKey, tc.inDefault)

			if got != tc.want {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}
		})
	}
}

func Test_Contains(t *testing.T) {
	testcases := map[string]struct {
		inMap map[string]int
		inKey string
		want  bool
	}{
		"if key exist, return true": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			inKey: "b",
			want:  true,
		},
		"if key does not exist, return false": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			inKey: "d",
			want:  false,
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Contains(tc.inMap, tc.inKey)

			if got != tc.want {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}
		})
	}
}

func Test_Keys(t *testing.T) {
	testcases := map[string]struct {
		inMap map[string]int
		want  []string
	}{
		"if empty map, returns empty slice": {
			inMap: map[string]int{},
			want:  []string{},
		},
		"if map contains a key, returns a slice which has single key": {
			inMap: map[string]int{
				"a": 1,
			},
			want: []string{"a"},
		},
		"if map contains some key, returns a slice which has same keys ": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: []string{"a", "b", "c"},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Keys(tc.inMap)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			sort.Slice(tc.want, func(i, j int) bool { return tc.want[i] < tc.want[j] })
			for i, g := range got {
				if g != tc.want[i] {
					t.Errorf("expect %v, but got %v", tc.want[i], g)
				}
			}
		})
	}
}

func Test_Vaues(t *testing.T) {
	testcases := map[string]struct {
		inMap map[string]int
		want  []int
	}{
		"if empty map, returns empty slice": {
			inMap: map[string]int{},
			want:  []int{},
		},
		"if map contains a key, returns a slice which has single value": {
			inMap: map[string]int{
				"a": 1,
			},
			want: []int{1},
		},
		"if map contains some key, returns a slice which has same values": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: []int{1, 2, 3},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Values(tc.inMap)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			sort.Slice(tc.want, func(i, j int) bool { return tc.want[i] < tc.want[j] })
			for i, g := range got {
				if g != tc.want[i] {
					t.Errorf("expect %v, but got %v", tc.want[i], g)
				}
			}
		})
	}
}

func Test_Entries(t *testing.T) {
	testcases := map[string]struct {
		inMap map[string]int
		want  []Entry[string, int]
	}{
		"if empty map, returns empty slice": {
			inMap: map[string]int{},
			want:  []Entry[string, int]{},
		},
		"if map contains a key, returns a slice which has single entry": {
			inMap: map[string]int{
				"a": 1,
			},
			want: []Entry[string, int]{
				{"a", 1},
			},
		},
		"if map contains some key, returns a slice which has same entry": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			want: []Entry[string, int]{
				{"a", 1},
				{"b", 2},
				{"c", 3},
			},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Entries(tc.inMap)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			sort.Slice(got, func(i, j int) bool { return got[i].Key < got[j].Key })
			sort.Slice(tc.want, func(i, j int) bool { return tc.want[i].Key < tc.want[j].Key })
			for i, g := range got {
				if g != tc.want[i] {
					t.Errorf("expect %v, but got %v", tc.want[i], g)
				}
			}
		})
	}
}

func Test_FilterKeys(t *testing.T) {
	testcases := map[string]struct {
		inMap      map[string]int
		inFilterFn func(string) bool
		want       map[string]int
	}{
		"if empty map, returns empty map": {
			inMap:      map[string]int{},
			inFilterFn: func(k string) bool { return true },
			want:       map[string]int{},
		},
		"if map contains matched key, returns a map which has correspond key": {
			inMap: map[string]int{
				"a": 1,
			},
			inFilterFn: func(k string) bool { return k == "a" },
			want: map[string]int{
				"a": 1,
			},
		},
		"if map contains some matched keys, returns a map which has correspond keys": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			inFilterFn: func(k string) bool { return k == "a" || k == "c" },
			want: map[string]int{
				"a": 1,
				"c": 3,
			},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := FilterKeys(tc.inMap, tc.inFilterFn)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			for k, v := range got {
				if v != tc.want[k] {
					t.Errorf("expect %v, but got %v", tc.want[k], v)
				}
			}
		})
	}
}

func Test_Transform(t *testing.T) {
	testcases := map[string]struct {
		inMap         map[string]int
		inTransformFn func(string, int) int
		want          map[string]int
	}{
		"if empty map, returns empty map": {
			inMap:         map[string]int{},
			inTransformFn: func(k string, v int) int { return v * 2 },
			want:          map[string]int{},
		},
		"if map contains matched key, returns a map which has correspond key": {
			inMap: map[string]int{
				"a": 1,
			},
			inTransformFn: func(k string, v int) int { return v * 2 },
			want: map[string]int{
				"a": 2,
			},
		},
		"if map contains some matched keys, returns a map which has correspond keys": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			inTransformFn: func(k string, v int) int { return v * 2 },
			want: map[string]int{
				"a": 2,
				"b": 4,
				"c": 6,
			},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := Transform(tc.inMap, tc.inTransformFn)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			for k, v := range got {
				if v != tc.want[k] {
					t.Errorf("expect %v, but got %v", tc.want[k], v)
				}
			}
		})
	}
}

func Test_MapValues(t *testing.T) {
	testcases := map[string]struct {
		inMap   map[string]int
		inMapFn func(string, int) string
		want    map[string]string
	}{
		"if empty map, returns empty map": {
			inMap:   map[string]int{},
			inMapFn: func(k string, v int) string { return fmt.Sprintf("%v", v) },
			want:    map[string]string{},
		},
		"if map contains matched key, returns a map which has correspond key": {
			inMap: map[string]int{
				"a": 1,
			},
			inMapFn: func(k string, v int) string { return fmt.Sprintf("%v", v) },
			want: map[string]string{
				"a": "1",
			},
		},
		"if map contains some matched keys, returns a map which has correspond keys": {
			inMap: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			inMapFn: func(k string, v int) string { return fmt.Sprintf("%v", v) },
			want: map[string]string{
				"a": "1",
				"b": "2",
				"c": "3",
			},
		},
	}

	for name, tc := range testcases {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := MapValues(tc.inMap, tc.inMapFn)

			if len(got) != len(tc.want) {
				t.Errorf("expect %v, but got %v", tc.want, got)
			}

			for k, v := range got {
				if v != tc.want[k] {
					t.Errorf("expect %v, but got %v", tc.want[k], v)
				}
			}
		})
	}
}
