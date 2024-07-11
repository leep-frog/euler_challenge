package p43

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P43() *ecmodels.Problem {
	return ecmodels.NoInputNode(43, func(o command.Output) {
		perms := combinatorics.StringPermutations([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
		primes := generator.Primes()

		var sum int
		for _, perm := range perms {
			if perm[0:1] == "0" {
				continue
			}

			yup := true
			for i := 1; i < 8; i++ {
				if parse.Atoi(perm[i:i+3])%primes.Nth(i-1) != 0 {
					yup = false
					break
				}
			}
			if yup {
				sum += parse.Atoi(perm)
			}
		}
		o.Stdoutln(sum)
	}, &ecmodels.Execution{
		Want:     "16695334890",
		Estimate: 6,
	})
}
