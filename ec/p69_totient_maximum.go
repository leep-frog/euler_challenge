package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P69() *problem {
	return intInputNode(69, func(o command.Output, n int) {
		p := generator.Primes().Iterator()
		prod := 1
		for ; prod < n; prod *= p.Next() {
		}
		o.Stdoutln(prod / p.Last())
	}, []*execution{
		{
			args: []string{"1000000"},
			want: "510510",
		},
		{
			args: []string{"10"},
			want: "6",
		},
	})
}
