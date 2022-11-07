package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P10() *problem {
	return intInputNode(10, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		for i, pn := 0, p.Nth(0); pn < n; i, pn = i+1, p.Nth(i+1) {
			sum += pn
		}
		o.Stdoutln(sum)
	}, []*execution{
		{
			args: []string{"10"},
			want: "17",
		},
		{
			args: []string{"2000000"},
			want: "142913828922",
		},
	})
}
