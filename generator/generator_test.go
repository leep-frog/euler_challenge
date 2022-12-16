package generator

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/command/cache"
)

func TestIsTriangular(t *testing.T) {
	for _, test := range []struct {
		t    int
		want bool
	}{
		{-1, false},
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{5, false},
		{6, true},
		{7, false},
		{8, false},
		{9, false},
		{10, true},
		{11, false},

		{53, false},
		{54, false},
		{55, true},
		{56, false},
		{57, false},
	} {
		t.Run(fmt.Sprintf("IsTriangular(%d)", test.t), func(t *testing.T) {
			if got := IsTriangular(test.t); got != test.want {
				t.Errorf("IsTriangular(%d) returned %v; want %v", test.t, got, test.want)
			}
		})
	}
}

func TestIsPentagonal(t *testing.T) {
	for _, test := range []struct {
		t    int
		want bool
	}{
		{-1, false},
		{1, true},
		{2, false},
		{3, false},
		{4, false},
		{5, true},
		{6, false},
		{7, false},
		{8, false},
		{9, false},
		{10, false},
		{11, false},
		{12, true},
		{13, false},
		{14, false},

		{143, false},
		{144, false},
		{145, true},
		{146, false},
		{147, false},
	} {
		t.Run(fmt.Sprintf("IsPentagonal(%d)", test.t), func(t *testing.T) {
			if got := IsPentagonal(test.t); got != test.want {
				t.Errorf("IsPentagonal(%d) returned %v; want %v", test.t, got, test.want)
			}
		})
	}
}

func TestIsHexagonal(t *testing.T) {
	for _, test := range []struct {
		t    int
		want bool
	}{
		{-1, false},
		{1, true},
		{2, false},
		{3, false},
		{4, false},
		{5, false},
		{6, true},
		{7, false},
		{8, false},
		{9, false},
		{10, false},
		{11, false},
		{12, false},
		{13, false},
		{14, false},
		{15, true},
		{16, false},
		{17, false},

		{40752, false},
		{40753, false},
		{40754, false},
		{40755, true},
		{40756, false},
		{40757, false},
		{40758, false},
	} {
		t.Run(fmt.Sprintf("IsHexagonal(%d)", test.t), func(t *testing.T) {
			if got := IsHexagonal(test.t); got != test.want {
				t.Errorf("IsHexagonal(%d) returned %v; want %v", test.t, got, test.want)
			}
		})
	}
}

func fakeCache(t *testing.T) {
	var _ *cache.Cache
	/*oldFunc := newCache
	newCache = func() *cache.Cache {
		return cache.NewTestCache(t)
	}
	t.Cleanup(func() {
		newCache = oldFunc
	})*/
}

func TestContains(t *testing.T) {
	fakeCache(t)
	for _, test := range []struct {
		name  string
		g     *Generator[int]
		want  bool
		nexts int
		v     int
	}{
		{
			name: "works for primes",
			g:    Primes().Generator,
			want: true,
			v:    19,
		},
		{
			name: "works when not in cycle",
			g:    Primes().Generator,
			v:    21,
		},
		{
			name:  "works for primes when already past",
			g:     Primes().Generator,
			want:  true,
			v:     19,
			nexts: 20,
		},
		{
			name:  "works when not in cycle and already past",
			g:     Primes().Generator,
			v:     21,
			nexts: 20,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			iter := test.g.Iterator()
			for i := 0; i < test.nexts; i++ {
				iter.Next()
			}
			if got := test.g.Contains(test.v); got != test.want {
				t.Errorf("InCycle(%d) returned %v; want %v", test.v, got, test.want)
			}
		})
	}
}

func TestGenerators(t *testing.T) {
	//fakeCache(t)
	for _, test := range []struct {
		name string
		g    *Generator[int]
		want []int
	}{
		{
			name: "Generates primes",
			g:    Primes().Generator,
			want: []int{
				2, 3, 5, 7, 11, 13,
			},
		},
		{
			name: "Generates best primes",
			g:    PrimesUpTo(101),
			want: []int{
				2, 3, 5, 7, 11, 13,
			},
		},
		{
			name: "Generates best primes",
			g:    FinalPrimes(11),
			want: []int{
				2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
			},
		},
		{
			name: "Generates best primes",
			g:    FinalPrimes(10),
			want: []int{
				2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
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
			name: "Generates pentagonals",
			g:    Pentagonals(),
			want: []int{
				1, 5, 12, 22, 35, 51,
			},
		},
		{
			name: "Generates hexagonals",
			g:    Hexagonals(),
			want: []int{
				1, 6, 15, 28, 45, 66,
			},
		},
		{
			name: "Generates fibonaccis",
			g:    Fibonaccis(),
			want: []int{
				1, 1, 2, 3, 5, 8, 13, 21,
			},
		},
		/* Useful for commenting out tests. */
	} {
		t.Run(test.name, func(t *testing.T) {
			iter := test.g.Iterator()
			if iter.Idx != 0 || len(test.g.values) != 0 {
				t.Errorf("Generator.len() returned %d; want 0", iter.Idx)
			}

			var nexts, lasts, nths []int
			for i := range test.want {
				nexts = append(nexts, iter.Next())
				lasts = append(lasts, iter.Last())
				nths = append(nths, test.g.Nth(i))
				if iter.Idx != i+1 || len(test.g.values) != i+1 {
					t.Errorf("Generator.len() returned %d; want %d", len(test.g.values), i+1)
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

/*func TestBigGenerators(t *testing.T) {
	fakeCache(t)
	for _, test := range []struct {
		name     string
		g        *Generator[*maths.Int]
		want     []int
		wantInts []*maths.Int
	}{
		{
			name: "Generates big fibonaccis",
			g:    BigFibonaccis(),
			want: []int{
				1, 1, 2, 3, 5, 8, 13, 21,
			},
		},
		{
			name: "Generates cubes",
			g:    PowerGenerator(3),
			want: []int{
				0, 1, 8, 27, 64, 125, 216, 343, 512,
			},
		},
		{
			name: "Generates powers of 4",
			g:    PowerGenerator(4),
			want: []int{
				0, 1, 16, 81, 256,
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			if len(test.g.values) != 0 {
				t.Errorf("Generator.len() returned %d; want 0", len(test.g.values))
			}

			want := test.wantInts
			if len(want) == 0 {
				for _, w := range test.want {
					want = append(want, maths.NewInt(w))
				}
			}
			var nexts, lasts, nths []*maths.Int
			iter := test.g.Iterator()
			for i := range want {
				nexts = append(nexts, iter.Next())
				lasts = append(lasts, iter.Last())
				nths = append(nths, test.g.Nth(i))
				if len(test.g.values) != i+1 || iter.Idx != i+1 {
					t.Errorf("Generator.len() returned %d; want %d", len(test.g.values), i+1)
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
*/
