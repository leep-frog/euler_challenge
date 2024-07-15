package p346

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P346() *ecmodels.Problem {
	return ecmodels.IntInputNode(346, func(o command.Output, pow int) {
		n := maths.Pow(10, pow)

		m := map[int]int{}
		for base := 2; base*base+base+1 <= n; base++ {
			for val, pow := base*base+base+1, 3; val <= n; val, pow = val+maths.Pow(base, pow), pow+1 {
				m[val]++
			}
		}

		var sum int
		for k := range m {
			sum += k
		}

		o.Stdoutln(sum + 1)
	}, []*ecmodels.Execution{
		{
			Args: []string{"3"},
			Want: "15864",
		},
		{
			Args:     []string{"12"},
			Want:     "336108797689259276",
			Estimate: 0.5,
		},
	})
}
