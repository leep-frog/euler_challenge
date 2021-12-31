package generator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerators(t *testing.T) {
	for _, test := range []struct {
		name string
		g    *Generator
		want []int
	}{
		{
			name: "Generates primes",
			g:    Primes(),
			want: []int{
				2, 3, 5, 7, 11, 13,
			},
		},
		{
			name: "Generates triangulars",
			g:    Triangulars(),
			want: []int{
				1, 3, 6, 10, 15, 21,
			},
		},
		{
			name: "Generates fibonaccis",
			g:    Fibonaccis(),
			want: []int{
				1, 2, 3, 5, 8, 13, 21,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.g.Len() != 0 {
				t.Errorf("Generator.Len() returned %d; want 0", test.g.Len())
			}

			var nexts, lasts, nths []int
			for i := range test.want {
				nexts = append(nexts, test.g.Next())
				lasts = append(lasts, test.g.Last())
				nths = append(nths, test.g.Nth(i))
			}

			if diff := cmp.Diff(test.want, nexts); diff != "" {
				t.Errorf("Generator.Next() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(test.want, lasts); diff != "" {
				t.Errorf("Generator.Last() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(test.want, nths); diff != "" {
				t.Errorf("Generator.Nth() returned incorrect values (-want, +got):\n%s", diff)
			}
		})
	}
}
