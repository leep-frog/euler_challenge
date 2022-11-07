package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P56() *problem {
	return noInputNode(56, func(o command.Output) {
		best := maths.Largest[int, int]()
		for a := 1; a < 100; a++ {
			for b := 1; b < 100; b++ {
				best.Check(maths.BigPow(a, b).DigitSum())
			}
		}
		o.Stdoutln(best.Best())
	}, &execution{
		want: "972",
	})
}
