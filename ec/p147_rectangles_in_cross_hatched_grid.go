package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P147() *problem {
	return intInputNode(147, func(o command.Output, n int) {
		W, H := 47, 43
		if n > 1 {
			W, H = 3, 2
		}

		diagCounts := map[int]map[int]int{
			1: {},
		}

		for w := 1; w <= W; w++ {
			diagCounts[w] = map[int]int{1: w - 1}
			diagCounts[1][w] = w - 1
		}

		diagCounts[2][2] = 9

		for w := 3; w <= W; w++ {
			diagCounts[w][2] = diagCounts[w-1][2] + 10
			diagCounts[2][w] = diagCounts[w][2]
		}

		for w := 3; w <= W; w++ {
			for h := 3; h <= w && h <= H; h++ {
				// The rectangle that is one shorter can be drawn in two positions (all the way right or left)
				// Removing the rectangle that is two shorter removes the doubly counted rectangles
				// that are counted by the two w-1 diagCounts.
				// Then, we just need to count the number of diagonals that include a square on the top and bottom.
				diagCounts[w][h] = 2*diagCounts[w-1][h] - diagCounts[w-2][h] + newDiagCount(w, h)
				diagCounts[h][w] = diagCounts[w][h]
			}
		}

		var sum int
		for w := 1; w <= W; w++ {
			for h := 1; h <= H && h <= w; h++ {
				orthoCount := (w * (w + 1) * h * (h + 1)) / 4
				total := orthoCount + diagCounts[w][h]

				if w == h || w > H {
					sum += total
				} else {
					sum += 2 * total
				}
			}
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"1"},
			want: "846910284",
		},
		{
			args: []string{"2"},
			want: "72",
		},
	})
}

func newDiagCount(nRows, nCols int) int {
	if nRows == 1 {
		return nCols - 1
	}
	if nCols == 1 {
		return nRows - 1
	}
	var count int
	for i := 0; i < nCols; i++ {
		count += topToBottom(i, nRows, nCols)
	}
	return count
}

func topToBottom(col, nRows, nCols int) int {
	var count int
	// The very top row of new diamonds
	if col < nCols-1 {
		lmBottom, lmTop := leftMost(col, nRows)
		rmBottom, rmTop := rightMost(col, nRows, nCols)

		//fmt.Println("un", lmBottom, lmTop, rmBottom, rmTop)
		if lmBottom <= rmBottom {
			count += rmBottom - lmBottom + 1
		}
		if lmTop <= rmTop {
			count += rmTop - lmTop + 1
		}
	}

	// The second row of new diamonds
	lmBottom, lmTop := leftMost(col, nRows)
	rmBottom, rmTop := rightMost(col-1, nRows, nCols)
	//fmt.Println("du", lmBottom, lmTop, rmBottom, rmTop)

	if lmBottom <= rmBottom {
		count += rmBottom - lmBottom + 1
	}
	if lmTop <= rmTop {
		count += rmTop - lmTop + 1
	}
	return count
}

func leftMost(col, nRows int) (int, int) {
	ret := col - (nRows - 1)
	if ret < 0 {
		return -(ret + 1), -(ret + 1)
	}
	return ret, ret + 1
}

func rightMost(col, nRows, nCols int) (int, int) {
	ret := col + nRows - 1
	// If bounce
	if ret >= nCols-1 {
		bottom := 2*nCols - 3 - ret
		return bottom, bottom + 1
	}
	return ret, ret
}
