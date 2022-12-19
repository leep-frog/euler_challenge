package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P7() *problem {
	return intInputNode(7, func(o command.Output, n int) {
		p := generator.Primes()
		// p := generator.SievedPrimes()
		// p := generator.BasicPrimes()
		o.Stdoutln(p.Nth(n - 1))
	}, []*execution{
		{
			args: []string{"6"},
			want: "13 13",
		},
		{
			args: []string{"9"},
			want: "23 23",
		},
		{
			args: []string{"10001"},
			want: "104743 104743",
		},
		{
			args:     []string{"10000001"},
			want:     "179424691",
			estimate: 36,
		},
	})
}
