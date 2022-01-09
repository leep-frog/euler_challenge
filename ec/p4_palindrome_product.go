package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
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

func P4() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest palindrome product of two, N-digit integers"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			start := 1
			for i := 1; i < d.Int(N); i++ {
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
		}),
	)
}
