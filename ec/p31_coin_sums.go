package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P31() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=31"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			coins := []int{
				1, 2, 5, 10, 20, 50, 100, 200,
			}
			o.Stdoutln(try(coins, 0, 0))
		}),
	)
}

func try(coins []int, index, curSum int) int {
	if curSum == 200 {
		return 1
	} else if curSum > 200 || index >= len(coins) {
		return 0
	}

	v := coins[index]
	var sum int
	for i := 0; curSum+v*i <= 200; i++ {
		sum += try(coins, index+1, curSum+v*i)
	}
	return sum
}
