package p58

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P58() *ecmodels.Problem {
	return ecmodels.NoInputNode(58, func(o command.Output) {
		cur := 1
		p := generator.Primes()
		count, numPrimes := 1, 0
		jump := 2
		for i := 1; ; i++ {
			for j := 0; j < 4; j++ {
				cur += jump
				count++
				if p.Contains(cur) {
					numPrimes++
				}
			}
			if numPrimes*10 < count {
				o.Stdoutln(2*i + 1)
				return
			}
			jump += 2
		}
	}, &ecmodels.Execution{
		Want: "26241",
	})
}
