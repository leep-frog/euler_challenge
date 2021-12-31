package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P21() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of amicable numbers up to n"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			sumMap := map[int]int{}
			for i := 1; i <= n; i++ {
				count := 0
				for j := 1; j*j <= i; j++ {
					if i%j == 0 {
						if j*j == i || j == 1 {
							count += j
						} else {
							count += j + i/j
						}
					}
				}
				sumMap[i] = count
			}

			var sum int
			for k, v := range sumMap {
				if sumMap[v] == k && v != k {
					sum += k + v
				}
			}

			// Divide by 2 since each pair is counted twice ((k, v) and (v, k)).
			o.Stdoutln(sum / 2)
		}),
	)
}
