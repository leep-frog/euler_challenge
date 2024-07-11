package p112

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P112() *ecmodels.Problem {
	return ecmodels.IntInputNode(112, func(o command.Output, n int) {
		var bouncy int
		for i := 1; ; i++ {
			if bouncyNumber(i) {
				bouncy++
			}
			// n/100 = bouncy/i
			// n*i = 100*bouncy
			if n*i == 100*bouncy {
				o.Stdoutln(i)
				return
			}
		}
	}, []*ecmodels.Execution{
		{
			Args:     []string{"99"},
			Want:     "1587000",
			Estimate: 0.25,
		},
		{
			Args: []string{"90"},
			Want: "21780",
		},
		{
			Args: []string{"50"},
			Want: "538",
		},
	})
}

func bouncyNumber(n int) bool {
	increasing, decreasing := true, true
	digits := maths.Digits(n)
	for idx, d := range digits {
		if idx > 0 && d < digits[idx-1] {
			increasing = false
		}
		if idx > 0 && d > digits[idx-1] {
			decreasing = false
		}
	}
	return !increasing && !decreasing
}
