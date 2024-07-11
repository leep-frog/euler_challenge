package p10

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P10() *ecmodels.Problem {
	return ecmodels.IntInputNode(10, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		for i, pn := 0, p.Nth(0); pn < n; i, pn = i+1, p.Nth(i+1) {
			sum += pn
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"10"},
			Want: "17",
		},
		{
			Args: []string{"2000000"},
			Want: "142913828922",
		},
	})
}
