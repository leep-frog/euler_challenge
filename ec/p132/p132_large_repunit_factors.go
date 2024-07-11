package p132

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p129"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P132() *ecmodels.Problem {
	return ecmodels.IntInputNode(132, func(o command.Output, n int) {
		rLen := maths.Pow(10, n)
		g := generator.Primes()
		var sum, count int
		prod := 1
		for i, pi := 0, g.Nth(0); count < 40 && len(strconv.Itoa(prod)) < rLen; i, pi = i+1, g.Nth(i+1) {
			if !p129.Repunitable(pi) {
				continue
			}
			k := p129.RepunitSmallest(pi)
			if rLen%k == 0 {
				sum += pi
				count++
				prod *= pi
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"9"},
			Want:     "843296",
			Estimate: 5,
		},
		{
			Args: []string{"1"},
			Want: "9414",
		},
	})
}
