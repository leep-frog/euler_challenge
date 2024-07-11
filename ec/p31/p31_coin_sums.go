package p31

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P31() *ecmodels.Problem {
	return ecmodels.NoInputNode(31, func(o command.Output) {
		coins := []int{
			1, 2, 5, 10, 20, 50, 100, 200,
		}
		o.Stdoutln(try(coins, 0, 0))
	}, &ecmodels.Execution{
		Want: "73682",
	})
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
