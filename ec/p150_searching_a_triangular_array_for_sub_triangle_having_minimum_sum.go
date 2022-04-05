package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func smallTriangle() []int {
	return []int{
		15,
		-14, -7,
		20, -13, -5,
		-3, 8, 23, -26,
		1, -4, -5, -18, 5,
		-16, 31, 2, 9, 28, 3,
	}
}

func bigTriangle() []int {
	var t int
	var seq []int
	big2, small2 := maths.Pow(2, 20), maths.Pow(2, 19)
	for k := 1; k <= 500500; k++ {
		t = (615949*t + 797807) % big2
		seq = append(seq, t-small2)
	}
	return seq
}

func P150() *problem {
	return intInputNode(150, func(o command.Output, n int) {
		// Construct triangle from sequence
		tri := smallTriangle()
		if n > 1 {
			tri = bigTriangle()
		}

		var topDown [][]int
		for rowLen, idx := 1, 0; idx+rowLen <= len(tri); rowLen, idx = rowLen+1, idx+rowLen {
			var row []int
			for i := idx; i < idx+rowLen; i++ {
				row = append(row, tri[i])
			}
			topDown = append(topDown, row)
		}

		// Cumulative rows of triangle
		var cumTopDown [][]int
		for i := 0; i < len(topDown); i++ {
			cumTopDown = append(cumTopDown, append([]int{0}, maths.Cumulative(topDown[i])...))
		}

		// iterate over top = (row, col) and edgeLen = k
		// top can be one of 500_500 positions, edgeLen can be up to 1000
		// So 500_500_000 iterations
		best := maths.Smallest[[]int, int]()
		for topRow := 0; topRow < len(topDown); topRow++ {
			for topCol := 0; topCol < len(topDown[topRow]); topCol++ {
				startSum := topDown[topRow][topCol]
				best.IndexCheck([]int{topRow, topCol, 0}, startSum)
				for edgeLen := 1; edgeLen+topRow < len(topDown); edgeLen++ {
					startSum += cumTopDown[topRow+edgeLen][topCol+edgeLen+1] - cumTopDown[topRow+edgeLen][topCol]
					best.IndexCheck([]int{topRow, topCol, edgeLen}, startSum)
				}
			}
		}
		o.Stdoutln(best.Best())
	})
}
