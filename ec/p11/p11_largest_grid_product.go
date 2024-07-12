package p11

import (
	"path/filepath"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/parse"
)

func P11() *ecmodels.Problem {
	return ecmodels.IntInputNode(11, func(o command.Output, n int) {
		ss := parse.ReadFileLines(filepath.Join("..", "input", "p11.txt"))
		var is [][]int
		for _, rowStr := range ss {
			var row []int
			cells := strings.Split(rowStr, " ")
			for _, cell := range cells {
				row = append(row, parse.Atoi(cell))
			}
			is = append(is, row)
		}

		max := 0
		// Check horizontal products
		for i := 0; i < len(is); i++ {
			for j := 0; j < len(is[i])-n; j++ {
				product := 1
				for k := j; k < j+n; k++ {
					product *= is[i][k]
				}
				if product > max {
					max = product
				}
			}
		}

		// Check vertical products
		for i := 0; i < len(is[0]); i++ {
			for j := 0; j < len(is)-n; j++ {
				product := 1
				for k := j; k < j+n; k++ {
					product *= is[k][i]
				}
				if product > max {
					max = product
				}
			}
		}

		// Check down right diagonal products
		for i := 0; i < len(is)-n; i++ {
			for j := 0; j < len(is[i])-n; j++ {
				product := 1
				for k := 0; k < n; k++ {
					product *= is[i+k][j+k]
				}
				if product > max {
					max = product
				}
			}
		}

		// Check down left diagonal products
		for i := n; i < len(is); i++ {
			for j := 0; j < len(is[i])-n; j++ {
				product := 1
				for k := 0; k < n; k++ {
					product *= is[i-k][j+k]
				}
				if product > max {
					max = product
				}
			}
		}

		o.Stdoutln(max)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "70600674",
		},
	})
}
