package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P5() *problem {
	return intInputNode(5, func(o command.Output, n int) {
		primer := generator.Primes()
		primes := map[int]int{}
		for i := 2; i < n; i++ {
			for p, cnt := range primer.PrimeFactors(i) {
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
	}, []*execution{
		{
			args: []string{"10"},
			want: "2520",
		},
		{
			args: []string{"20"},
			want: "232792560",
		},
	})
}
