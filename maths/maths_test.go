package maths

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func tn(name string) string {
	return fmt.Sprintf("[maxInt = %d] %s", maxInt(), name)
}

func TestPermutations(t *testing.T) {
	for _, test := range []struct {
		name  string
		parts []string
		want  []string
	}{
		{
			name:  "small",
			parts: []string{"0", "1", "2"},
			want:  []string{"012", "021", "102", "120", "201", "210"},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := Permutations(test.parts)
			sort.Strings(got)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Permutations(%v) returned incorrect value (-want, +got):\n%s", test.parts, diff)
			}
		})
	}
}

func TestComps(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name       string
			a          *Int
			b          *Int
			wantLT     bool
			wantLTE    bool
			wantEQ     bool
			wantGTE    bool
			wantGT     bool
			wantMagLT  bool
			wantMagLTE bool
			wantMagEQ  bool
			wantMagGTE bool
			wantMagGT  bool
		}{
			{
				name:       "pos equal",
				a:          NewInt(123),
				b:          NewInt(123),
				wantLTE:    true,
				wantEQ:     true,
				wantGTE:    true,
				wantMagLTE: true,
				wantMagEQ:  true,
				wantMagGTE: true,
			},
			{
				name:       "neg equal",
				a:          NewInt(-321),
				b:          NewInt(-321),
				wantLTE:    true,
				wantEQ:     true,
				wantGTE:    true,
				wantMagLTE: true,
				wantMagEQ:  true,
				wantMagGTE: true,
			},
			{
				name:       "small pos, big pos",
				a:          NewInt(123),
				b:          NewInt(234),
				wantLT:     true,
				wantLTE:    true,
				wantMagLT:  true,
				wantMagLTE: true,
			},
			{
				name:       "big pos, small pos",
				a:          NewInt(62),
				b:          NewInt(2),
				wantGT:     true,
				wantGTE:    true,
				wantMagGT:  true,
				wantMagGTE: true,
			},
			{
				name:       "small neg, big neg",
				a:          NewInt(-123),
				b:          NewInt(-234),
				wantGT:     true,
				wantGTE:    true,
				wantMagLT:  true,
				wantMagLTE: true,
			},
			{
				name:       "big neg, small neg",
				a:          NewInt(-6234),
				b:          NewInt(-2),
				wantLT:     true,
				wantLTE:    true,
				wantMagGT:  true,
				wantMagGTE: true,
			},
			{
				name:       "small pos, big neg",
				a:          NewInt(4),
				b:          NewInt(-999),
				wantGT:     true,
				wantGTE:    true,
				wantMagLT:  true,
				wantMagLTE: true,
			},
			{
				name:       "big pos, small neg",
				a:          NewInt(4444),
				b:          NewInt(-9),
				wantGT:     true,
				wantGTE:    true,
				wantMagGT:  true,
				wantMagGTE: true,
			},
			{
				name:       "small neg, big pos",
				a:          NewInt(-239),
				b:          NewInt(8746321111),
				wantLT:     true,
				wantLTE:    true,
				wantMagLT:  true,
				wantMagLTE: true,
			},
			{
				name:       "big neg, small pos",
				a:          NewInt(-239792037),
				b:          NewInt(10101),
				wantLT:     true,
				wantLTE:    true,
				wantMagGT:  true,
				wantMagGTE: true,
			},
			{
				name:       "equal neg and pos",
				a:          NewInt(-724913),
				b:          NewInt(724913),
				wantLT:     true,
				wantLTE:    true,
				wantMagLTE: true,
				wantMagEQ:  true,
				wantMagGTE: true,
			},
			{
				name:       "equal pos and neg",
				a:          NewInt(724913),
				b:          NewInt(-724913),
				wantGT:     true,
				wantGTE:    true,
				wantMagLTE: true,
				wantMagEQ:  true,
				wantMagGTE: true,
			},
			/* Useful for commenting out tests*/
		} {
			t.Run(tn(test.name), func(t *testing.T) {
				if got := test.a.LT(test.b); got != test.wantLT {
					t.Errorf("%v.LT(%v) returned %v; want %v", test.a, test.b, got, test.wantLT)
				}

				if got := test.a.LTE(test.b); got != test.wantLTE {
					t.Errorf("%v.LTE(%v) returned %v; want %v", test.a, test.b, got, test.wantLTE)
				}

				if got := test.a.EQ(test.b); got != test.wantEQ {
					t.Errorf("%v.EQ(%v) returned %v; want %v", test.a, test.b, got, test.wantEQ)
				}

				if got := test.a.GTE(test.b); got != test.wantGTE {
					t.Errorf("%v.GTE(%v) returned %v; want %v", test.a, test.b, got, test.wantGTE)
				}

				if got := test.a.GT(test.b); got != test.wantGT {
					t.Errorf("%v.GT(%v) returned %v; want %v", test.a, test.b, got, test.wantGT)
				}

				// Magnitude comparisons
				if got := test.a.MagLT(test.b); got != test.wantMagLT {
					t.Errorf("%v.MagLT(%v) returned %v; want %v", test.a, test.b, got, test.wantMagLT)
				}

				if got := test.a.MagLTE(test.b); got != test.wantMagLTE {
					t.Errorf("%v.MagLTE(%v) returned %v; want %v", test.a, test.b, got, test.wantMagLTE)
				}

				if got := test.a.MagEQ(test.b); got != test.wantMagEQ {
					t.Errorf("%v.MagEQ(%v) returned %v; want %v", test.a, test.b, got, test.wantMagEQ)
				}

				if got := test.a.MagGTE(test.b); got != test.wantMagGTE {
					t.Errorf("%v.MagGTE(%v) returned %v; want %v", test.a, test.b, got, test.wantMagGTE)
				}

				if got := test.a.MagGT(test.b); got != test.wantMagGT {
					t.Errorf("%v.MagGT(%v) returned %v; want %v", test.a, test.b, got, test.wantMagGT)
				}
			})
		}
	}
}

func maxIs(t *testing.T) []int {
	oldMaxDigits := maxDigits
	t.Cleanup(func() {
		maxIntCached = 0
		maxDigits = oldMaxDigits
	})
	return []int{1, 2, 3, oldMaxDigits}
}

func setMax(md int) {
	maxIntCached = 0
	maxDigits = md
}

func TestIntPlus(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for idx, test := range []struct {
			input      []int64
			want       *Int
			wantString string
		}{
			{
				input:      []int64{0, 4, 0, 8, 0},
				want:       NewInt(12),
				wantString: "12",
			},
			{
				input:      []int64{0, 4, -3, 8, 0},
				want:       NewInt(9),
				wantString: "9",
			},
			{
				input:      []int64{0, -3, 4, 0, 8},
				want:       NewInt(9),
				wantString: "9",
			},
			{
				input:      []int64{100, -87},
				want:       NewInt(13),
				wantString: "13",
			},
			{
				input:      []int64{-3010, 220},
				want:       NewInt(-2790),
				wantString: "-2790",
			},
			{
				input:      []int64{220, -3010},
				want:       NewInt(-2790),
				wantString: "-2790",
			},
			{
				input:      []int64{12345678, -12345600},
				want:       NewInt(78),
				wantString: "78",
			},
			{
				input:      []int64{10000, 0},
				want:       NewInt(10000),
				wantString: "10000",
			},
		} {
			t.Run(tn(fmt.Sprintf("Add Test %d", idx)), func(t *testing.T) {
				got := SumI(test.input...)
				if diff := cmp.Diff(test.want, got, CmpOpts()...); diff != "" {
					t.Errorf("Sum(%v) produced incorrect result (-want, +got):\n%s", test.input, diff)
				}
			})
		}
	}
}

func TestIntToInt(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name  string
			fName string
			f     func(int) int
			input int
			want  int
		}{
			// Abs
			{
				name:  "Absolute value positive number",
				fName: "Abs",
				f:     Abs,
				input: 5,
				want:  5,
			},
			{
				name:  "Absolute value zero",
				fName: "Abs",
				f:     Abs,
				input: 0,
				want:  0,
			},
			{
				name:  "Absolute value negative number",
				fName: "Abs",
				f:     Abs,
				input: -14,
				want:  14,
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				if got := test.f(test.input); got != test.want {
					t.Errorf("%s(%d) returned %d; want %d", test.fName, test.input, got, test.want)
				}
			})
		}
	}
}

func TestIntToBool(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name  string
			fName string
			f     func(int) bool
			input int
			want  bool
		}{
			// IsSquare
			{
				name:  "IsSquare with square",
				fName: "IsSquare",
				f:     IsSquare,
				input: 256,
				want:  true,
			},
			{
				name:  "IsSquare with non-square",
				fName: "IsSquare",
				f:     IsSquare,
				input: 257,
			},
			{
				name:  "IsSquare with zero",
				fName: "IsSquare",
				f:     IsSquare,
				want:  true,
			},
			{
				name:  "IsSquare with negative number",
				fName: "IsSquare",
				f:     IsSquare,
				input: -4,
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				if got := test.f(test.input); got != test.want {
					t.Errorf("%s(%d) returned %v; want %v", test.fName, test.input, got, test.want)
				}
			})
		}
	}
}

func TestDiv(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name string
			// Test the following:
			// a / b = c + wantRem
			a       *Int
			b       *Int
			c       *Int
			wantRem *Int
			// If false, test the following as well
			// a / c = b + wantRem
			noReverseTest bool
		}{
			{
				name: "small by small",
				a:    NewInt(189),
				b:    NewInt(9),
				c:    NewInt(21),
			},
			{
				name: "small neg by small",
				a:    NewInt(-189),
				b:    NewInt(9),
				c:    NewInt(-21),
			},
			{
				name:    "small by small with remainder",
				a:       NewInt(194),
				b:       NewInt(9),
				c:       NewInt(21),
				wantRem: NewInt(5),
			},
			{
				name: "small by 2",
				a:    MustIntFromString("24"),
				b:    NewInt(2),
				c:    MustIntFromString("12"),
			},
			{
				name: "big by small",
				a:    MustIntFromString("335367096786357081410764800000"),
				b:    NewInt(2),
				c:    MustIntFromString("167683548393178540705382400000"),
			},
			{
				name: "bigger by small with remainder",
				a:    MustIntFromString("335367096786357081410764800000"),
				b:    NewInt(222),
				c:    MustIntFromString("1510662598136743609958400000"),
			},
			{
				name:    "biggerer by small-ish",
				a:       MustIntFromString("335367096786357081410764800000"),
				b:       NewInt(22222),
				c:       MustIntFromString("15091670272088789551379929"),
				wantRem: NewInt(17762),
			},
			{
				name:    "biggest by small-ish",
				a:       MustIntFromString("335367096786357081410764800000"),
				b:       NewInt(23456),
				c:       MustIntFromString("14297710470086846922355252"),
				wantRem: NewInt(9088),
			},
			{
				name:          "big by big",
				a:             MustIntFromString("335367096786357081410764800000"),
				b:             MustIntFromString("67734834945737458"),
				c:             MustIntFromString("4951176112778"),
				wantRem:       MustIntFromString("61261549777761676"),
				noReverseTest: true,
			},
			/* Useful for commenting out tests. */
		} {
			t.Run(tn(test.name), func(t *testing.T) {
				quo, rem := test.a.Div(test.b)
				if diff := cmp.Diff(test.c, quo, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect quotient (-want, +got):\n%s", test.a, test.b, diff)
				}
				if diff := cmp.Diff(test.wantRem, rem, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect remainder (-want, +got):\n%s", test.a, test.b, diff)
				}

				// Reverse should also work
				if test.noReverseTest {
					return
				}
				quo, rem = test.a.Div(test.c)
				if diff := cmp.Diff(test.b, quo, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect quotient (-want, +got):\n%s", test.a, test.c, diff)
				}
				if diff := cmp.Diff(test.wantRem, rem, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect remainder (-want, +got):\n%s", test.a, test.c, diff)
				}
			})
		}
	}
}

func TestDivInt(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name    string
			a       *Int
			b       uint16
			want    *Int
			wantRem uint16
		}{
			{
				name: "simple",
				a:    NewInt(189),
				b:    9,
				want: NewInt(21),
			},
			{
				name:    "simple with remainder",
				a:       NewInt(194),
				b:       9,
				want:    NewInt(21),
				wantRem: 5,
			},
			{
				name: "big",
				a:    MustIntFromString("335367096786357081410764800000"),
				b:    2,
				want: MustIntFromString("167683548393178540705382400000"),
			},
			{
				name: "bigger",
				a:    MustIntFromString("335367096786357081410764800000"),
				b:    222,
				want: MustIntFromString("1510662598136743609958400000"),
			},
			{
				name:    "biggerer",
				a:       MustIntFromString("335367096786357081410764800000"),
				b:       22222,
				want:    MustIntFromString("15091670272088789551379929"),
				wantRem: 17762,
			},
			{
				name:    "biggest",
				a:       MustIntFromString("335367096786357081410764800000"),
				b:       23456,
				want:    MustIntFromString("14297710470086846922355252"),
				wantRem: 9088,
			},
			/* Useful for commenting out tests. */
		} {
			t.Run(tn(test.name), func(t *testing.T) {
				if diff := cmp.Diff(test.want, test.a.DivInt(test.b), CmpOpts()...); diff != "" {
					t.Errorf("DivInt(%v, %d) returned incorrect value (-want, +got):\n%s", test.a, test.b, diff)
				}

				if diff := cmp.Diff(test.wantRem, test.a.ModInt(test.b)); diff != "" {
					t.Errorf("ModInt(%v, %d) returned incorrect value (-want, +got):\n%s", test.a, test.b, diff)
				}

				bInt := NewInt(int64(test.b))
				remInt := NewInt(int64(test.wantRem))
				quo, rem := test.a.Div(bInt)
				if diff := cmp.Diff(test.want, quo, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect quotient (-want, +got):\n%s", test.a, bInt, diff)
				}
				if diff := cmp.Diff(remInt, rem, CmpOpts()...); diff != "" {
					t.Errorf("Div(%v, %v) returned incorrect remainder (-want, +got):\n%s", test.a, bInt, diff)
				}
			})
		}
	}
}

func TestIntsToInt(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			name  string
			fName string
			f     func(...int) int
			input []int
			want  int
		}{
			// Min
			{
				name:  "Min value",
				fName: "Min",
				f:     Min,
				input: []int{0, 4, -2, 9},
				want:  -2,
			},
			{
				name:  "Min value with single input",
				fName: "Min",
				f:     Min,
				input: []int{9},
				want:  9,
			},
			{
				name:  "Min value with no input",
				fName: "Min",
				f:     Min,
				want:  0,
			},
			// Max
			{
				name:  "Max value",
				fName: "Max",
				f:     Max,
				input: []int{0, 4, -2, 9},
				want:  9,
			},
			{
				name:  "Max value with single input",
				fName: "Max",
				f:     Max,
				input: []int{3},
				want:  3,
			},
			{
				name:  "Max value with no input",
				fName: "Max",
				f:     Max,
				want:  0,
			},
		} {
			t.Run(test.name, func(t *testing.T) {
				if got := test.f(test.input...); got != test.want {
					t.Errorf("%s(%d) returned %d; want %d", test.fName, test.input, got, test.want)
				}
			})
		}
	}
}
