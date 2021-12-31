package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

// TODO: move all of these to helper directory
func getPrimeFactors(n int, p *generator.Generator) map[int]int {
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := p.Nth(i)
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func P3() *command.Node {
	return command.SerialNodes(
		command.Description("Find the largest prime factor of N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			factors := getPrimeFactors(d.Int(N), generator.Primes())

			max := 0
			for f := range factors {
				if f > max {
					max = f
				}
			}

			o.Stdoutln(max)
		}),
	)
}
