package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P14() *command.Node {
	return command.SerialNodes(
		command.Description("Find the longest Collatz series under N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			found := map[int]int{}
			n := d.Int(N)
			var max, maxI int
			for i := 2; i < n; i++ {
				count := 1
				for j := i; j != 1; {
					if j%2 == 0 {
						j /= 2
					} else {
						j = 3*j + 1
					}

					if v, ok := found[j]; ok {
						count += v
						break
					}
					count++
				}
				found[i] = count
				if count > max {
					max = count
					maxI = i
				}
			}
			o.Stdoutln(maxI)
		}))
}
