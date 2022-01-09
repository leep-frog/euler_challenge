package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P5() *command.Node {
	return command.SerialNodes(
		command.Description("Find the smallest integer that is a multiple of all integers up to N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {

			primer := generator.Primes()
			primer.Next()
			primes := map[int]int{}
			for i := 2; i < d.Int(N); i++ {
				for p, cnt := range generator.PrimeFactors(i, primer) {
					primes[p] = maths.Max(cnt, primes[p])
				}
			}
			product := 1
			for p, cnt := range primes {
				for i := 0; i < cnt; i++ {
					product *= p
				}
			}
			o.Stdoutln(product)
		}),
	)
}
