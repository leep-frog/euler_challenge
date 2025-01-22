package p518

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P518() *ecmodels.Problem {
	return ecmodels.IntInputNode(518, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int

		prev := 1000000

		for i := 0; p.Nth(i) < n; i++ {
			a := p.Nth(i) + 1

			if a > prev {
				fmt.Println(a)
				prev += 1000000
			}

			// First, do all integer sequences
			for j := 2; a*j*j < n; j++ {
				b, c := a*j, a*(j*j)
				if b <= n && c <= n && p.Contains(b-1) && p.Contains(c-1) {
					sum += a + b + c - 3
				}
			}

			sum += dfs(p, p.PrimeFactorsFast(a), n, a, 1)
		}
		fmt.Println(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "",
		},
		{
			Args: []string{"2"},
			Want: "",
		},
	})
}

func dfs(p *generator.Prime, factors [][]int, n int, a int, denom int) int {
	if len(factors) == 0 {

		if denom == 1 {
			return 0
		}

		var sum int

		for numer := denom + 1; ((a/denom)/denom)*numer*numer < n; numer++ {
			if !p.Coprimes(numer, denom) {
				continue
			}

			b := (a / denom) * numer
			c := (b / denom) * numer
			if b <= n && c <= n && p.Contains(b-1) && p.Contains(c-1) {
				sum += a + b + c - 3
			}
		}

		return sum
	}

	factor, cnt := factors[0][0], factors[0][1]
	var sum int
	for curCnt := 0; curCnt <= cnt; curCnt += 2 {
		sum += dfs(p, factors[1:], n, a, denom*maths.Pow(factor, curCnt/2))
	}
	return sum
}
