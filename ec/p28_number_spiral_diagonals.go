package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P28() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of the diagonals on an n by n spiral"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			start := 3
			sum := 1
			for i := 0; i < (n-1)/2; i++ {
				offset := (i + 1) * 2
				sum += 4*start + 6*offset
				start += offset*4 + 2
			}
			o.Stdoutln(sum)
		}),
	)
}
