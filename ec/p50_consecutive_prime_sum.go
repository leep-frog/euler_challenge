package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P50() *problem {
	return intInputNode(50, func(o command.Output, n int) {
		best := maths.Largest[int, int]()
		primes := generator.Primes()
		for i := 0; primes.Nth(i) < n; i++ {
			pi := primes.Nth(i)
			sum := pi + primes.Nth(i+1)
			for j := 2; sum < n; j++ {
				if primes.Contains(sum) {
					best.IndexCheck(sum, j)
				}
				sum += primes.Nth(i + j)
			}
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	}, []*execution{
		{
			args: []string{"1000000"},
			want: "997651 543",
		},
		{
			args: []string{"1000"},
			want: "953 21",
		},
		{
			args: []string{"100"},
			want: "41 6",
		},
	})
}
