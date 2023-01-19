package maths

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRange(t *testing.T) {
	for _, test := range []struct {
		name       string
		a, b, want []int
	}{
		{
			name: "Returns single range",
			a:    []int{1, 4, 7, 12},
			want: []int{1, 4, 7, 12},
		},
		{
			name: "Merges non-overlapping ranges",
			a:    []int{1, 4, 7, 12},
			b:    []int{5, 6},
			want: []int{1, 12},
		},
		{
			name: "Merges overlapping ranges",
			a:    []int{1, 4, 7, 12},
			b:    []int{3, 8},
			want: []int{1, 12},
		},
		{
			name: "Merges partially overlapping ranges",
			a:    []int{1, 4, 7, 12},
			b:    []int{6, 8},
			want: []int{1, 4, 6, 12},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := NewRange(test.a...).Merge(NewRange(test.b...))
			if diff := cmp.Diff(test.want, got.inflectionPoints); diff != "" {
				t.Errorf("NewRange(%v).Merge(NewRange(%v)) returned incorrect values (-want, +got):\n%s", test.a, test.b, diff)
			}

			// Also test reverse
			got = NewRange(test.b...).Merge(NewRange(test.a...))
			if diff := cmp.Diff(test.want, got.inflectionPoints); diff != "" {
				t.Errorf("(Reversed order) NewRange(%v).Merge(NewRange(%v)) returned incorrect values (-want, +got):\n%s", test.b, test.a, diff)
			}
		})
	}
}
