package maths

import (
	"fmt"
	"testing"
)

func TestDeterminant(t *testing.T) {
	for _, test := range []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name: "2x2",
			matrix: [][]int{
				{3, 4},
				{5, 6},
			},
			want: -2,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 0, 0},
				{0, 4, 9},
				{0, 6, 4},
			},
			want: -190,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 3, 7},
				{2, 4, 9},
				{3, 6, 4},
			},
			want: -133,
		},
		{
			name: "3x3",
			matrix: [][]int{
				{5, 3, 7, 2},
				{2, 4, 9, 3},
				{3, 6, 4, 4},
				{5, 6, 7, 8},
			},
			want: -531,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			fmt.Println(test.name, "==================")
			if got := Determinant(test.matrix); test.want != got {
				t.Errorf("Determinant(%v) returned %d; want %d", test.matrix, got, test.want)
			}
		})
	}
}
