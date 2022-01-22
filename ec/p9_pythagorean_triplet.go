package eulerchallenge

import (
	"math"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P9() *problem {
	return intInputNode(9, func(o command.Output, n int) {
		for a := 1; a < n; a++ {
			for b := a + 1; b+a < n; b++ {
				c2 := (a*a + b*b)
				c := int(math.Sqrt(float64(c2)))
				if a+b+c == 1000 && maths.IsSquare(c2) {
					o.Stdoutf("%d", a*b*c)
					return
				}
			}
		}
		o.Terminatef("failed to find triplet")
	})
}
