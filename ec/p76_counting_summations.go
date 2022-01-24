package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P76() *problem {
	return intInputNode(76, func(o command.Output, n int) {
		// Subtract one because the search includes the single digit summation ([]int{n}).
		o.Stdoutln(dfs76(n, 1) - 1)
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
