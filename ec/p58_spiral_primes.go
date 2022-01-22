package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P58() *problem {
	return noInputNode(58, func(o command.Output) {
		cur := 1
		p := generator.Primes()
		count, numPrimes := 1, 0
		jump := 2
		for i := 1; ; i++ {
			for j := 0; j < 4; j++ {
				cur += jump
				count++
				if generator.IsPrime(cur, p) {
					numPrimes++
				}
			}
			if numPrimes*10 < count {
				o.Stdoutln(2*i + 1)
				return
			}
			jump += 2
		}
	})
}
