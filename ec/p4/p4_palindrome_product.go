package p4

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func isPalindrome(i int) bool {
	str := fmt.Sprintf("%d", i)
	for j := 0; j < len(str); j++ {
		if str[j] != str[len(str)-j-1] {
			return false
		}
	}
	return true
}

func P4() *ecmodels.Problem {
	return ecmodels.IntInputNode(4, func(o command.Output, n int) {
		start := 1
		for i := 1; i < n; i++ {
			start *= 10
		}
		end := start * 10

		var biggestPalindrome int
		for i := start; i < end; i++ {
			for j := i + 1; j < end; j++ {
				p := i * j
				if p > biggestPalindrome && isPalindrome(p) {
					biggestPalindrome = p
				}
			}
		}
		o.Stdoutln(biggestPalindrome)
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "9009",
		},
		{
			Args: []string{"3"},
			Want: "906609",
		},
	})
}
