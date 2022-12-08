package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P47() *problem {
	return intInputNode(47, func(o command.Output, n int) {
		p := generator.Primes()
		var row int
		for i := 1; ; i++ {
			if len(p.PrimeFactors(i)) >= n {
				row++
				if row == n {
					o.Stdoutln(i - (n - 1))
					return
				}
			} else {
				row = 0
			}
		}
	}, []*execution{
		{
			args:     []string{"4"},
			want:     "134043",
			estimate: 0.35,
		},
		{
			args: []string{"3"},
			want: "644",
		},
		{
			args: []string{"2"},
			want: "14",
		},
	})
}
