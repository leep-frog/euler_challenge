package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
)

func P71() *problem {
	return intInputNode(71, func(o command.Output, n int) {
		best := fraction.New(1, 4)
		for den := 1; den < n; den++ {
			if den%7 == 0 {
				continue
			}
			newF := fraction.New(3*den/7, den)
			if maths.LTE(best, newF) {
				best = newF
			}
		}
		o.Stdoutln(best)
	})
}
