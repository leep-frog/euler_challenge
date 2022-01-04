package generator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/maths"
)

func TestContains(t *testing.T) {
	for _, test := range []struct {
		name  string
		g     *Generator[int]
		want  bool
		nexts int
		v     int
	}{
		{
			name: "works for primes",
			g:    Primes(),
			want: true,
			v:    19,
		},
		{
			name: "works when not in cycle",
			g:    Primes(),
			v:    21,
		},
		{
			name:  "works for primes when already past",
			g:     Primes(),
			want:  true,
			v:     19,
			nexts: 20,
		},
		{
			name:  "works when not in cycle and already past",
			g:     Primes(),
			v:     21,
			nexts: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < test.nexts; i++ {
				test.g.Next()
			}
			if got := test.g.Contains(test.v); got != test.want {
				t.Errorf("InCycle(%d) returned %v; want %v", test.v, got, test.want)
			}
		})
	}
}

func TestGenerators(t *testing.T) {
	for _, test := range []struct {
		name string
		g    *Generator[int]
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
				1, 1, 2, 3, 5, 8, 13, 21,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.g.len() != 0 {
				t.Errorf("Generator.len() returned %d; want 0", test.g.len())
			}

			var nexts, lasts, nths []int
			for i := range test.want {
				nexts = append(nexts, test.g.Next())
				lasts = append(lasts, test.g.last())
				nths = append(nths, test.g.Nth(i))
				if test.g.len() != i+1 {
					t.Errorf("Generator.len() returned %d; want %d", test.g.len(), i+1)
				}
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

func TestBigGenerators(t *testing.T) {
	for _, test := range []struct {
		name string
		g    *Generator[*maths.Int]
		want []int
	}{
		{
			name: "Generates big fibonaccis",
			g:    BigFibonaccis(),
			want: []int{
				1, 1, 2, 3, 5, 8, 13, 21,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.g.len() != 0 {
				t.Errorf("Generator.len() returned %d; want 0", test.g.len())
			}

			var want, nexts, lasts, nths []*maths.Int
			for i, w := range test.want {
				want = append(want, maths.NewInt(int64(w)))
				nexts = append(nexts, test.g.Next())
				lasts = append(lasts, test.g.last())
				nths = append(nths, test.g.Nth(i))
				if test.g.len() != i+1 {
					t.Errorf("Generator.len() returned %d; want %d", test.g.len(), i+1)
				}
			}

			if diff := cmp.Diff(want, nexts, maths.CmpOpts()...); diff != "" {
				t.Errorf("Generator.Next() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(want, lasts, maths.CmpOpts()...); diff != "" {
				t.Errorf("Generator.Last() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(want, nths, maths.CmpOpts()...); diff != "" {
				t.Errorf("Generator.Nth() returned incorrect values (-want, +got):\n%s", diff)
			}
		})
	}
}
