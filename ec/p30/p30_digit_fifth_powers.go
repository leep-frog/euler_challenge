package p30

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P30() *ecmodels.Problem {
	return ecmodels.IntInputNode(30, func(o command.Output, n int) {
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
	}, []*ecmodels.Execution{
		{
			Args: []string{"5"},
			Want: "443839",
		},
		{
			Args: []string{"4"},
			Want: "19316",
		},
	})
}
