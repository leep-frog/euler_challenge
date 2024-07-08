package eulerchallenge

import (
	"sort"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P41() *problem {
	return noInputNode(41, func(o command.Output) {
		// Can't be 9 or 8 digits because sum of 1 through 8|9 is divisible by 3
		possibilities := combinatorics.StringPermutations([]string{"1", "2", "3", "4", "5", "6", "7"})
		sort.Sort(sort.Reverse(sort.StringSlice(possibilities)))
		p := generator.Primes()
		for _, possStr := range possibilities {
			if p.Contains(parse.Atoi(possStr)) {
				o.Stdoutln(possStr)
				return
			}
		}
	}, &execution{
		want: "7652413",
	})
}
