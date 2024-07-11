package p204

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P204() *ecmodels.Problem {
	return ecmodels.IntInputNode(204, func(o command.Output, n int) {
		p := generator.Primes()

		upTo := 5
		if n == 9 {
			upTo = 100
		}

		var opts []int
		for i := 0; p.Nth(i) <= upTo; i++ {
			opts = append(opts, p.Nth(i))
		}
		o.Stdoutln(1 + calc204(opts, maths.Pow(10, n), 1))

	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "1105",
		},
		{
			Args:     []string{"9"},
			Want:     "2944730",
			Estimate: 1,
		},
	})
}

func calc204(opts []int, n, cur int) int {
	if cur > n {
		return 0
	}

	if len(opts) == 0 {
		return 1
	}

	var count int
	for prod := 1; prod*cur < n; prod *= opts[0] {
		count += calc204(opts[1:], n, cur*prod)
	}
	return count
}
