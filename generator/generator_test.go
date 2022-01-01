package generator

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/maths"
)

func TestContains(t *testing.T) {
	for _, test := range []struct {
		name  string
		g     *Generator
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
		g    *Generator
		ig   *IntGenerator
		want []int
	}{
		{
			name: "Generates primes",
			g:    Primes(),
			ig:   PrimesInt(),
			want: []int{
				2, 3, 5, 7, 11, 13,
			},
		},
		{
			name: "Generates triangulars",
			g:    Triangulars(),
			ig:   TriangularsInt(),
			want: []int{
				1, 3, 6, 10, 15, 21,
			},
		},
		{
			name: "Generates fibonaccis",
			g:    Fibonaccis(),
			ig:   FibonaccisInt(),
			want: []int{
				1, 1, 2, 3, 5, 8, 13, 21,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.g.Len() != 0 {
				t.Errorf("Generator.Len() returned %d; want 0", test.g.Len())
			}

			var nexts, lasts, nths []int
			var nextsI, lastsI, nthsI []*maths.Int
			for i := range test.want {
				nexts = append(nexts, test.g.Next())
				lasts = append(lasts, test.g.Last())
				nths = append(nths, test.g.Nth(i))

				nextsI = append(nextsI, test.ig.Next())
				lastsI = append(lastsI, test.ig.Last())
				nthsI = append(nthsI, test.ig.Nth(i))
			}

			var wantI []*maths.Int
			for _, w := range test.want {
				wantI = append(wantI, maths.NewInt(int64(w)))
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

			// IntGenerator
			if diff := cmp.Diff(wantI, nextsI, maths.CmpOpts()...); diff != "" {
				t.Errorf("IntGenerator.Next() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(wantI, lastsI, maths.CmpOpts()...); diff != "" {
				t.Errorf("IntGenerator.Last() returned incorrect values (-want, +got):\n%s", diff)
			}

			if diff := cmp.Diff(wantI, nthsI, maths.CmpOpts()...); diff != "" {
				t.Errorf("IntGenerator.Nth() returned incorrect values (-want, +got):\n%s", diff)
			}
		})
	}
}
