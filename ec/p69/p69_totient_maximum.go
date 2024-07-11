package p69

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P69() *ecmodels.Problem {
	return ecmodels.IntInputNode(69, func(o command.Output, n int) {
		p := generator.Primes().Iterator()
		prod := 1
		for ; prod < n; prod *= p.Next() {
		}
		o.Stdoutln(prod / p.Last())
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000000"},
			Want: "510510",
		},
		{
			Args: []string{"10"},
			Want: "6",
		},
	})
}
