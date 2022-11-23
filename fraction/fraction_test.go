package fraction

import (
	"fmt"
	"math/big"
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
			// Test New
			if diff := cmp.Diff(&Fraction{test.wantN, test.wantD}, New(test.n, test.d)); diff != "" {
				t.Errorf("New(%d, %d) returned incorrect fraction (-want, +got):\n%s", test.n, test.d, diff)
			}

			// Test NewI
			fI := NewRational(test.n, test.d)
			wantI := &Rational{}
			if test.wantD != 0 {
				wantI = &Rational{big.NewRat(int64(test.wantN), int64(test.wantD))}
			}
			fmt.Println("FI", fI, wantI)
			if diff := cmp.Diff(wantI, fI, CmpOpts()...); diff != "" {
				t.Errorf("NewI(%d, %d) returned incorrect fraction (-want, +got):\n%s", test.n, test.d, diff)
			}
		})
	}
}

// Also tests Times and Reciprocal
func TestDiv(t *testing.T) {
	for _, test := range []struct {
		name string
		f    *Fraction
		g    *Fraction
		want *Fraction
	}{
		{
			"Divide by 1",
			New(1, 1),
			New(1, 1),
			New(1, 1),
		},
		{
			"Divide by integers",
			New(9, 1),
			New(71, 1),
			New(9, 71),
		},
		{
			"Divide by integer reciprocal",
			New(9, 1),
			New(1, 71),
			New(9*71, 1),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.f.Div(test.g)); diff != "" {
				t.Errorf("%v.Div(%v) returned incorrect fraction (-want, +got):\n%s", test.f, test.g, diff)
			}

			fi := test.f.ToRational()
			gi := test.g.ToRational()
			wantI := test.want.ToRational()
			if diff := cmp.Diff(wantI, fi.Div(gi), CmpOpts()...); diff != "" {
				t.Errorf("%v.Div(%v) returned incorrect fraction (-want, +got):\n%s", test.f, test.g, diff)
			}
		})
	}
}

// Also tests plus
func TestMinus(t *testing.T) {
	for _, test := range []struct {
		name string
		f    *Fraction
		g    *Fraction
		want *Fraction
	}{
		{
			"1 - 1",
			New(1, 1),
			New(1, 1),
			New(0, 1),
		},
		{
			"Same fraction",
			New(9, 1),
			New(9, 1),
			New(0, 1),
		},
		{
			"Same number",
			New(9, 3),
			New(18, 6),
			New(0, 18),
		},
		{
			"Different denominators",
			New(1, 3),
			New(1, 24),
			New(7*3, 24*3),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.f.Minus(test.g)); diff != "" {
				t.Errorf("%v.Minus(%v) returned incorrect fraction (-want, +got):\n%s", test.f, test.g, diff)
			}

			fi := test.f.ToRational()
			gi := test.g.ToRational()
			wantI := test.want.ToRational()
			if diff := cmp.Diff(wantI, fi.Minus(gi), CmpOpts()...); diff != "" {
				t.Errorf("%v.Minus(%v) returned incorrect fraction (-want, +got):\n%s", test.f, test.g, diff)
			}
		})
	}
}

func TestLT(t *testing.T) {
	for _, test := range []struct {
		name        string
		f           *Fraction
		g           *Fraction
		want        bool
		wantReverse bool
	}{
		{
			name: "same number",
			f:    New(1, 1),
			g:    New(1, 1),
		},
		{
			name:        "different numbers",
			f:           New(2, 1),
			g:           New(1, 1),
			wantReverse: true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.f.LT(test.g)); diff != "" {
				t.Errorf("(%v).LT(%v) returned wrong value (-want, +got):\n%s", test.f, test.g, diff)
			}
			if diff := cmp.Diff(test.wantReverse, test.g.LT(test.f)); diff != "" {
				t.Errorf("(%v).LT(%v) returned wrong value (-want, +got):\n%s", test.g, test.f, diff)
			}

			fi := test.f.ToRational()
			gi := test.g.ToRational()
			if diff := cmp.Diff(test.want, fi.LT(gi)); diff != "" {
				t.Errorf("(%v).LT(%v) returned wrong value (-want, +got):\n%s", test.f, test.g, diff)
			}
			if diff := cmp.Diff(test.wantReverse, gi.LT(fi)); diff != "" {
				t.Errorf("(%v).LT(%v) returned wrong value (-want, +got):\n%s", gi, fi, diff)
			}

		})
	}
}

func TestSimplify(t *testing.T) {
	p := generator.Primes()
	for _, test := range []struct {
		f    *Fraction
		want *Fraction
	}{
		// Negative numbers
		{
			&Fraction{-1, -1},
			&Fraction{1, 1},
		},
		{
			&Fraction{-1, 1},
			&Fraction{-1, 1},
		},
		{
			&Fraction{1, -1},
			&Fraction{-1, 1},
		},
		{
			&Fraction{1, 1},
			&Fraction{1, 1},
		},
		// Zero over zero
		{
			New(0, 0),
			New(0, 0),
		},
		// number over zero
		{
			New(3, 0),
			New(1, 0),
		},
		{
			New(-3, 0),
			New(1, 0),
		},
		// Zero over number
		{
			New(0, 3),
			New(0, 1),
		},
		{
			New(0, -3),
			New(0, 1),
		},
		// Regular fractions
		{
			New(4, 3),
			New(4, 3),
		},
		{
			New(4, 2),
			New(2, 1),
		},
		{
			New(2, 4),
			New(1, 2),
		},
		{
			New(2*2*2*2*2*7, 2*2*2),
			New(2*2*7, 1),
		},
		{
			New(2*2*2*2*2*3*7, 2*2*2*7*13),
			New(2*2*3, 13),
		},
		/* Useful for commenting out tests. */
	} {
		t.Run(fmt.Sprintf("(%v).Simplify()", test.f), func(t *testing.T) {
			if diff := cmp.Diff(test.want, Simplify(test.f.N, test.f.D, p)); diff != "" {
				t.Errorf("Simplify(%v) returned incorrect fraction (-want, +got):\n%s", test.f, diff)
			}

			if diff := cmp.Diff(test.want, test.f.Copy().Simplify(p)); diff != "" {
				t.Errorf("(%v).Simplify() returned incorrect fraction (-want, +got):\n%s", test.f, diff)
			}

			// Don't use Copy() which fixed signs
			if diff := cmp.Diff(test.want, (&Fraction{test.f.N, test.f.D}).Simplify(p)); diff != "" {
				t.Errorf("(%v).Simplify() returned incorrect fraction (-want, +got):\n%s", test.f, diff)
			}
		})
	}
}

func TestIsTriangular(t *testing.T) {
	p := generator.Primes()
	for _, test := range []struct {
		f    *Fraction
		want *Fraction
	}{
		{New(6, 4), New(3, 2)},
		{New(1488, 66), New(248, 11)},
		{New(7*5*3*3*3, 3*3*7*7*7*5*8), New(3, 7*7*8)},
	} {
		t.Run(fmt.Sprintf("(%v).Simplify", test.f), func(t *testing.T) {
			if diff := cmp.Diff(test.want, test.f.Simplify(p)); diff != "" {
				t.Errorf("(%v).Simplify() returned incorrect result (-want, +got):\n%s", test.f, diff)
			}
		})
	}
}
