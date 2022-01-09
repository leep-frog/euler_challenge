package eulerchallenge

import (
	"github.com/leep-frog/command"
)

const (
	N = "N"
)

func P1() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all numbers less than N that are divisble by 3 or 5"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			var sum int
			for i := 1; i < d.Int(N); i++ {
				if i%5 == 0 || i%3 == 0 {
					sum += i
				}
			}
			o.Stdoutln(sum)
		}),
	)
}
