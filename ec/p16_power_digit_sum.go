package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P16() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of the digits of 2^n"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := maths.NewInt(int64(d.Int(N)))

			two := maths.NewInt(2)
			pow := maths.NewInt(1)
			for n.GT(maths.Zero()) {
				pow = pow.Times(two)
				n.MM()
			}

			o.Stdoutln(pow.DigitSum())
		}),
	)
}
