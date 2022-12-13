package maths

import (
	"fmt"
	"math/big"

	"github.com/leep-frog/euler_challenge/linkedlist"
)

func BiggifyMatrix(matrix [][]float64) [][]*big.Rat {
	var m [][]*big.Rat
	for _, r := range matrix {
		var row []*big.Rat
		for _, v := range r {
			row = append(row, big.NewRat(0, 1).SetFloat64(v))
		}
		m = append(m, row)
	}
	return m
}

func BiggifyIntMatrix(matrix [][]int) [][]*big.Rat {
	var m [][]*big.Rat
	for _, r := range matrix {
		var row []*big.Rat
		for _, v := range r {
			row = append(row, big.NewRat(int64(v), 1))
		}
		m = append(m, row)
	}
	return m
}

func SmallifyMatrix(matrix [][]*big.Rat) [][]float64 {
	var m [][]float64
	for _, r := range matrix {
		var row []float64
		for _, v := range r {
			if v == nil {
				row = append(row, 0)
			} else {
				f, _ := v.Float64()
				row = append(row, f)
			}
		}
		m = append(m, row)
	}
	return m
}

func IdentityMatrix(n int) [][]*big.Rat {
	var m [][]*big.Rat
	for i := 0; i < n; i++ {
		r := make([]*big.Rat, n)
		r[i] = big.NewRat(1.0, 1.0)
		m = append(m, r)
	}
	return m
}

func MultiplyMatrices(this, that [][]*big.Rat) [][]*big.Rat {
	if len(this) == 0 || len(that) == 0 {
		panic("cannot multiply empty matrices")
	}

	if len(this[0]) != len(that) {
		panic("cannot multiply matrices with unmatched dimensions")
	}

	var result [][]*big.Rat
	for iThis := 0; iThis < len(this); iThis++ {
		var row []*big.Rat
		for iThat := 0; iThat < len(that[0]); iThat++ {
			sum := big.NewRat(0, 1)
			for j := 0; j < len(this[0]); j++ {
				sum.Add(sum, big.NewRat(0, 1).Mul(this[iThis][j], that[j][iThat]))
			}
			row = append(row, sum)
		}
		result = append(result, row)
	}
	return result
}

// Transpose transposes the provided matrix.
func Transpose(matrix [][]*big.Rat) [][]*big.Rat {
	var m [][]*big.Rat
	if len(matrix) == 0 {
		return m
	}
	for j := range Range(len(matrix[0])) {
		var col []*big.Rat
		for i := range Range(len(matrix)) {
			col = append(col, big.NewRat(0, 1).Set(matrix[i][j]))
		}
		m = append(m, col)
	}
	return m
}

// Inverse gets the inverse of the provided matrix.
func Inverse(matrix [][]*big.Rat) [][]*big.Rat {
	im := AdjugateMatrix(matrix)
	d := Determinant(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			im[i][j].Quo(im[i][j], d)
		}
	}
	return im
}

// AdjugateMatrix creates the adjugate matrix for the provided matrix.
// This is mainly used for calculating the inverse of a matrix.
func AdjugateMatrix(matrix [][]*big.Rat) [][]*big.Rat {
	matrix = Transpose(matrix)

	detCache := map[string]*big.Rat{}

	nRows := len(matrix)
	nCols := len(matrix[0])
	if nRows != nCols {
		panic("can only get adjugate of a square matrix")
	}

	if nRows == 1 {
		return [][]*big.Rat{{big.NewRat(1, 1)}}
	}

	var rs, cs []int
	for i := range Range(nRows) {
		rs = append(rs, i)
	}
	for i := range Range(nCols) {
		cs = append(cs, i)
	}

	rows, cols := linkedlist.NewList(rs), linkedlist.NewList(cs)

	var adj [][]*big.Rat
	firstRow := true
	for r := rows; r != nil; r = r.Next {
		rHead := rows
		if firstRow {
			rHead = rows.Next
		}
		var adjRow []*big.Rat
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

			d := determinant(matrix, nRows-1, nCols-1, rHead, cHead, detCache)
			if (r.Value+c.Value)%2 == 1 {
				d.Mul(d, big.NewRat(-1, 1))
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
}

// Determinant calculates the determinant of the provided matrix.
func Determinant(matrix [][]*big.Rat) *big.Rat {
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

	return determinant(matrix, nRows, nCols, linkedlist.NewList(rs), linkedlist.NewList(cs), map[string]*big.Rat{})
}

func determinant(matrix [][]*big.Rat, nRows, nCols int, rows, cols *linkedlist.Node[int], detCache map[string]*big.Rat) *big.Rat {
	code := fmt.Sprintf("%v:%v", rows, cols)
	if detCache != nil {
		if r, ok := detCache[code]; ok {
			return big.NewRat(0, 1).Set(r)
		}
	}

	if nRows == 1 {
		return matrix[rows.Value][cols.Value]
	}

	pos := true
	det := big.NewRat(0, 1)
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

		subDet := big.NewRat(0, 1).Mul(cell, determinant(matrix, nRows-1, nCols-1, rHead, cHead, detCache))
		if pos {
			det.Add(det, subDet)
		} else {
			det.Sub(det, subDet)
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

	detCache[code] = det
	return det
}

func CrossProduct(x1, y1, x2, y2 int) *big.Rat {
	return Determinant(BiggifyIntMatrix([][]int{
		{x1, y1},
		{x2, y2},
	}))
}

func CrossProductSign(x1, y1, x2, y2 int) int {
	return CrossProduct(x1, y1, x2, y2).Sign()
}
