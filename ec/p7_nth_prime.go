package eulerchallenge

import (
	"fmt"
	"time"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P7() *problem {
	return intInputNode(7, func(o command.Output, n int) {
		// ;o.Stdoutln(generator.Primes().Nth(n - 1))
		fmt.Println(time.Now())
		a := generator.PrimesUpTo(180_000_000).Nth(n - 1)
		fmt.Println(time.Now())
		b := generator.Primes().Nth(n - 1)
		fmt.Println(time.Now())
		c := generator.BetterPrimes().Nth(n - 1)
		fmt.Println(time.Now())
		o.Stdoutln(a, b, c)
	}, []*execution{
		{
			args: []string{"6"},
			want: "13 13",
		},
		{
			args: []string{"9"},
			want: "23 23",
		},
		{
			args: []string{"10001"},
			want: "104743 104743",
		},
		{
			args: []string{"10000001"},
			want: "",
		},
	})
}
