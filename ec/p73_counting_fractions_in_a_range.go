package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P73() *problem {
	return intInputNode(73, func(o command.Output, n int) {
		p := generator.Primes()
		var unique int
		for den := 4; den <= n; den++ {
			for num := den / 3; num*2 < den; num++ {
				if num*3 <= den {
					continue
				}
				if p.Coprimes(num, den) {
					continue
				}
				unique++
			}
		}
		o.Stdoutln(unique)
	}, []*execution{
		{
			args:     []string{"12000"},
			want:     "7295372",
			estimate: 8,
		},
		{
			args: []string{"8"},
			want: "3",
		},
	})
}
