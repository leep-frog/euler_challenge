package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P235() *problem {
	return noInputNode(235, func(o command.Output) {
		left := 1.0
		right := 1.1
		prevMiddle := left
		middle := (left + right) / 2.0
		for maths.Abs(middle-prevMiddle) > 0.000_000_000_000_1 {
			mv := 600_000_000_000 + s(middle)
			if mv > 0 {
				left = middle
			} else {
				right = middle
			}
			prevMiddle = middle
			middle = (left + right) / 2.0
		}

		// Printf rounds for us
		o.Stdoutf("%.12f\n", middle)
	}, &execution{
		want: "1.002322108633",
	})
}

func u(k int, r_k float64) float64 {
	return float64(900-3*k) * r_k
}

func s(r float64) float64 {
	r_k := 1.0
	sum := 0.0
	for k := 1; k <= 5000; k++ {
		sum += u(k, r_k)
		r_k *= r
	}
	return sum
}
