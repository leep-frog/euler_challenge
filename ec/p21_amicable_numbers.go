package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
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
				for _, div := range maths.Divisors(i) {
					if div != i {
						count += div
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
