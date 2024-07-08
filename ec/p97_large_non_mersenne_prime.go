package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P97() *problem {
	return noInputNode(97, func(o command.Output) {
		// coef * 2^exp + 1
		coef := 28433
		exp := 7830457

		mod := maths.Pow(10, 10)

		v := coef
		for i := 0; i < exp; i++ {
			v *= 2
			v = v % mod
		}
		o.Stdoutln(strconv.Itoa((v + 1) % mod))
	}, &execution{
		want: "8739992577",
	})
}
