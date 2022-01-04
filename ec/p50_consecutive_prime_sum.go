package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P50() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=50"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			best := maths.Largest()
			primes := generator.Primes()
			for i := 0; primes.Nth(i) < n; i++ {
				pi := primes.Nth(i)
				sum := pi + primes.Nth(i+1)
				for j := 2; sum < n; j++ {
					if generator.IsPrime(sum, primes) {
						best.IndexCheck(sum, j)
					}
					sum += primes.Nth(i + j)
				}
			}
			o.Stdoutln(best.BestIndex(), best.Best())
		}),
	)
}
