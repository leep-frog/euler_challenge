package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P64() *problem {
	return intInputNode(64, func(o command.Output, n int) {
		var count int
		for k := 2; k <= n; k++ {
			_, period := maths.SquareRootPeriod(k)
			if len(period)%2 == 1 {
				count++
			}
		}
		o.Stdoutln(count)
	})
}
