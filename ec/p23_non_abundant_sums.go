package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P23() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			var abundant []int
			for i := 1; i <= n; i++ {
				divs := maths.Divisors(i)
				var sum int
				for _, d := range divs {
					if d != i {
						sum += d
					}
				}
				if sum > i {
					abundant = append(abundant, i)
				}
			}

			isSum := map[int]bool{}
			for i, ai := range abundant {
				for j := i; j < len(abundant); j++ {
					aj := abundant[j]
					isSum[ai+aj] = true
				}
			}

			var sum int
			for i := 1; i <= n; i++ {
				if !isSum[i] {
					sum += i
				}
			}
			o.Stdoutln(sum)
		}),
	)
}
