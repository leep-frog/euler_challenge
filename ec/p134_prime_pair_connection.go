package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P134() *problem {
	return noInputNode(134, func(o command.Output) {
		g := generator.Primes()
		var sum int
		for i := 2; g.Nth(i) < 1000000; i++ {
			p1, p2 := g.Nth(i), g.Nth(i+1)

			iter := maths.Pow(10, len(strconv.Itoa(p1)))
			v := p1
			for ; v%p2 != 0; v += iter {
			}
			sum += v
		}
		o.Stdoutln(sum)
	})
}
