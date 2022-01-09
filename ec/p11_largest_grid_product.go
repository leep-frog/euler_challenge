package eulerchallenge

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P11() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest N-length product in the grid"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			ss := parse.ReadFileLines("p11.txt")
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
		}),
	)
}
