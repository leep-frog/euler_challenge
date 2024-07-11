package p130

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p129"
	"github.com/leep-frog/euler_challenge/generator"
)

func P130() *ecmodels.Problem {
	return ecmodels.IntInputNode(130, func(o command.Output, n int) {
		diffs := []int{2, 4, 2, 2}
		g := generator.Primes()

		var count, sum int
		for i, j := 3, 1; ; i, j = i+diffs[j], (j+1)%len(diffs) {
			if g.Contains(i) {
				continue
			}
			k := p129.RepunitSmallest(i)

			if (i-1)%k == 0 {
				count++
				sum += i
				if count >= n {
					o.Stdoutln(sum)
					return
				}
			}
		}
	}, []*ecmodels.Execution{
		{
			Args: []string{"25"},
			Want: "149253",
		},
		{
			Args: []string{"5"},
			Want: "1985",
		},
	})
}
