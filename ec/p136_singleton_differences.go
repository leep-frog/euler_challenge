package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

// I solved using this brute force approach, but took 20 minutes.
// Saw cool solution in problem thread that is implemented below
func brute136(n int) int {
	g := generator.Primes()
	var count int
	for k := 2; k < n; k++ {
		if k%1000000 == 0 {
			fmt.Println(k)
		}
		if diophantineDifferenceExactCount(k, 1, g) {
			fmt.Println(k)
			count++
		}
	}
	return count
}

// After going through brute and hints from forum, noticed that valid n's are one of the following:
// (1) n=4*p with p odd prime or p=1.
// (2) n=4*4*p with p odd prime or p=1.
// or (3) n=p with p prime and p+1 mod 4 = 0.
func elegant136(n int) int {
	g := generator.Primes()
	count := 2
	for i := 1; g.Nth(i) < n; i++ {
		p := g.Nth(i)
		if 4*p <= n {
			count++
		}
		if 4*4*p <= n {
			count++
		}
		if (p+1)%4 == 0 {
			count++
		}
	}
	return count
}

func P136() *problem {
	return intInputNode(136, func(o command.Output, n int) {
		o.Stdoutln(elegant136(n))
	}, []*execution{
		{
			args: []string{"100"},
			want: "25",
		},
		{
			args:     []string{"50_000_000"},
			want:     "2544559",
			estimate: 6.5,
		},
	})
}
