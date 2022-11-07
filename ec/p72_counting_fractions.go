package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P72() *problem {
	return intInputNode(72, func(o command.Output, n int) {
		p := generator.Primes()
		got := map[int]int{}
		count := 0

		for i := 2; i <= n; i++ {
			if p.Contains(i) {
				got[i] = i - 1
			} else {
				k := 1
				for num, cnt := range generator.PrimeFactors(i, p) {
					k *= got[num] * maths.Pow(num, cnt-1)
				}
				got[i] = k
			}
			count += got[i]
		}
		o.Stdoutln(count)
	}, []*execution{
		{
			args:     []string{"1000000"},
			want:     "303963552391",
			estimate: 6,
		},
		{
			args: []string{"8"},
			want: "21",
		},
	})
}
