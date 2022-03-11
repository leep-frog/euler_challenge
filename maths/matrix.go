package maths

import (
	"github.com/leep-frog/euler_challenge/linkedlist"
)

/*type Matrix[T Mathable] struct {
	values *[][]T
	rows []int
	cols []int

	determinantCache map[string]T
}

func NewMatrix[T Mathable](values [][]T) *Matrix[T] {
	var rows, cols []int
	for i := range values {
		rows = append(rows, i)
	}
	if len(values) > 0 {
	for i := range values[0] {
		cols = append(cols, i)
	}
}
	return &Matrix[T]{&values, rows, cols, map[string]T{}}
}*/

func IdentityMatrix[T Mathable](n int) [][]T {
	var m [][]T
	for i := 0; i < n; i++ {
		r := make([]T, n)
		r[i] = 1
		m = append(m, r)
	}
	return m
}

func MultiplyMatrices[T Mathable](this, that [][]T) [][]T {
	if len(this) == 0 || len(that) == 0 {
		panic("cannot multiply empty matrices")
	}

	if len(this[0]) != len(that) {
		panic("cannot multiply matrices with unmatched dimensions")
	}

	var result [][]T
	for iThis := 0; iThis < len(this); iThis++ {
		var row []T
		for iThat := 0; iThat < len(that[0]); iThat++ {
			var sum T
			for j := 0; j < len(this[0]); j++ {
				sum += this[iThis][j]*that[j][iThat]
			}
			row = append(row, sum)
		}
		result = append(result, row)
	}
	return result
}

func Transpose[T Mathable](matrix [][]T) [][]T {
	var m [][]T
	if len(matrix) == 0 {
		return m
	}
	for j := range Range(len(matrix[0])) {
		var col []T
		for i := range Range(len(matrix)) {
			col = append(col, matrix[i][j])
		}
		m = append(m, col)
	}
	return m
}

func Inverse[T Mathable](matrix [][]T) [][]T {
	im := AdjugateMatrix(matrix)
	d := Determinant(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			im[i][j] /= d
		}
	}
	return im
}

func AdjugateMatrix[T Mathable](matrix [][]T) [][]T {
	//d := Determinant(matrix)
	matrix = Transpose(matrix)

	nRows := len(matrix)
	nCols := len(matrix[0])
	if nRows != nCols {
		panic("can only get adjugate of a square matrix")
	}

	var rs, cs []int
	for i := range Range(nRows) {
		rs = append(rs, i)
	}
	for i := range Range(nCols) {
		cs = append(cs, i)
	}

	rows, cols := linkedlist.NewList(rs), linkedlist.NewList(cs)

	var adj [][]T
	firstRow := true
	for r := rows; r != nil; r = r.Next {
		rHead := rows
		if firstRow {
			rHead = rows.Next
		}
		var adjRow []T
		firstCol := true
		for c := cols; c != nil; c = c.Next {
			cHead := cols
			if firstCol {
				cHead = cols.Next
			}

			// Temporarily remove nodes
			if r.Prev != nil {
				r.Prev.Next = r.Next
			}
			if r.Next != nil {
				r.Next.Prev = r.Prev
			}
			if c.Prev != nil {
				c.Prev.Next = c.Next
			}
			if c.Next != nil {
				c.Next.Prev = c.Prev
			}

			//fmt.Println(c.Value, r.Value, rHead, cHead)
			d := determinant(matrix, nRows - 1, nCols - 1, rHead, cHead,)
			if (r.Value + c.Value) % 2 == 1 {
				d *= -1
			}
			adjRow = append(adjRow, d)

			// Re-add removed nodes
			if r.Prev != nil {
				r.Prev.Next = r
			}
			if r.Next != nil {
				r.Next.Prev = r
			}
			if c.Prev != nil {
				c.Prev.Next = c
			}
			if c.Next != nil {
				c.Next.Prev = c
			}

			firstCol = false
		}
		adj = append(adj, adjRow)
		firstRow = false
	}
	return adj

	//return determinant(matrix, nRows, nCols, rows, cols)
}

func Determinant[T Mathable](matrix [][]T) T {
	if len(matrix) == 0 {
		panic("can't get determinant of an empty matrix")
	}
	nRows := len(matrix)
	nCols := len(matrix[0])
	if nRows != nCols {
		panic("can only get determinant of a square matrix")
	}

	var rs, cs []int
	for i := range Range(nRows) {
		rs = append(rs, i)
	}
	for i := range Range(nCols) {
		cs = append(cs, i)
	}

	return determinant(matrix, nRows, nCols, linkedlist.NewList(rs), linkedlist.NewList(cs))
}

func determinant[T Mathable](matrix [][]T, nRows, nCols int, rows, cols *linkedlist.Node[int]) T {
	// TODO: cache already computed determinants
	if nRows == 1 {
		return matrix[rows.Value][cols.Value]
	}

	pos := true
	var det T
		firstCol := true
		rHead := rows.Next
		for col := cols; col != nil; col = col.Next {
			cell := matrix[rows.Value][col.Value]
			cHead := cols
			if firstCol {
				cHead = cHead.Next
			}

			if col.Prev != nil {
				col.Prev.Next = col.Next
			}
			if col.Next != nil {
				col.Next.Prev = col.Prev
			}
			
			subDet := cell * determinant(matrix, nRows - 1, nCols - 1, rHead, cHead)
			if pos {
				det += subDet
			} else {
				det -= subDet
			}
			pos = !pos

			if col.Prev != nil {
				col.Prev.Next = col
			}
			if col.Next != nil {
				col.Next.Prev = col
			}

			firstCol = false
	}

	return det
}

