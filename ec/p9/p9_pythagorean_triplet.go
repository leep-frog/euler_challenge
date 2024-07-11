package p9

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P9() *ecmodels.Problem {
	return ecmodels.IntInputNode(9, func(o command.Output, n int) {
		for a := 1; a < n; a++ {
			for b := a + 1; b+a < n; b++ {
				c2 := (a*a + b*b)
				c := int(math.Sqrt(float64(c2)))
				if a+b+c == 1000 && maths.IsSquare(c2) {
					o.Stdoutln(a * b * c)
					return
				}
			}
		}
		o.Terminatef("failed to find triplet")
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000"},
			Want: "31875000",
		},
	})
}
