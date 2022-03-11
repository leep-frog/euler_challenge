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
			if cell == 0 {
				continue
			}
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
