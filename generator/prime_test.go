package generator

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrimePi(t *testing.T) {
	for _, test := range []struct {
		x    int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 4},
		{8, 4},
		{9, 4},
		{10, 4},
		{11, 5},
		{12, 5},
		{13, 6},
		{14, 6},
		{1_000, 168},
		{1_000_000, 78_498},
		{1_000_000_000, 50_847_534},
		{1_000_000_000_000, 0},
	} {
		t.Run(fmt.Sprintf("PrimePi(%d)", test.x), func(t *testing.T) {
			if got := PrimePi(test.x); got != test.want {
				t.Errorf("PrimePi(%d) returned %d; want %d", test.x, got, test.want)
			}
		})
	}
}

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

			if got := primes.FactorCount(test.n); got != len(test.want) {
				t.Errorf("FactorCount(%d) returned %d; want %d", test.n, got, len(test.want))
			}

			clearCaches()

			if diff := cmp.Diff(test.want, primes.Factors(test.n)); diff != "" {
				t.Errorf("Factors(%d) returned incorrect result (-want, +got):\n%s", test.n, diff)
			}

			clearCaches()

		})
	}
}
