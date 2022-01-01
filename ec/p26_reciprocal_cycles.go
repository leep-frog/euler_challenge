package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func cycleLen(num int) int {
	rem := 1
	remMap := map[int]int{}

	for pos := 0; ; pos++ {
		rem = (rem % num) * 10
		if rem == 0 {
			return 0
		}
		if v, ok := remMap[rem]; ok {
			return pos - v
		}
		remMap[rem] = pos
	}
}

func P26() *command.Node {
	return command.SerialNodes(
		command.Description("Find the unit fraction up to n that has the longest recurring cycle"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			var max, maxI int
			for i := 1; i < n; i++ {
				if v := cycleLen(i); v > max {
					max = v
					maxI = i
				}
			}
			o.Stdoutln(maxI)
		}),
	)
}
