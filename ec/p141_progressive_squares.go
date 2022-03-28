package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P141() *problem {
	return intInputNode(141, func(o command.Output, n int) {
		p := generator.Primes()
		squares := generator.SmallPowerGenerator(2)
		var sum int
		for a := 1; a*a < n; a++ {
			factors := generator.Factors(a, p)
			for _, f := range factors {
				if squares.Contains(f) {
					den := maths.Sqrt(f)
					for num := 1; num < den; num++ {
						if generator.Coprimes(num, den, p) {
							continue
						}
						k := a*a*num/den + a*num*num/(den*den)
						if squares.Contains(k) {
							sum += k
						}
					}
				}
			}
		}
		o.Stdoutln(sum)
	})
}
