package p206

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/rgx"
)

func P206() *ecmodels.Problem {
	return ecmodels.NoInputNode(206, func(o command.Output) {
		// The last digit is a zero, so we know that the second to last digit is
		// also a zero (based on it being a square value)
		r := rgx.New("^1[0-9]2[0-9]3[0-9]4[0-9]5[0-9]6[0-9]7[0-9]8[0-9]9$" /* _0 */)

		// The number is 17 digits long, which means the square root is >= sqrt(10^16) = 10^8
		// (but start with a value that ends in 3 for the reason below)
		start := 100_000_003

		// Since it ends in 9, the last digit must be a 3 or a 7
		// Therefore, we start at 3 and then alternately increment by 4 and 6
		for i, offset := start, 1; ; i, offset = i+3+offset, (offset+2)%4 {
			if _, ok := r.Match(fmt.Sprintf("%d", i*i)); ok {
				o.Stdoutf("%d0\n", i)
				return
			}
		}
	}, &ecmodels.Execution{
		Want: "1389019170",
	})
}
