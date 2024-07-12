package p8

import (
	"path/filepath"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/parse"
)

func P8() *ecmodels.Problem {
	return ecmodels.IntInputNode(8, func(o command.Output, n int) {
		s := parse.ReadFileInput(filepath.Join("..", "input", "p8.txt"))
		var is []int
		for i := 0; i < len(s); i++ {
			is = append(is, parse.Atoi(s[i:i+1]))
		}

		var max int

		for i := n; i < len(s); i++ {
			product := 1
			for j := i - n; j < i; j++ {
				product *= is[j]
			}
			if product > max {
				max = product
			}
		}

		o.Stdoutln(max)
	}, []*ecmodels.Execution{
		{
			Args: []string{"4"},
			Want: "5832",
		},
		{
			Args: []string{"13"},
			Want: "23514624000",
		},
	})
}
