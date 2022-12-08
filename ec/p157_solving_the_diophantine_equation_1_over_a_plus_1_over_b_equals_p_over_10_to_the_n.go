package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P157() *problem {
	return intInputNode(157, func(o command.Output, n int) {
		primes := generator.Primes()

		// We only need to check when a and b are relative primes.
		// Otherwise:
		// let c*A = a and c*B = b, where c is the factor a and b share.
		// 1/cA + 1/cB = P/ten
		// (cA + cB)/(ABc*c) = P/ten
		// (A + B)/ABc = P/ten
		// (A + B)/AB = cP/ten
		// Since P is a variable, then this just simplifies to solving the problem
		// for A and B (rather than a=c*A and b=c*B)
		cnt := 0
		ten := 1
		for i := 1; i <= n; i++ {
			ten *= 10
			for twos := 1; ten%twos == 0; twos *= 2 {
				for fives := 1; ten%fives == 0; fives *= 5 {
					// a = twos, b = fives
					// P = ten * (a + b) / ab
					p := ten * (twos + fives) / (twos * fives)
					cnt += primes.FactorCount(p)

					if twos != 1 && fives != 1 {
						// a = 1, b = twos*fives
						p := ten * (twos*fives + 1) / (twos * fives)
						cnt += primes.FactorCount(p)
					}
				}
			}
		}

		o.Stdoutln(cnt)
	}, []*execution{
		{
			args: []string{"1"},
			want: "20",
		},
		{
			args: []string{"9"},
			want: "53490",
		},
	})
}
