package fraction

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/generator"
)

func TestNew(t *testing.T) {
	for _, test := range []struct {
		n     int
		d     int
		wantN int
		wantD int
	}{
		{0, 0, 0, 0},
		// Zero denominator
		{1, 0, 1, 0},
		{-1, 0, 1, 0},
		// Zero numerator
		{0, 1, 0, 1},
		{0, -1, 0, 1},
		// Non-zero numerator and denominator
		{1, 1, 1, 1},
		{-1, -1, 1, 1},
		{-1, 1, -1, 1},
		{1, -1, -1, 1},
	} {
		t.Run(fmt.Sprintf("New(%d, %d)", test.n, test.d), func(t *testing.T) {

			if diff := cmp.Diff(New(test.n, test.d), &Fraction[int]{test.wantN, test.wantD}); diff != "" {
				t.Errorf("New(%d, %d) returned incorrect fraction (-want, +got):\n%s", test.n, test.d, diff)
			}
		})
	}
}

func TestIsTriangular(t *testing.T) {
	p := generator.Primes()
	for _, test := range []struct {
		f    *Fraction[int]
		want *Fraction[int]
	}{
		{New(6, 4), New(3, 2)},
		{New(1488, 66), New(248, 11)},
		{New(7*5*3*3*3, 3*3*7*7*7*5*8), New(3, 7*7*8)},
	} {
		t.Run(fmt.Sprintf("(%v).Simplify", test.f), func(t *testing.T) {
			if diff := cmp.Diff(test.want, Simplify(test.f.N, test.f.D, p)); diff != "" {
				t.Errorf("(%v).Simplify() returned incorrect result (-want, +got):\n%s", test.f, diff)
			}
		})
	}
}
