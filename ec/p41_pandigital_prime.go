package eulerchallenge

import (
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P41() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=41"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			// Can't be 9 or 8 digits because sum of 1 through 8|9 is divisible by 3
			possibilities := maths.Permutations([]string{"1", "2", "3", "4", "5", "6", "7"})
			sort.Sort(sort.Reverse(sort.StringSlice(possibilities)))
			p := generator.Primes()
			for _, possStr := range possibilities {
				if generator.IsPrime(parse.Atoi(possStr), p) {
					o.Stdoutln(possStr)
					return
				}
			}
		}),
	)
}
