package binary

import (
	"fmt"
	"testing"
)

func TestBinaryBinaryOps(t *testing.T) {
	for _, test := range []struct {
		name           string
		a              int
		b              int
		wantA          string
		wantB          string
		wantDifference int
		wantBinaryDiff string
	}{
		{
			name:           "same number",
			a:              5,
			b:              5,
			wantA:          "101",
			wantB:          "101",
			wantDifference: 0,
			wantBinaryDiff: "0",
		},
		{
			name:           "Power of 2 minus one less than it",
			a:              16,
			b:              15,
			wantA:          "10000",
			wantB:          "1111",
			wantDifference: 1,
			wantBinaryDiff: "1",
		},
		{
			name:           "Power of 2 minus one less than it",
			a:              23,
			b:              9,
			wantA:          "10111",
			wantB:          "1001",
			wantDifference: 14,
			wantBinaryDiff: "1110",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			aBin := BinaryFromInt(test.a)
			if got := aBin.ToInt(); got != test.a {
				t.Errorf("BinaryFromInt(%d).ToInt() returned %d", test.a, got)
			}
			if got := aBin.String(); got != test.wantA {
				t.Errorf("BinaryFromInt(%d).ToString() returned %s; want %s", test.a, got, test.wantA)
			}

			bBin := BinaryFromInt(test.b)
			if got := bBin.ToInt(); got != test.b {
				t.Errorf("BinaryFromInt(%d).ToInt() returned %d", test.b, got)
			}
			if got := bBin.String(); got != test.wantB {
				t.Errorf("BinaryFromInt(%d).String() returned %s; want %s", test.b, got, test.wantB)
			}

			diffBin := aBin.Minus(bBin)
			if got := diffBin.ToInt(); got != test.wantDifference {
				t.Errorf("BinaryFromInt(%d).Minus(BinaryFromInt(%d).Int()) returned %d; want %d", test.a, test.b, got, test.wantDifference)
			}
			if got := diffBin.String(); got != test.wantBinaryDiff {
				t.Errorf("BinaryFromInt(%d).Minus(BinaryFromInt(%d).String()) returned %s; want %s", test.a, test.b, got, test.wantBinaryDiff)
			}
		})
	}
}

func TestUnaryBinaryOps(t *testing.T) {
	for _, test := range []struct {
		name              string
		b                 int
		wantB             string
		wantDouble        int
		wantDoublePlusOne int
		wantHalf          int
		wantIsZero        bool
	}{
		{
			name:              "Zero",
			b:                 0,
			wantB:             "0",
			wantDouble:        0,
			wantDoublePlusOne: 1,
			wantHalf:          0,
			wantIsZero:        true,
		},
		{
			name:              "One",
			b:                 1,
			wantB:             "1",
			wantDouble:        2,
			wantDoublePlusOne: 3,
			wantHalf:          0,
		},
		{
			name:              "Two",
			b:                 2,
			wantB:             "10",
			wantDouble:        4,
			wantDoublePlusOne: 5,
			wantHalf:          1,
		},
		{
			name:              "Three",
			b:                 3,
			wantB:             "11",
			wantDouble:        6,
			wantDoublePlusOne: 7,
			wantHalf:          1,
		},
		{
			name:              "Four",
			b:                 4,
			wantB:             "100",
			wantDouble:        8,
			wantDoublePlusOne: 9,
			wantHalf:          2,
		},
		{
			name:              "Five",
			b:                 5,
			wantB:             "101",
			wantDouble:        10,
			wantDoublePlusOne: 11,
			wantHalf:          2,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println("TESTING ==============", test.b)
			bin := BinaryFromInt(test.b)
			if got := bin.ToInt(); got != test.b {
				t.Errorf("BinaryFromInt(%d).ToInt() returned %d", test.b, got)
			}
			if got := bin.String(); got != test.wantB {
				t.Errorf("BinaryFromInt(%d).ToString() returned %s; want %s", test.b, got, test.wantB)
			}

			cp := bin.Copy()
			cp.Double()
			if cp.ToInt() != test.wantDouble {
				t.Errorf("BinaryFromInt(%d).Double() returned %d; want %d", test.b, cp.ToInt(), test.wantDouble)
			}

			cp = bin.Copy()
			cp.DoublePlusOne()
			if cp.ToInt() != test.wantDoublePlusOne {
				t.Errorf("BinaryFromInt(%d).DoublePlusOne() returned %d; want %d", test.b, cp.ToInt(), test.wantDoublePlusOne)
			}

			cp = bin.Copy()
			cp.Half()
			if cp.ToInt() != test.wantHalf {
				t.Errorf("BinaryFromInt(%d).Half() returned %d; want %d", test.b, cp.ToInt(), test.wantHalf)
			}

			if bin.IsZero() != test.wantIsZero {
				t.Errorf("BinaryFromInt(%d).IsZero() returned %v; want %v", test.b, bin.IsZero(), test.wantIsZero)
			}
		})
	}
}
