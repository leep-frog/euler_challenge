package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"golang.org/x/exp/maps"
)

func P154() *problem {
	return intsInputNode(154, 2, func(o command.Output, ns []int) {
		// Pascal Pyramid layer
		n := ns[0]
		// Divisble by
		div := ns[1]

		// Create the
		divFactors := generator.PrimeFactors(div, generator.Primes())
		factors := maps.Keys(divFactors)
		var minNeeded []int
		for _, f := range factors {
			minNeeded = append(minNeeded, divFactors[f])
		}
		fmt.Println("FACTORS", factors)
		fmt.Println("MINNEED", minNeeded)

		fc := NewFC(n, factors)

		cnt := 0
		for i := 0; i <= n; i++ {
			if i%1_000 == 0 {
				fmt.Println(i)
			}
			coef := fc.Choose(n, i)
			for j := 0; j <= i; j++ {
				v := fc.Choose(i, j)
				good := true
				for i, mn := range minNeeded {
					if coef[i]+v[i] < mn {
						good = false
						break
					}
				}
				if good {
					cnt++
				}
			}
		}
		fmt.Println(cnt)

		o.Stdoutln(cnt)
	}, []*execution{
		{
			args: []string{"200000", "1000000000000"},
			want: "479742450 YOU",
		},
	})
}

type FactorialChecker struct {
	factors []int

	divs [][]int
}

// Returns the number of each factor in the numerator
func (fc *FactorialChecker) Choose(n, k int) []int {
	// n! / (k!(n-k)!)

	var rs []int
	for i := range fc.factors {
		// Number of the factors in n!
		nf := fc.divs[n][i]
		// Number of the factors in k!
		kf := fc.divs[k][i]
		// Number of the factors in k!
		knf := fc.divs[n-k][i]

		// Number of factors in the numerator:
		rs = append(rs, nf-kf-knf)
	}
	return rs
}

// Note: only works if factors are primes.
func NewFC(k int, factors []int) *FactorialChecker {
	// currentF is an array of [factor, factor^x] where factor^x is
	// bigger than the current number we are on in the iteration below.
	var currentF [][]int
	for _, f := range factors {
		currentF = append(currentF, []int{f, f})
	}

	lf := len(factors)
	divs := [][]int{make([]int, lf, lf)}
	for i := 1; i <= k; i++ {
		var newRow []int
		for ci, cf := range currentF {
			v := divs[i-1][ci]
			for s := i; s%cf[1] == 0; s, v = s/cf[1], v+1 {
			}
			newRow = append(newRow, v)
		}
		divs = append(divs, newRow)
	}
	return &FactorialChecker{factors, divs}
}

type GreedyPascalTriangleRow struct {
	rows [][]int
}

func PascalIdx(row, col int) int {
	return ShapeNumber(col+1, row+1)
}

func ShapeNumber(s, n int) int {
	// https://en.wikipedia.org/wiki/Polygonal_number#Formula
	return ((s-2)*n*n - (s-4)*n) / 2
}

// 0 indexed
func (gpt *GreedyPascalTriangleRow) HalfRow(r int) []int {

	for ln := len(gpt.rows); ln <= r; ln++ {
		if ln <= 1 {
			gpt.rows = append(gpt.rows, []int{})
			continue
		}
		if ln == 2 {
			gpt.rows = append(gpt.rows, []int{2})
			continue
		}

		if ln%2000 == 0 {
			fmt.Println(ln)
		}

		row := gpt.rows[ln-1]
		newRow := []int{ln}
		if ln%2 == 0 {
			for j := 1; j < (ln-1)/2; j++ {
				newRow = append(newRow, row[j-1]+row[j])
			}
			newRow = append(newRow, 2*row[len(row)-1])
		} else {
			for j := 1; j < ln/2; j++ {
				newRow = append(newRow, row[j-1]+row[j])
			}
		}

		gpt.rows = append(gpt.rows, newRow)
	}
	return gpt.rows[r]
}

type PascalTriangle struct {
	Rows [][]int
}

func (pt *PascalTriangle) Row(i int) []int {
	for len(pt.Rows) <= i {
		if len(pt.Rows)%2000 == 0 {
			fmt.Println(len(pt.Rows))
		}
		if len(pt.Rows) == 0 {
			pt.Rows = append(pt.Rows, []int{1})
			continue
		}

		prevRow := pt.Rows[len(pt.Rows)-1]
		newRow := []int{1}
		for i, v := range prevRow {
			if i == len(prevRow)-1 {
				newRow = append(newRow, 1)
			} else {
				newRow = append(newRow, v+prevRow[i+1])
			}
		}
		pt.Rows = append(pt.Rows, newRow)
	}
	return pt.Rows[i]
}

type PascalPyramid struct{}

func (py *PascalPyramid) Layer(i int) [][]int {
	return nil
}
