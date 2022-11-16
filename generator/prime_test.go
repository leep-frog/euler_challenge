package generator

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFactoring(t *testing.T) {
	for _, test := range []struct {
		n    int
		want []int
	}{
		{
			0,
			nil,
		},
		{
			1,
			[]int{1},
		},
		{
			2,
			[]int{1, 2},
		},
		{
			3,
			[]int{1, 3},
		},
		{
			4,
			[]int{1, 2, 4},
		},
		{
			5,
			[]int{1, 5},
		},
		{
			299,
			[]int{1, 13, 23, 299},
		},
		{
			307,
			[]int{1, 307},
		},
		{
			// Ensure functions are efficient for large primes.
			15485867, // = Primes().Nth(1_000_000),
			[]int{1, 15485867},
		},
		{
			// Ensure functions are efficient for large primes.
			15485866, // = Primes().Nth(1_000_000) - 1,
			[]int{1, 2, 11, 22, 703903, 1407806, 7742933, 15485866},
		},
	} {
		t.Run(fmt.Sprintf("Factors fo %d", test.n), func(t *testing.T) {
			clearCaches()

			primes := Primes()

			if got := FactorCount(test.n, primes); got != len(test.want) {
				t.Errorf("FactorCount(%d) returned %d; want %d", test.n, got, len(test.want))
			}

			clearCaches()

			if diff := cmp.Diff(test.want, Factors(test.n, primes)); diff != "" {
				t.Errorf("Factors(%d) returned incorrect result (-want, +got):\n%s", test.n, diff)
			}

			clearCaches()

		})
	}
}
