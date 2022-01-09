package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P8() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest in series product for N integers"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			s := parse.ReadFileInput("p8.txt")
			var is []int
			for i := 0; i < len(s); i++ {
				is = append(is, parse.Atoi(s[i:i+1]))
			}

			var max int

			for i := d.Int(N); i < len(s); i++ {
				product := 1
				for j := i - d.Int(N); j < i; j++ {
					product *= is[j]
				}
				if product > max {
					max = product
				}
			}

			o.Stdoutln(max)
		}),
	)
}
