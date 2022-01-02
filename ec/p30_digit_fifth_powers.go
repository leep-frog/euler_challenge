package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P30() *command.Node {
	return command.SerialNodes(
		command.Description("Find the sum of all numbers that can be written as the sum of the fifth power of their digits"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			v := 10
			for i := 0; i < n; i++ {
				v *= 10
			}

			var superSum int
			for i := 2; i < v; i++ {
				str := strconv.Itoa(i)
				var sum int
				for j := 0; j < len(str); j++ {
					digit := maths.Pow(parse.Atoi(str[j:j+1]), n)
					sum += digit
				}
				if sum == i {
					superSum += i
				}
			}
			o.Stdoutln(superSum)
		}),
	)
}
