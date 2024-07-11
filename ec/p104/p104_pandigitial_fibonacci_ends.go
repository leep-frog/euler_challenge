package p104

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P104() *ecmodels.Problem {
	return ecmodels.NoInputNode(104, func(o command.Output) {
		frontA, frontB := 1, 1
		backA, backB := 1, 1

		for i := 2; ; i++ {
			fLen := len(parse.Itos(frontB))
			bLen := len(parse.Itos(backB))
			if bLen > 9 && fLen >= 9 {
				panBack := maths.Pandigital(backB % 1_000_000_000)
				panFront := maths.Pandigital(maths.Chop(frontB, 0, 9))
				if panFront && panBack {
					o.Stdoutln(i)
					return
				}
			}

			t := frontA + frontB
			frontA = frontB
			frontB = t
			if frontA > (math.MaxInt / 10) {
				frontA = frontA / 10
				frontB = frontB / 10
			}
			backA, backB = backB%100_000_000_000, (backA+backB)%100_000_000_000
		}
	}, &ecmodels.Execution{
		Want:     "329468",
		Estimate: 0.5,
	})
}
