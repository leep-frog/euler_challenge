package p207

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
)

func P207() *ecmodels.Problem {
	return ecmodels.IntInputNode(207, func(o command.Output, n int) {
		// 4^t = 2^t + k
		// 4^t - 2^t = k
		// 2^t * 2^t - 2^t = k
		// 2^t * (2^t - 1) = k
		// Iterate over x and plugin values to x * (x - 1) to get numbers
		// Perfect solutions are when x = 2^t and where t is an integer
		// (so when x is a power of 2)

		threshold := fraction.New(1, 12345)

		nextPerfect := 2

		var perfectCount, count int
		for x := 2; ; x++ {
			count++
			if x == nextPerfect {
				perfectCount++
				nextPerfect *= 2
			}

			fmt.Println(x*(x-1), fmt.Sprintf("%d/%d", perfectCount, count))

			ratio := fraction.New(perfectCount, count)
			if ratio.LT(threshold) {
				fmt.Println("DONZO", x*(x-1))
				break
			}
		}
		// o.Stdoutln(n)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}
