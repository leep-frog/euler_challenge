package maths

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leep-frog/euler_challenge/linkedlist"
)

type Matrix[T Mathable] struct {
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
}

func Determinant[T Mathable](matrix [][]T) T {
	nRows := len(matrix)
	nCols := len(matrix[0])

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
	if nRows != nCols {
		panic("can only get determinant of a square matrix")
	}
	if nRows == 0 {
		panic("can't get determinant of an empty matrix")
	}
	
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

/*func (m *Matrix[T]) Determinant() T {
	if len(m.rows) != len(m.cols) {
		panic("can only get determinant of a square matrix")
	}
	if len(m.rows) == 0 {
		panic("can't get determinant of an empty matrix")
	}

	if len(m.rows) == 0 {
		return m.Get(0, 0)
	}

	code := m.subsetCode()
	if v, ok := m.determinantCache[code]; ok {
		return v
	}

	for _, r := range m.rows {
		for _, c := range m.cols {

		}
	}
}*/

// Get the code for a matrix subset
func (m *Matrix[T]) subsetCode() string {
	var r, c []string
	for _, ri := range m.rows {
		r = append(r, strconv.Itoa(ri))
	}
	for _, ci := range m.cols {
		c = append(c, strconv.Itoa(ci))
	}
	return fmt.Sprintf("%s:%s", strings.Join(r, "_"), strings.Join(c, "_"))
}

/*func (m *Matrix[T]) Get(r, c int) T {
	return m.values[m.rows[r]][m.cols[c]]
}*/
