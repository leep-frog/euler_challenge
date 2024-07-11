package maths

/*func TestMatrix(t *testing.T) {
	for _, test := range []struct {
		name          string
		matrix        [][]float64
		wantDet       float64
		wantTranspose [][]float64
		wantRotate    [][]float64
	}{
		{
			name: "1x1",
			matrix: [][]float64{
				{3},
			},
			wantTranspose: [][]float64{
				{3},
			},
			wantRotate: [][]float64{
				{3},
			},
			wantDet: 3,
		},
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
			wantRotate: [][]float64{
				{5, 3},
				{6, 4},
			},
			wantDet: -2,
		},
		{
			name: "2x2",
			matrix: [][]float64{
				{4, 9},
				{6, 4},
			},
			wantTranspose: [][]float64{
				{4, 6},
				{9, 4},
			},
			wantRotate: [][]float64{
				{6, 4},
				{4, 9},
			},
			wantDet: -38,
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
			wantRotate: [][]float64{
				{0, 0, 5},
				{6, 4, 0},
				{4, 9, 0},
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
			wantRotate: [][]float64{
				{3, 2, 5},
				{6, 4, 3},
				{4, 9, 7},
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
			wantRotate: [][]float64{
				{5, 0, 1},
				{6, 1, 2},
				{0, 4, 3},
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
			wantRotate: [][]float64{
				{5, 3, 2, 5},
				{6, 6, 4, 3},
				{7, 4, 9, 7},
				{8, 4, 3, 2},
			},
			wantDet: -531,
		},
		/* Useful for commenting out tests. * /
	} {
		t.Run(test.name, func(t *testing.T) {
			m := BiggifyMatrix(test.matrix)

			if got := Determinant(m); big.NewRat(0, 1).SetFloat64(test.wantDet).Cmp(got) != 0 {
				t.Errorf("Determinant(%v) returned %v; wantDet %v", test.matrix, got, test.wantDet)
			}

			if diff := cmp.Diff(test.wantTranspose, SmallifyMatrix(Transpose(m))); diff != "" {
				t.Errorf("Transpose(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(test.wantRotate, SmallifyMatrix(Rotate(m))); diff != "" {
				t.Errorf("Rotate(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
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
	}{
		{
			name: "1x1",
			matrix: [][]float64{
				{3},
			},
			wantTranspose: [][]float64{
				{3},
			},
			wantDet: 3,
		},
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
		/ * Useful for commenting out tests. * /
	} {
		t.Run(test.name, func(t *testing.T) {
			m := BiggifyMatrix(test.matrix)

			opts := []cmp.Option{
				cmpopts.EquateApprox(0, 0.00000000000001),
			}

			d, _ := Determinant(m).Float64()
			if diff := cmp.Diff(test.wantDet, float64(d)); diff != "" {
				t.Errorf("Determinant(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(test.wantTranspose, SmallifyMatrix(Transpose(m))); diff != "" {
				t.Errorf("Transpose(%v) produced diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(SmallifyMatrix(IdentityMatrix(len(test.matrix))), SmallifyMatrix(MultiplyMatrices(Inverse(m), m)), opts...); diff != "" {
				t.Errorf("Inverse*Matrix (%v) returned diff (-want, +got):\n%s", test.matrix, diff)
			}

			if diff := cmp.Diff(SmallifyMatrix(IdentityMatrix(len(test.matrix))), SmallifyMatrix(MultiplyMatrices(m, Inverse(m))), cmpopts.EquateApprox(0, 0.00000000000001)); diff != "" {
				t.Errorf("Matrix*Inverse (%v) returned diff (-want, +got):\n%s", test.matrix, diff)
			}
		})
	}
}
*/
