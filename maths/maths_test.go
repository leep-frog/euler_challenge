package maths

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/parse"
)

func tn(name string) string {
	return fmt.Sprintf("[maxInt = %d] %s", maxInt(), name)
}

func TestPandigital(t *testing.T) {
	for _, test := range []struct {
		v    int
		want bool
	}{
		{
			v:    1,
			want: true,
		},
		{
			v:    1234,
			want: true,
		},
		{
			v:    35124,
			want: true,
		},
		{
			v: 351241,
		},
		{
			v: 350124,
		},
		{
			v: 13,
		},
	} {
		t.Run(fmt.Sprintf("Pandigital_%d", test.v), func(t *testing.T) {
			if got := Pandigital(test.v); got != test.want {
				t.Errorf("Pandigital(%d) returned %v; want %v", test.v, got, test.want)
			}
		})
	}
}

func TestPalindromes(t *testing.T) {
	for _, test := range []struct {
		n    int
		want []int
	}{
		{},
		{
			n:    1,
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			n:    2,
			want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99},
		},
		{
			n: 3,
			want: []int{
				101, 111, 121, 131, 141, 151, 161, 171, 181, 191,
				202, 212, 222, 232, 242, 252, 262, 272, 282, 292,
				303, 313, 323, 333, 343, 353, 363, 373, 383, 393,
				404, 414, 424, 434, 444, 454, 464, 474, 484, 494,
				505, 515, 525, 535, 545, 555, 565, 575, 585, 595,
				606, 616, 626, 636, 646, 656, 666, 676, 686, 696,
				707, 717, 727, 737, 747, 757, 767, 777, 787, 797,
				808, 818, 828, 838, 848, 858, 868, 878, 888, 898,
				909, 919, 929, 939, 949, 959, 969, 979, 989, 999,
			},
		},
	} {
		t.Run(fmt.Sprintf("Palindromes_%d", test.n), func(t *testing.T) {
			if diff := cmp.Diff(test.want, Palindromes(test.n)); diff != "" {
				t.Errorf("Palindromes(%d) returned incorrect values (-want, +got):\n%s", test.n, diff)
			}
		})
	}
}

func TestToBinary(t *testing.T) {
	for _, test := range []struct {
		i    int
		want string
	}{
		{
			i:    585,
			want: "1001001001",
		},
		{
			i:    13,
			want: "1101",
		},
	} {
		t.Run(fmt.Sprintf("ToBinary(%d)", test.i), func(t *testing.T) {
			wantB := NewBinary(test.want)
			got := ToBinary(test.i)

			if diff := cmp.Diff(test.want, got.String(), CmpOpts()...); diff != "" {
				t.Errorf("ToBinary(%d) produced incorrect string:\n%s", test.i, diff)
			}

			if diff := cmp.Diff(wantB, got, CmpOpts()...); diff != "" {
				t.Errorf("ToBinary(%d) produced incorrect struct:\n%s", test.i, diff)
			}

			if diff := cmp.Diff(test.i, got.ToInt()); diff != "" {
				t.Errorf("ToBinary(%d).ToInt() returned incorrect int:\n%s", test.i, diff)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	for _, test := range []struct {
		name    string
		parts   []string
		want    []string
		wantRot []string
		wantSet [][]int
	}{
		{
			name:    "small",
			parts:   []string{"0", "1", "2"},
			want:    []string{"012", "021", "102", "120", "201", "210"},
			wantRot: []string{"012", "120", "201"},
			wantSet: [][]int{
				{0},
				{0, 1},
				{0, 1, 2},
				{0, 2},
				{1},
				{1, 2},
				{2},
			},
		},
		{
			name:    "duplicates",
			parts:   []string{"1", "0", "1"},
			want:    []string{"011", "101", "110"},
			wantRot: []string{"101", "011", "110"},
			wantSet: [][]int{
				{1},
				{1, 1},
				{1, 1, 0},
				{1, 0},
				{0},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			gots := Permutations(test.parts)
			var got []string
			for _, g := range gots {
				got = append(got, strings.Join(g, ""))
			}
			sort.Strings(got)
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Permutations(%v) returned incorrect value (-want, +got):\n%s", test.parts, diff)
			}

			gotRot := Rotations(test.parts)
			if diff := cmp.Diff(test.wantRot, gotRot); diff != "" {
				t.Errorf("Rotations(%v) returned incorrect value (-want, +got):\n%s", test.parts, diff)
			}

			var iParts []int
			for _, p := range test.parts {
				iParts = append(iParts, parse.Atoi(p))
			}

			gotSet := ChooseAllSets(iParts)
			if diff := cmp.Diff(test.wantSet, gotSet); diff != "" {
				t.Errorf("Sets(%v) returned incorrect values (-want, +got):\n%s", test.parts, diff)
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

				if got := test.a.NEQ(test.b); got != !test.wantEQ {
					t.Errorf("%v.NEQ(%v) returned %v; want %v", test.a, test.b, got, !test.wantEQ)
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

				if got := test.a.MagNEQ(test.b); got != !test.wantMagEQ {
					t.Errorf("%v.MagNEQ(%v) returned %v; want %v", test.a, test.b, got, !test.wantMagEQ)
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
			input      []int
			want       *Int
			wantString string
		}{
			{
				input:      []int{0, 4, 0, 8, 0},
				want:       NewInt(12),
				wantString: "12",
			},
			{
				input:      []int{0, 4, -3, 8, 0},
				want:       NewInt(9),
				wantString: "9",
			},
			{
				input:      []int{0, -3, 4, 0, 8},
				want:       NewInt(9),
				wantString: "9",
			},
			{
				input:      []int{100, -87},
				want:       NewInt(13),
				wantString: "13",
			},
			{
				input:      []int{-3010, 220},
				want:       NewInt(-2790),
				wantString: "-2790",
			},
			{
				input:      []int{220, -3010},
				want:       NewInt(-2790),
				wantString: "-2790",
			},
			{
				input:      []int{12345678, -12345600},
				want:       NewInt(78),
				wantString: "78",
			},
			{
				input:      []int{10000, 0},
				want:       NewInt(10000),
				wantString: "10000",
			},
		} {
			t.Run(tn(fmt.Sprintf("Add Test %d", idx)), func(t *testing.T) {
				var inputs []*Int
				for _, in := range test.input {
					inputs = append(inputs, NewInt(in))
				}
				got := Sum(inputs...)
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
				f:     Abs[int],
				input: 5,
				want:  5,
			},
			{
				name:  "Absolute value zero",
				fName: "Abs",
				f:     Abs[int],
				input: 0,
				want:  0,
			},
			{
				name:  "Absolute value negative number",
				fName: "Abs",
				f:     Abs[int],
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

func TestIntFromString(t *testing.T) {
	for _, md := range maxIs(t) {
		setMax(md)
		for _, test := range []struct {
			s      string
			want   *Int
			wantOK bool
		}{
			{
				s:      "0",
				want:   Zero(),
				wantOK: true,
			},
			{
				s:      "1",
				want:   One(),
				wantOK: true,
			},
			{
				s:      "01",
				want:   One(),
				wantOK: true,
			},
			{
				s:      "000001",
				want:   One(),
				wantOK: true,
			},
			{
				s:      "001001",
				want:   NewInt(1001),
				wantOK: true,
			},
			{
				s:      "67734834945737458",
				want:   NewInt(67734834945737458),
				wantOK: true,
			},
		} {
			t.Run(tn(fmt.Sprintf("IntFromString(%s)", test.s)), func(t *testing.T) {
				got, ok := IntFromString(test.s)
				if diff := cmp.Diff(test.wantOK, ok, CmpOpts()...); diff != "" {
					t.Errorf("IntFromString(%s) returned incorrect OK value (-want, +got):\n%s", test.s, diff)
				}
				if diff := cmp.Diff(test.want, got, CmpOpts()...); diff != "" {
					t.Errorf("IntFromString(%s) returned incorrect value (-want, +got):\n%s", test.s, diff)
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
				quo, rem := test.a.Divide(test.b)
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
				quo, rem = test.a.Divide(test.c)
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
			b       int
			want    *Int
			wantRem int
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

				bInt := NewInt(test.b)
				remInt := NewInt(test.wantRem)
				quo, rem := test.a.Divide(bInt)
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
				f:     Min[int],
				input: []int{0, 4, -2, 9},
				want:  -2,
			},
			{
				name:  "Min value with single input",
				fName: "Min",
				f:     Min[int],
				input: []int{9},
				want:  9,
			},
			{
				name:  "Min value with no input",
				fName: "Min",
				f:     Min[int],
				want:  0,
			},
			// Max
			{
				name:  "Max value",
				fName: "Max",
				f:     Max[int],
				input: []int{0, 4, -2, 9},
				want:  9,
			},
			{
				name:  "Max value with single input",
				fName: "Max",
				f:     Max[int],
				input: []int{3},
				want:  3,
			},
			{
				name:  "Max value with no input",
				fName: "Max",
				f:     Max[int],
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

func TestRomanNumerals(t *testing.T) {
	for _, test := range []struct {
		numeral    string
		simplified string
		decimal    int
	}{
		{"I", "I", 1},
		{"II", "II", 2},
		{"III", "III", 3},
		{"IIII", "IV", 4},
		{"IIIII", "V", 5},
		{"IIIIII", "VI", 6},
		{"IIIIIII", "VII", 7},
		{"IIIIIIII", "VIII", 8},
		{"IIIIIIIII", "IX", 9},
		{"IIIIIIIIII", "X", 10},
		{"IIIIIIIIIII", "XI", 11},
		{"IIIIIIIIIIII", "XII", 12},
		{"IIIIIIIIIIIII", "XIII", 13},
	} {
		t.Run(fmt.Sprintf("RomanNumeral:%s=%d", test.numeral, test.decimal), func(t *testing.T) {
			if diff := cmp.Diff(test.simplified, RomanNumeral(test.decimal).String()); diff != "" {
				t.Errorf("RomanNumeral(%d) produced diff (-want, +got):\n%s", test.decimal, diff)
			}

			if diff := cmp.Diff(test.decimal, NumeralFromString(test.numeral).ToInt()); diff != "" {
				t.Errorf("NumeralFromString(%s) produced diff (-want, +got):\n%s", test.numeral, diff)
			}

			if diff := cmp.Diff(test.decimal, NumeralFromString(test.simplified).ToInt()); diff != "" {
				t.Errorf("NumeralFromString(%s) produced diff (-want, +got):\n%s", test.simplified, diff)
			}
		})
	}
}

func TestChooseSets(t *testing.T) {
	ints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, test := range []struct {
		parts []string
		n     int
		want  [][]string
	}{
		{},
		{
			parts: ints,
		},
		{
			parts: ints,
			n:     1,
			want: [][]string{
				{"0"},
				{"1"},
				{"2"},
				{"3"},
				{"4"},
				{"5"},
				{"6"},
				{"7"},
				{"8"},
				{"9"},
			},
		},
		{
			parts: []string{"1", "2", "3", "4", "5"},
			n:     2,
			want: [][]string{
				{"1", "2"},
				{"1", "3"},
				{"1", "4"},
				{"1", "5"},
				{"2", "3"},
				{"2", "4"},
				{"2", "5"},
				{"3", "4"},
				{"3", "5"},
				{"4", "5"},
			},
		},
	} {
		t.Run(fmt.Sprintf("ChooseSetsOfLength(%v, %d)", test.parts, test.n), func(t *testing.T) {
			got := ChooseSetsOfLength(test.parts, test.n)
			sort.SliceStable(got, func(i, j int) bool { return strings.Join(got[i], "") < strings.Join(got[j], "") })
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("ChooseSets(%v, %d) produced diff (-want, +got):\n%s", test.parts, test.n, diff)
			}
		})
	}
}
