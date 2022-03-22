package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P130() *problem {
	return intInputNode(130, func(o command.Output, n int) {
		diffs := []int{2, 4, 2, 2}
		g := generator.Primes()

		var count, sum int
		for i, j := 3, 1; ; i, j = i+diffs[j], (j+1)%len(diffs) {
			if generator.IsPrime(i, g) {
				continue
			}
			// Build map from one digit to required multiplier
			mults := make([]int, 10, 10)
			for m := 1; m <= 9; m++ {
				prod := i * m
				digits := maths.Digits(prod)
				mults[digits[len(digits)-1]] = m * i
			}

			k := 1
			for init := mults[1] / 10; init != 0; {
				k++
				need := (11 - (init % 10)) % 10
				init = (init + mults[need]) / 10
			}

			if (i-1)%k == 0 {
				count++
				sum += i
				if count >= n {
					o.Stdoutln(sum)
					return
				}
			}
		}
	})
}
