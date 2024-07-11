package p87

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P87() *ecmodels.Problem {
	return ecmodels.IntInputNode(87, func(o command.Output, n int) {
		p := generator.Primes()
		unique := map[int]bool{}
		for a := 0; ; a++ {
			av := p.Nth(a) * p.Nth(a)
			if av >= n {
				break
			}
			for b := 0; ; b++ {
				bv := p.Nth(b)*p.Nth(b)*p.Nth(b) + av
				if bv >= n {
					break
				}
				for c := 0; ; c++ {
					cv := p.Nth(c)*p.Nth(c)*p.Nth(c)*p.Nth(c) + bv
					if cv > n {
						break
					}
					unique[cv] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	}, []*ecmodels.Execution{
		{
			Args:     []string{"50000000"},
			Want:     "1097343",
			Estimate: 0.3,
		},
		{
			Args: []string{"50"},
			Want: "4",
		},
	})
}
