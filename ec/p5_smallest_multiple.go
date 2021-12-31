package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func addPrimeFactors(i int, primer *generator.Generator, primes map[int]int) {
	if i == primer.Last() {
		primes[i] = 1
		primer.Next()
		return
	}

	for p, cnt := range primes {
		curCnt := 0
		for i%p == 0 {
			curCnt++
			i = i / p
		}
		if curCnt > cnt {
			primes[p] = curCnt
		}
		if i == 1 {
			break
		}
	}
}

func P5() *command.Node {
	return command.SerialNodes(
		command.Description("Find the smallest integer that is a multiple of all integers up to N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {

			primer := generator.Primes()
			primer.Next()
			primes := map[int]int{}
			for i := 2; i < d.Int(N); i++ {
				addPrimeFactors(i, primer, primes)
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
