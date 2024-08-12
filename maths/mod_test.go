package maths

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolveMod(t *testing.T) {
	for _, test := range []struct {
		a, mod, target, want int
	}{
		{1, 7, 0, 7},
		{1, 7, 1, 1},
		{1, 7, 2, 2},
		{1, 7, 3, 3},
		{1, 7, 4, 4},
		{1, 7, 5, 5},
		{1, 7, 6, 6},

		{2, 7, 0, 7},
		{2, 7, 1, 4},
		{2, 7, 2, 1},
		{2, 7, 3, 5},
		{2, 7, 4, 2},
		{2, 7, 5, 6},
		{2, 7, 6, 3},

		{3, 7, 0, 7},
		{3, 7, 1, 5},
		{3, 7, 2, 3},
		{3, 7, 3, 1},
		{3, 7, 4, 6},
		{3, 7, 5, 4},
		{3, 7, 6, 2},

		{4, 7, 0, 7},
		{4, 7, 1, 2},
		{4, 7, 2, 4},
		{4, 7, 3, 6},
		{4, 7, 4, 1},
		{4, 7, 5, 3},
		{4, 7, 6, 5},

		{5, 7, 0, 7},
		{5, 7, 1, 3},
		{5, 7, 2, 6},
		{5, 7, 3, 2},
		{5, 7, 4, 5},
		{5, 7, 5, 1},
		{5, 7, 6, 4},

		{6, 7, 0, 7},
		{6, 7, 1, 6},
		{6, 7, 2, 5},
		{6, 7, 3, 4},
		{6, 7, 4, 3},
		{6, 7, 5, 2},
		{6, 7, 6, 1},

		{8, 7, 0, 7},
		{8, 7, 1, 1},
		{8, 7, 2, 2},
		{8, 7, 3, 3},
		{8, 7, 4, 4},
		{8, 7, 5, 5},
		{8, 7, 6, 6},
		/* Useful for commenting out tests. */
	} {
		t.Run(fmt.Sprintf("%dx = %d (mod %d)", test.a, test.target, test.mod), func(t *testing.T) {
			if diff := cmp.Diff(test.want, SolveMod(test.a, test.mod, test.target)); diff != "" {
				t.Errorf("Incorrect result (-want, +got):\n%s", diff)
			}
		})
	}
}

func TestCoprimes(t *testing.T) {
	for _, test := range []struct {
		a, b, want  int
		wantCoprime bool
	}{
		{8, 11, 1, true},
		{8, 4, 4, false},
		{8, 24, 8, false},
		{24, 8, 8, false},
		{17, 27, 1, true},
		{84, 27, 3, false},
		{90, 27, 9, false},
		{90, 54, 18, false},
		{42823, 6409, 17, false},
		{42824, 6409, 1, true},
	} {
		t.Run(fmt.Sprintf("Gcd(%d, %d)", test.a, test.b), func(t *testing.T) {
			// Test GCD(a, b) and GCD(b, a)
			if diff := cmp.Diff(test.want, Gcd(test.a, test.b)); diff != "" {
				t.Errorf("Gcd(%d, %d) returned incorrect result (-want, +got):\n%s", test.a, test.b, diff)
			}
			if diff := cmp.Diff(test.want, Gcd(test.b, test.a)); diff != "" {
				t.Errorf("Gcd(%d, %d) returned incorrect result (-want, +got):\n%s", test.b, test.a, diff)
			}

			// Test Coprime(a, b) and Coprime(b, a)
			if Coprime(test.a, test.b) != test.wantCoprime {
				t.Errorf("Coprime(%d, %d) returned incorrect result. Want %v, got %v", test.a, test.b, test.wantCoprime, !test.wantCoprime)
			}
			if Coprime(test.b, test.a) != test.wantCoprime {
				t.Errorf("Coprime(%d, %d) returned incorrect result. Want %v, got %v", test.b, test.a, test.wantCoprime, !test.wantCoprime)
			}
		})
	}
}
