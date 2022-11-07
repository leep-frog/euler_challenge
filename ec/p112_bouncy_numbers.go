package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P112() *problem {
	return intInputNode(112, func(o command.Output, n int) {
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
	}, []*execution{
		{
			args:     []string{"99"},
			want:     "1587000",
			estimate: 0.25,
		},
		{
			args: []string{"90"},
			want: "21780",
		},
		{
			args: []string{"50"},
			want: "538",
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
