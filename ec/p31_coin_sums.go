package eulerchallenge

import (
	"github.com/leep-frog/command/command"
)

func P31() *problem {
	return noInputNode(31, func(o command.Output) {
		coins := []int{
			1, 2, 5, 10, 20, 50, 100, 200,
		}
		o.Stdoutln(try(coins, 0, 0))
	}, &execution{
		want: "73682",
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
