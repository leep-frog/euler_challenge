package binary

import (
	"testing"
)

func TestBinary(t *testing.T) {
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
