package maths

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMatrix(t *testing.T) {
	for _, test := range []struct {
		name          string
		matrix        [][]int
		wantDet       int
		wantTranspose [][]int
	}{
		{
			name: "2x2",
			matrix: [][]int{
				{3, 4},
				{5, 6},
			},
			wantTranspose: [][]int{
				{3, 5},
				{4, 6},
			},
			wantDet: -2,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 0, 0},
				{0, 4, 9},
				{0, 6, 4},
			},
			wantTranspose: [][]int{
				{5, 0, 0},
				{0, 4, 6},
				{0, 9, 4},
			},
			wantDet: -190,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 3, 7},
				{2, 4, 9},
				{3, 6, 4},
			},
			wantTranspose: [][]int{
				{5, 2, 3},
				{3, 4, 6},
				{7, 9, 4},
			},
			wantDet: -133,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 3, 7, 2},
				{2, 4, 9, 3},
				{3, 6, 4, 4},
				{5, 6, 7, 8},
			},
			wantTranspose: [][]int{
				{5, 2, 3, 5},
				{3, 4, 6, 6},
				{7, 9, 4, 7},
				{2, 3, 4, 8},
			},
			wantDet: -531,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := Determinant(test.matrix); test.wantDet != got {
				t.Errorf("Determinant(%v) returned %d; wantDet %d", test.matrix, got, test.wantDet)
			}

			if diff := cmp.Diff(test.wantTranspose, Transpose(test.matrix)); diff != "" {
				t.Errorf("Transpose(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}
		})
	}
}
