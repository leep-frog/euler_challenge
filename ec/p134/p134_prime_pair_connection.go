package p134

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P134() *ecmodels.Problem {
	return ecmodels.NoInputNode(134, func(o command.Output) {
		g := generator.Primes()
		var sum int
		for i := 2; g.Nth(i) < 1000000; i++ {
			p1, p2 := g.Nth(i), g.Nth(i+1)

			// p1 + offset*x mod p2 == 0
			// offset*x mod p2 == (p2-p1)
			// where offset = 10^k
			offset := maths.Pow(10, len(strconv.Itoa(p1)))
			x := maths.SolveMod(offset, p2, p2-p1)
			sum += p1 + offset*x
		}
		o.Stdoutln(sum)
	}, &ecmodels.Execution{
		Want:     "18613426663617118",
		Estimate: 0.5,
	})
}
