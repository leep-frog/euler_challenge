package p106

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

func P106() *ecmodels.Problem {
	return ecmodels.IntInputNode(106, func(o command.Output, n int) {
		options := []string{
			"A", // set A
			"B", // set B
			"N", // neither
		}
		var count int
		for _, perm := range combinatorics.GenerateCombos(&combinatorics.Combinatorics[string]{
			Parts:            options,
			MinLength:        n,
			MaxLength:        n,
			AllowReplacement: true,
			OrderMatters:     true,
		}) {
			var moreAs, moreBs bool
			var aCount, bCount int

			for i := len(perm) - 1; i >= 0; i-- {
				part := perm[i]
				if part == "A" {
					aCount++
				} else if part == "B" {
					bCount++
				}
				moreAs = moreAs || aCount > bCount
				moreBs = moreBs || bCount > aCount
			}

			if moreAs && moreBs && aCount == bCount {
				count++
			}
		}
		// Divide by 2 due to symmetry
		o.Stdoutln(count / 2)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"12"},
			Want:     "21384",
			Estimate: 0.5,
		},
		{
			Args: []string{"7"},
			Want: "70",
		},
		{
			Args: []string{"4"},
			Want: "1",
		},
	})
}
