package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func P52() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=52"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			start := "1"
			end := "1"
			for {
				start += "0"
				end += "6"
				sn := parse.Atoi(start)
				en := parse.Atoi(end)
				for i := sn + 1; i <= en; i++ {
					allSame := true
					for j := 2; j <= n; j++ {
						if !sameDigits(i, i*j) {
							allSame = false
							break
						}
					}
					if allSame {
						o.Stdoutln(i)
						return
					}
				}
			}
		}),
	)
}
