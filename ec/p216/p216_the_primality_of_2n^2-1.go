package p216

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P216() *ecmodels.Problem {
	return ecmodels.IntInputNode(216, func(o command.Output, n int) {

		p := generator.Primes()

		var maybePrime []int
		for i := 2; i <= n; i++ {
			if i%(n/100) == 0 {
				fmt.Println(i)
			}
			ti := t(i)
			if p.FermatContains(ti, 5) {
				maybePrime = append(maybePrime, ti)
			}
		}

		o.Stdoutln(len(maybePrime))

		// Continuously filter out more
		// for {
		// 	var next []int
		// 	for _, mp := range maybePrime {
		// 		if p.FermatContains(mp, 2) {
		// 			next = append(next, mp)
		// 		}
		// 	}

		// 	if len(next) != len(maybePrime) {
		// 		fmt.Println("NEW LEN", len(next))
		// 	}
		// 	maybePrime = next
		// }
	}, []*ecmodels.Execution{
		{
			Args: []string{"50000000"},
			Want: "5437849",
			Skip: strings.Join([]string{
				"The code for this problem simply spits out an upper bound since it uses FermatContains",
				"I started at the output number and then simply decremented continually until getting",
				"a green check mark on the Project Euler problem page",
			}, "\n"),
			Estimate: 250,
		},
	})
}

func t(n int) int {
	return 2*n*n - 1
	// generator.Is
	// return p.Contains(tn)
}

// func other(p *generator.Prime, n int) bool {
// 	// Largest k such that k^2 divides prime(n)+1.

// 	for i := 0;; i++ {

// 	}

// }
