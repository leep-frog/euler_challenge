package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P73() *problem {
	return intInputNode(73, func(o command.Output, n int) {
		lower := fraction.New(1, 3)
		upper := fraction.New(1, 2)
		p := generator.Primes()
		unique := map[string]bool{}
		for den := 4; den <= n; den++ {
			for num := den / 3; ; num++ {
				f := fraction.New(num, den)
				if maths.GTE(f, upper) {
					break
				}
				if maths.GT(f, lower) {
					unique[fraction.Simplify(f.N, f.D, p).String()] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	})
}
