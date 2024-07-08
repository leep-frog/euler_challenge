package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P179() *problem {
	return intInputNode(179, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		prev := 1
		for i := 2; i < 10_000_000; i++ {
			fn := p.FactorCount(i)
			if fn == prev {
				sum++
			}
			prev = fn
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args:     []string{"1"},
			want:     "986262",
			estimate: 7,
		},
	})
}
