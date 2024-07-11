package p35

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P35() *ecmodels.Problem {
	return ecmodels.IntInputNode(35, func(o command.Output, n int) {
		checked := map[string]bool{}
		unique := map[string]bool{}
		p := generator.Primes()
		for i := 0; p.Nth(i) < n; i++ {
			pn := p.Nth(i)
			spn := strconv.Itoa(pn)
			if checked[spn] {
				continue
			}
			var digits []string
			for j := 0; j < len(spn); j++ {
				digits = append(digits, spn[j:j+1])
			}

			allPrime := true
			rots := combinatorics.Rotations(digits)
			for _, rot := range rots {
				checked[rot] = true
				if !p.Contains(parse.Atoi(rot)) {
					allPrime = false
				}
			}
			if allPrime {
				for _, rot := range rots {
					unique[rot] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000000"},
			Want:     "55",
			Estimate: 0.25,
		},
		{
			Args: []string{"100"},
			Want: "13",
		},
	})
}
