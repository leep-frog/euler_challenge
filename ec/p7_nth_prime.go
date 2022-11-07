package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P7() *problem {
	return intInputNode(7, func(o command.Output, n int) {
		o.Stdoutln(generator.Primes().Nth(n - 1))
	}, []*execution{
		{
			args: []string{"6"},
			want: "13",
		},
		{
			args: []string{"10001"},
			want: "104743",
		},
	})
}
