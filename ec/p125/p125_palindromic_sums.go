package p125

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P125() *ecmodels.Problem {
	return ecmodels.IntInputNode(125, func(o command.Output, n int) {
		set := map[int]bool{}
		for i := 1; i*i <= n; i++ {
			sum := i * i
			for j := i + 1; sum+j*j <= n; j++ {
				sum += j * j
				if maths.Palindrome(sum) {
					// Some palindromes can be written in different ways
					set[sum] = true
				}
			}
		}
		var total int
		for k := range set {
			total += k
		}
		o.Stdoutln(total)
	}, []*ecmodels.Execution{
		{
			Args: []string{"100000000"},
			Want: "2906969179",
		},
		{
			Args: []string{"1000"},
			Want: "4164",
		},
	})
}
