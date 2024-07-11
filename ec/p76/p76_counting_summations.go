package p76

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P76() *ecmodels.Problem {
	return ecmodels.IntInputNode(76, func(o command.Output, n int) {
		// Subtract one because the search includes the single digit summation ([]int{n}).
		o.Stdoutln(dfs76(n, 1) - 1)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"100"},
			Want:     "190569291",
			Estimate: 3,
		},
		{
			Args: []string{"5"},
			Want: "6",
		},
	})
}

func dfs76(remaining, value int) int {
	if remaining == 0 {
		return 1
	}
	var count int
	for i := value; i <= remaining; i++ {
		count += dfs76(remaining-i, i)
	}
	return count
}
