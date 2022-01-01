package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P20() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of the digits of n!"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			o.Stdoutln(maths.Facotiral(n).DigitSum())
		}),
	)
}
