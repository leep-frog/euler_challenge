package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
)

func P106() *problem {
	return intInputNode(106, func(o command.Output, n int) {
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
	}, []*execution{
		{
			args:     []string{"12"},
			want:     "21384",
			estimate: 0.5,
		},
		{
			args: []string{"7"},
			want: "70",
		},
		{
			args: []string{"4"},
			want: "1",
		},
	})
}
