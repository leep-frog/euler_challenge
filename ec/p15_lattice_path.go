package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

/*
Need to go left n times and right n times
So equivalent to the number of binary strings of length 2n with exactly n 1s and 0s
Which is
(2n choose n)
 = (2n)! / ((n!)*n!)
 = (2n)! / 2(n!)
 = (2n * (2n-1) * ... * (n + 1)) / (n * (n-1) * ... * 1)
 =
*/

func P15() *problem {
	return intInputNode(15, func(o command.Output, ni int) {
		n := maths.NewInt(int64(ni))

		var top, bottom, i = maths.NewInt(1), maths.NewInt(1), maths.NewInt(1)
		for ; i.LTE(n); i.PP() {
			top = top.Times(i.Plus(n))
			bottom = bottom.Times(i)
		}
		o.Stdoutln(top.Div(bottom))
	})
}
