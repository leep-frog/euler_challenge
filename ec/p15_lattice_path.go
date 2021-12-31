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

func P15() *command.Node {
	return command.SerialNodes(
		command.Description("Find the number of unique lattice paths for an n x n grid"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := maths.NewInt(int64(d.Int(N)))

			var top, bottom, i = maths.NewInt(1), maths.NewInt(1), maths.NewInt(1)
			for ; i.LTE(n); i.PP() {
				top = top.Times(i.Plus(n))
				bottom = bottom.Times(i)
			}
			r, _ := top.Div(bottom)
			o.Stdoutln(r)
		}),
	)
}
