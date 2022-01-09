package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P6() *command.Node {
	return command.SerialNodes(
		command.Description("Find the difference between the sum of squares and the square of sums up to N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			var squareSums, sumSquares int
			for i := 1; i <= d.Int(N); i++ {
				sumSquares += i * i
				squareSums += i
			}
			squareSums *= squareSums
			o.Stdoutln(squareSums - sumSquares)
		}),
	)
}
