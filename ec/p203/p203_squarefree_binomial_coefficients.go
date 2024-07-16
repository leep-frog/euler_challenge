package p203

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P203() *ecmodels.Problem {
	return ecmodels.IntInputNode(203, func(o command.Output, n int) {

		p := generator.Primes()

		set := map[int]bool{}
		for row := []int{1}; len(row) <= n; row = updateRow(row) {
			for _, v := range row {
				if !hasSquare(v, p) {
					set[v] = true
				}
			}
		}

		var sum int
		for v := range set {
			sum += v
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "105",
		},
		{
			Args: []string{"51"},
			Want: "34029210557338",
		},
	})
}

func updateRow(row []int) []int {
	next := []int{1}
	for i, v := range row[1:] {
		next = append(next, v+row[i])
	}
	return append(next, 1)
}

func hasSquare(k int, p *generator.Prime) bool {
	for _, cnt := range p.PrimeFactors(k) {
		if cnt > 1 {
			return true
		}
	}
	return false
}
