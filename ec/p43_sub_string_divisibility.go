package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P43() *problem {
	return noInputNode(43, func(o command.Output) {
		perms := maths.StringPermutations([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
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
	})
}
