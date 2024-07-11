package p36

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P36() *ecmodels.Problem {
	return ecmodels.IntInputNode(36, func(o command.Output, n int) {
		var palins []int
		for j, prod := 1, 1; prod < n; j++ {
			palins = append(palins, maths.Palindromes(j)...)
			prod *= 10
		}

		var sum int
		for _, palin := range palins {
			if maths.ToBinary(palin).Palindrome() {
				sum += palin
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000000"},
			Want: "872187",
		},
		{
			Args: []string{"10"},
			Want: "25",
		},
	})
}
