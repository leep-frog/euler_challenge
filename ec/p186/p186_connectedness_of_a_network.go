package p186

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/unionfind"
)

const (
	people = 1_000_000
)

func P186() *ecmodels.Problem {
	return ecmodels.NoInputNode(186, func(o command.Output) {

		uf := unionfind.New[int]()

		lfg := &laggedFibGenerator{}

		var calls int
		for n := 1; uf.LargestSetSize() < 990_000; n++ {
			a, b := lfg.at(2*n-1), lfg.at(2*n)
			if a != b {
				calls++
				uf.Merge(a, b)
			}
		}

		o.Stdoutln(calls)
	}, &ecmodels.Execution{
		Want:     "2325629",
		Estimate: 1.5,
	})
}

type laggedFibGenerator struct {
	s []int
}

func (lfg *laggedFibGenerator) at(idx int) int {
	for len(lfg.s) <= idx {
		k := len(lfg.s)

		if k == 0 {
			lfg.s = append(lfg.s, 0)
		} else if k <= 55 {
			lfg.s = append(lfg.s, (100_003-200_003*k+300_007*k*k*k)%people)
		} else {
			lfg.s = append(lfg.s, (lfg.s[k-24]+lfg.s[k-55])%people)
		}
	}
	return lfg.s[idx]
}
