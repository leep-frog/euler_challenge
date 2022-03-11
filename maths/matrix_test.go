package maths

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestMatrix(t *testing.T) {
	for _, test := range []struct {
		name          string
		matrix        [][]int
		wantDet       int
		wantTranspose [][]int
		wantAdj       [][]int
	}{
		/*{
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
		},*/
		{
			name: "3x3",
			matrix: [][]int{
				{1, 2, 3},
				{0, 1, 4},
				{5, 6, 0},
			},
			wantTranspose: [][]int{
				{1, 0, 5},
				{2, 1, 6},
				{3, 4, 0},
			},
			wantAdj: [][]int{
				{-24, 18, 5},
				{20, -15, -4},
				{-5, 4, 1},
			},
			wantDet: 1,
		},
		/*{
			name: "4x4",
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
		/* Useful for commenting out tests. */
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := Determinant(test.matrix); test.wantDet != got {
				t.Errorf("Determinant(%v) returned %d; wantDet %d", test.matrix, got, test.wantDet)
			}

			if diff := cmp.Diff(test.wantTranspose, Transpose(test.matrix)); diff != "" {
				t.Errorf("Transpose(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(test.wantAdj, AdjugateMatrix(test.matrix)); diff != "" {
				t.Errorf("AdjugateMatrix(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}
		})
	}
}

func TestMatrixInverse(t *testing.T) {
	for _, test := range []struct {
		name          string
		matrix        [][]float64
		wantDet       float64
		wantTranspose [][]float64
		wantAdj       [][]float64
	}{
		{
			name: "2x2",
			matrix: [][]float64{
				{3, 4},
				{5, 6},
			},
			wantTranspose: [][]float64{
				{3, 5},
				{4, 6},
			},
			wantDet: -2,
		},
		{
			name: "3x3",
			matrix: [][]float64{
				{5, 0, 0},
				{0, 4, 9},
				{0, 6, 4},
			},
			wantTranspose: [][]float64{
				{5, 0, 0},
				{0, 4, 6},
				{0, 9, 4},
			},
			wantDet: -190,
		},
		{
			name: "3x3",
			matrix: [][]float64{
				{5, 3, 7},
				{2, 4, 9},
				{3, 6, 4},
			},
			wantTranspose: [][]float64{
				{5, 2, 3},
				{3, 4, 6},
				{7, 9, 4},
			},
			wantDet: -133,
		},
		{
			name: "3x3",
			matrix: [][]float64{
				{1, 2, 3},
				{0, 1, 4},
				{5, 6, 0},
			},
			wantTranspose: [][]float64{
				{1, 0, 5},
				{2, 1, 6},
				{3, 4, 0},
			},
			wantAdj: [][]float64{
				{-24, 18, 5},
				{20, -15, -4},
				{-5, 4, 1},
			},
			wantDet: 1,
		},
		{
			name: "4x4",
			matrix: [][]float64{
				{5, 3, 7, 2},
				{2, 4, 9, 3},
				{3, 6, 4, 4},
				{5, 6, 7, 8},
			},
			wantTranspose: [][]float64{
				{5, 2, 3, 5},
				{3, 4, 6, 6},
				{7, 9, 4, 7},
				{2, 3, 4, 8},
			},
			wantDet: -531,
		},
		/* Useful for commenting out tests. */
	} {
		t.Run(test.name, func(t *testing.T) {
			if got := Determinant(test.matrix); test.wantDet != got {
				t.Errorf("Determinant(%v) returned %0.2f; wantDet %0.2f", test.matrix, got, test.wantDet)
			}

			if diff := cmp.Diff(test.wantTranspose, Transpose(test.matrix)); diff != "" {
				t.Errorf("Transpose(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(IdentityMatrix[float64](len(test.matrix)), MultiplyMatrices(Inverse(test.matrix), test.matrix), cmpopts.EquateApprox(0, 0.00000000000001)); diff != "" {
				t.Errorf("Inverse*Matrix (%v) returned diff (-want, +got):\n%s", test.matrix, diff)
			}

			/*if diff := cmp.Diff(test.wantAdj, AdjugateMatrix(test.matrix)); diff != "" {
				t.Errorf("AdjugateMatrix(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}*/
		})
	}
}
