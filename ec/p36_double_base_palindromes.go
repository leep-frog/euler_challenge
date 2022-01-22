package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P36() *problem {
	return intInputNode(36, func(o command.Output, n int) {
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
	})
}
