package p345

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func P345() *ecmodels.Problem {
	return ecmodels.FileInputNode(345, func(lines []string, o command.Output) {
		grid := functional.Map(parse.SplitWhitespace(lines), parse.AtoiArray)

		rcs := map[int]bool{}
		for i := range grid {
			rcs[i] = true
		}

		o.Stdoutln(dp(grid, rcs, map[string]int{}))
	}, []*ecmodels.Execution{
		{
			Args: []string{"-x"},
			Want: "3315",
		},
		{
			Want: "13938",
		},
	})
}

func dp(grid [][]int, remainingCols map[int]bool, cache map[string]int) int {
	rowIdx := len(grid) - len(remainingCols)
	if rowIdx == len(grid) {
		return 0
	}

	colKeys := maps.Keys(remainingCols)
	slices.Sort(colKeys)
	code := fmt.Sprintf("%v", colKeys)
	if v, ok := cache[code]; ok {
		return v
	}

	best := maths.Largest[int, int]()
	for _, rc := range colKeys {
		delete(remainingCols, rc)
		best.Check(grid[rowIdx][rc] + dp(grid, remainingCols, cache))
		remainingCols[rc] = true
	}

	cache[code] = best.Best()
	return best.Best()
}
