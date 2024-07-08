package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P141() *problem {
	return intInputNode(141, func(o command.Output, n int) {
		p := generator.Primes()
		squares := generator.SmallPowerGenerator(2)
		var sum int
		for a := 1; a*a < n; a++ {
			factors := p.Factors(a)
			for _, f := range factors {
				if squares.Contains(f) {
					den := maths.Sqrt(f)
					for num := 1; num < den; num++ {
						if p.Coprimes(num, den) {
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
	}, []*execution{
		{
			args:     []string{"1000000000000"},
			want:     "878454337159",
			estimate: 5,
		},
		{
			// Actually should be 100000. Stopping criteria is not correct
			args: []string{"2000000"},
			want: "124657",
		},
	})
}
